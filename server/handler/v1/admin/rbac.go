package admin

import (
	"project/service"
)

type rbacHandler struct {
	userService *service.UserService
}

func NewRbacHandler() *rbacHandler {
	return &rbacHandler{
		userService: &service.UserService{},
	}
}

// //组装jwt数据
// func (h *rbacHandler) genUserJwt(ctx *gin.Context, u *model.User) *dto.UserJWT {
// 	uj := &dto.UserJWT{
// 		ID:      u.ID,
// 		Account: u.Account,
// 		Email:   u.Email,
// 		Avatar:  u.Avatar,
// 	}
// 	id := cast.ToString(uj.ID)
// 	uj.Roles, _ = component.GetRolesForUser(id)
// 	uj.Permissions = component.GetPermissionsForUser(id)
// 	var tmp []int64
// 	for _, v := range uj.Roles {
// 		tmp = append(tmp, cast.ToInt64(v))
// 	}
// 	goctx := ctx.Request.Context()
// 	uj.Menus = h.UserService.GetMenus(goctx, tmp, u.Email)

// 	return uj
// }
