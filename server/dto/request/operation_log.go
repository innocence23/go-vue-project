package request

import "project/entity"

type OperationLogSearch struct {
	entity.OperationLog
	PageInfo
}
