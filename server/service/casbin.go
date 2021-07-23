package service

import (
	"errors"
	"project/dto/request"
	"project/model/system"
	"project/zvar"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

type CasbinService struct {
}

func (casbinService *CasbinService) Update(roleId string, casbinInfos []request.CasbinInfo) error {
	casbinService.ClearCasbin(0, roleId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		cm := system.CasbinModel{
			Ptype:  "p",
			RoleId: roleId,
			Path:   v.Path,
			Method: v.Method,
		}
		rules = append(rules, []string{cm.RoleId, cm.Path, cm.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if success == false {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := zvar.DB.Table("casbin_rule").Model(&system.CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

func (casbinService *CasbinService) GetPermByRoleId(roleId string) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, roleId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(zvar.DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(zvar.Config.Casbin.ModelPath, a)
		syncedEnforcer.AddFunction("ParamsMatch", casbinService.ParamsMatchFunc)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}

func (casbinService *CasbinService) ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	return util.KeyMatch2(key1, key2)
}

func (casbinService *CasbinService) ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return casbinService.ParamsMatch(name1, name2), nil
}
