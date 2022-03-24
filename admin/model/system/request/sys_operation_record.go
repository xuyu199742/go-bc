package request

import (
	"go-bc/admin/model/common/request"
	"go-bc/admin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
