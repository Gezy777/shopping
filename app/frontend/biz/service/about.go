package service

import (
	"context"

	common "frontend/hertz_gen/common"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type AboutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAboutService(Context context.Context, RequestContext *app.RequestContext) *AboutService {
	return &AboutService{RequestContext: RequestContext, Context: Context}
}


func (h *AboutService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	return utils.H{}, nil
}

