package homestayOrder

import (
	"context"

	"go-bc/app/order/cmd/api/internal/svc"
	"go-bc/app/order/cmd/api/internal/types"
	"go-bc/app/order/cmd/rpc/order"
	"go-bc/common/ctxdata"
	"go-bc/common/tool"
	"go-bc/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomestayOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserHomestayOrderListLogic {
	return UserHomestayOrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderListLogic) UserHomestayOrderList(req types.UserHomestayOrderListReq) (*types.UserHomestayOrderListResp, error) {

	userId := ctxdata.GetUidFromCtx(l.ctx) //当前登陆用户id

	resp, err := l.svcCtx.OrderRpc.UserHomestayOrderList(l.ctx, &order.UserHomestayOrderListReq{
		UserId:      userId,
		TraderState: req.TradeState,
		PageSize:    req.PageSize,
		LastId:      req.LastId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("获取用户订单列表失败"), "err : %v ,req:%+v", err, req)
	}

	var typesUserHomestayOrderList []types.UserHomestayOrderListView

	if len(resp.List) > 0 {

		for _, homestayOrder := range resp.List {

			var typeHomestayOrder types.UserHomestayOrderListView
			_ = copier.Copy(&typeHomestayOrder, homestayOrder)

			typeHomestayOrder.OrderTotalPrice = tool.Fen2Yuan(homestayOrder.OrderTotalPrice)

			typesUserHomestayOrderList = append(typesUserHomestayOrderList, typeHomestayOrder)
		}
	}

	return &types.UserHomestayOrderListResp{
		List: typesUserHomestayOrderList,
	}, nil
}
