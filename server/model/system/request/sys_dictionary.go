package request

import (
	"project/model/common/request"
	"project/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
