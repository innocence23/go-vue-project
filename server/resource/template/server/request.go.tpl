package request

import (
	"project/model/autocode"
	"project/dto/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}