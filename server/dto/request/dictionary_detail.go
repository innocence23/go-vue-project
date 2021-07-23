package request

import (
	"project/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	PageInfo
}
