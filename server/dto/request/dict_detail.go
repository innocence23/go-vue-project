package request

import (
	"project/entity"
)

type DictDetailSearch struct {
	entity.DictDetail
	PageInfo
}
