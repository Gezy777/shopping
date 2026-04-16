package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/product/biz/model"
	product "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type GetProductByNameService struct {
	ctx context.Context
}

// NewGetProductByNameService new GetProductByNameService
func NewGetProductByNameService(ctx context.Context) *GetProductByNameService {
	return &GetProductByNameService{ctx: ctx}
}

// Run create note info
func (s *GetProductByNameService) Run(req *product.GetProductByNameReq) (resp *product.GetProductByNameResp, err error) {
	// Finish your business logic.
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	p, err := productQuery.GetByName(req.Name)
	if err != nil {
		return nil, err
	}

	return &product.GetProductByNameResp{
		Id: uint32(p.ID),
	}, nil
}
