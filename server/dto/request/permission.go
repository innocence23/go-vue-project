package request

import (
	"project/model/system"
)

// api分页条件查询及排序结构体
type SearchPermissionParams struct {
	system.Permission
	PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}