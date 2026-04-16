package eino

import (
	"context"

	"frontend/biz/service"
	"frontend/biz/utils"
	eino "frontend/hertz_gen/eino"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Eino .
// @router /eino [POST]
func Eino(ctx context.Context, c *app.RequestContext) {
	var err error
	var req eino.EinoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewEinoService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "about", utils.WarpResponse(ctx, c, resp))
}
