package service

import (
	"log"
	"project/zvar"
)

type RbacService struct {
}

func (rbacService *RbacService) CheckPermission(account string, permission string, method string) bool {
	if account == "admin" {
		return true
	}
	roles, _ := zvar.Enforcer.GetRolesForUser(account)
	for _, role := range roles {
		ok := zvar.Enforcer.HasPermissionForUser(role, permission, method)
		if ok {
			return ok
		}
	}
	return false
}

// 授权用户到角色
func (rbacService *RbacService) AddRoleForUser(account string, role string) (bool, error) {
	return zvar.Enforcer.AddRoleForUser(account, role)
}

// 授权用户到角色 批量
func (rbacService *RbacService) AddRolesForUser(account string, role []string) (bool, error) {
	rbacService.DeleteRolesForUser(account)
	return zvar.Enforcer.AddRolesForUser(account, role)
}

//  添加权限到角色
func (rbacService *RbacService) AddPermissionForUser(permission string, method string, role string) (bool, error) {
	return zvar.Enforcer.AddPermissionForUser(role, permission, method)
}

// 获取用户角色
func (rbacService *RbacService) GetRolesForUser(account string) ([]string, error) {
	return zvar.Enforcer.GetRolesForUser(account)
}

// 获取用户（角色）权限
func (rbacService *RbacService) GetPermissionsForRole(role string) []map[string]string {
	var permissions []map[string]string
	currentPermissions := zvar.Enforcer.GetPermissionsForUser(role)
	for _, currentPermission := range currentPermissions {
		permissions = append(permissions, map[string]string{
			"method":     currentPermission[2],
			"permission": currentPermission[1],
		})
	}
	return permissions
}

// 获取用户权限
func (rbacService *RbacService) GetPermissionsForUser(account string) []map[string]string {
	var permissions []map[string]string
	roles, _ := rbacService.GetRolesForUser(account)
	for _, role := range roles {
		rolePermissions := rbacService.GetPermissionsForRole(role)
		for _, rolePermission := range rolePermissions {
			permissions = append(permissions, rolePermission)
		}
	}
	return permissions
}

// 删除用户的所有角色
func (rbacService *RbacService) DeleteRolesForUser(account string) (bool, error) {
	return zvar.Enforcer.DeleteRolesForUser(account)
}

// 删除角色的权限
func (rbacService *RbacService) DeletePermissionsForUser(role string) (bool, error) {
	return zvar.Enforcer.DeletePermissionsForUser(role)
}

// 删除拥有对应角色的(用户角色权限)
func (rbacService *RbacService) DeleteRoleForUsers(role string) bool {
	users, err := zvar.Enforcer.GetUsersForRole(role)
	if err != nil {
		log.Fatal("获取具有角色的用户")
	}
	for _, user := range users {
		_, _ = zvar.Enforcer.DeleteRoleForUser(user, role)
	}
	return true
}
