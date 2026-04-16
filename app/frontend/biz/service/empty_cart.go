package service

import (
	"context"

	frontendUtils "frontend/utils"

	common "frontend/hertz_gen/common"
	"frontend/infra/rpc"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	_, err = rpc.CartClient.EmptyCart(h.Context, &cart.EmptyCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	}) 
	if err != nil {
		return nil, err
	}
	return
}
