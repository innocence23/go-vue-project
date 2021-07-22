package request

import (
	"project/model/autocode"
	"project/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}