package service

import (
	"context"

	"encoding/json"
	"fmt"
	"frontend/biz/model"
	eino "frontend/hertz_gen/eino"
	"frontend/infra/rpc"

	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
    rpcproduct "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
    rpccart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
    frontendUtils "frontend/utils"


)

type EinoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

type Mess struct {
    Name string `json:"name"`
    Num int `json:"num"`
}

func NewEinoService(Context context.Context, RequestContext *app.RequestContext) *EinoService {
	return &EinoService{RequestContext: RequestContext, Context: Context}
}

func (h *EinoService) Run(req *eino.EinoReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	m := model.Model

	if req.Ask == "" {
		fmt.Println("1")
	}

	pre := "请提取以下句子中的名词和量词,只返回一个json格式数据,名词的关键字设为name,量词的关键字设为num:"

	// 准备消息
	messages := []*schema.Message{
		//schema.UserMessage("请提取以下句子中的名词和量词,只返回一个json格式数据,名词的关键字设为name,量词的关键字设为num:下单一箱牛奶"),
		schema.UserMessage(pre + req.Ask),
	}
    
	// 生成回复
	response, err := m.Generate(h.Context, messages)
	if err != nil {

		fmt.Println(err)

		panic(err)
	}
	fmt.Println(response.Content)
	var mess Mess

	err = json.Unmarshal([]byte(response.Content), &mess)
	if err != nil {
		fmt.Println("json wrong:", err)
	}

    p, err := rpc.ProductClient.GetProductByName(h.Context, &rpcproduct.GetProductByNameReq{Name: mess.Name})
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    fmt.Println(p.Id)
    _, err = rpc.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
		Item: &rpccart.CartItem{
			ProductId: p.Id,
			Quantity: uint32(mess.Num),
		},
	})
    if err != nil {
        panic(err)
    }
	return utils.H{
		"productname" : mess.Name,
		"num": mess.Num,
	}, nil
}
