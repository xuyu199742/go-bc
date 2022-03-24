package request

import (
	"go-bc/admin/model/common/request"
	"go-bc/admin/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
