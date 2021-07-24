package service

import (
	"errors"
	"project/entity"
	"project/zvar"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type BaseMenuService struct {
}

func (baseMenuService *BaseMenuService) Create(menu entity.Menu) error {
	if !errors.Is(zvar.DB.Where("name = ?", menu.Name).First(&entity.Menu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return zvar.DB.Create(&menu).Error
}

func (baseMenuService *BaseMenuService) Update(menu entity.Menu) (err error) {
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort
	err = zvar.DB.Model(&entity.Menu{}).Where("id = ?", menu.ID).Updates(upDateMap).Error
	return err
}

func (baseMenuService *BaseMenuService) Find(id float64) (err error, menu entity.Menu) {
	err = zvar.DB.Where("id = ?", id).First(&menu).Error
	return
}

func (baseMenuService *BaseMenuService) Hidden(id float64) (err error) {
	err = zvar.DB.Where("id = ?", id).Update("hidden", true).Error
	return
}

func (baseMenuService *BaseMenuService) Display(id float64) (err error) {
	err = zvar.DB.Where("id = ?", id).Update("hidden", false).Error
	return
}

//------ TreeList （列表）
func (baseMenuService *BaseMenuService) TreeList() (err error, list interface{}, total int64) {
	var menuList []entity.Menu
	err, treeMap := baseMenuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = baseMenuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

func (baseMenuService *BaseMenuService) getBaseChildrenList(menu *entity.Menu, treeMap map[string][]entity.Menu) (err error) {
	menu.Children = treeMap[cast.ToString(menu.ID)]
	for i := 0; i < len(menu.Children); i++ {
		err = baseMenuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (baseMenuService *BaseMenuService) getMenuTreeMap(roleId string) (err error, treeMap map[string][]entity.Menu) {
	var allMenus []entity.Menu
	treeMap = make(map[string][]entity.Menu)
	err = zvar.DB.Where("role_id = ?", roleId).Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//------ end

//------ GetMenuTree （用户菜单）
func (baseMenuService *BaseMenuService) GetMenuTree(roleId string) (err error, menus []entity.Menu) {
	err, menuTree := baseMenuService.getMenuTreeMap(roleId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = baseMenuService.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

func (baseMenuService *BaseMenuService) getBaseMenuTreeMap() (err error, treeMap map[string][]entity.Menu) {
	var allMenus []entity.Menu
	treeMap = make(map[string][]entity.Menu)
	err = zvar.DB.Order("sort").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

func (baseMenuService *BaseMenuService) getChildrenList(menu *entity.Menu, treeMap map[string][]entity.Menu) (err error) {
	menu.Children = treeMap[cast.ToString(menu.ID)]
	for i := 0; i < len(menu.Children); i++ {
		err = baseMenuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//------ end
