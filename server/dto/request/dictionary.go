package request

import (
	"project/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	PageInfo
}
