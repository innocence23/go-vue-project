package request

import (
	"project/model/system"
)

type OperationLogSearch struct {
	system.OperationLog
	PageInfo
}
