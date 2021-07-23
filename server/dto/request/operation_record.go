package request

import (
	"project/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	PageInfo
}
