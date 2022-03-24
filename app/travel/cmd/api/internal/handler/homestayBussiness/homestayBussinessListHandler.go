package homestayBussiness

import (
	"net/http"

	"go-bc/app/travel/cmd/api/internal/logic/homestayBussiness"
	"go-bc/app/travel/cmd/api/internal/svc"
	"go-bc/app/travel/cmd/api/internal/types"
	"go-bc/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HomestayBussinessListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HomestayBussinessListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayBussiness.NewHomestayBussinessListLogic(r.Context(), ctx)
		resp, err := l.HomestayBussinessList(req)
		result.HttpResult(r, w, resp, err)
	}
}
