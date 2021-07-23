package request

import (
	"project/model/autocode"
)

// 如果含有time.Time 请自行import time包
type AutoCodeExampleSearch struct {
	autocode.AutoCodeExample
	PageInfo
}
