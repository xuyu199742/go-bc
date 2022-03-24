package request

import (
	"go-bc/admin/model/common/request"
	"go-bc/admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
