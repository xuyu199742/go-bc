package homestayOrder

import (
	"net/http"

	"go-bc/app/order/cmd/api/internal/logic/homestayOrder"
	"go-bc/app/order/cmd/api/internal/svc"
	"go-bc/app/order/cmd/api/internal/types"
	"go-bc/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserHomestayOrderListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserHomestayOrderListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := homestayOrder.NewUserHomestayOrderListLogic(r.Context(), ctx)
		resp, err := l.UserHomestayOrderList(req)
		result.HttpResult(r, w, resp, err)
	}
}
