package request

import (
	"project/entity"
)

type DictSearch struct {
	entity.Dict
	PageInfo
}
