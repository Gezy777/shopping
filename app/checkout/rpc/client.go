package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/conf"
	checkoutUtils "github.com/cloudwego/biz-demo/gomall/app/checkout/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	PaymentClient paymentservice.Client
	ProductClient productcatalogservice.Client
	CartClient cartservice.Client
	OrderClient orderservice.Client
	once sync.Once
)

func Init() {
	once.Do(func() {
		initPaymentClient()
		initProductClient()
		initCartClient()
		initOrderClient()
	})
}

func initPaymentClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	PaymentClient, err = paymentservice.NewClient("payment", client.WithResolver(r))
	checkoutUtils.MustHandleError(err)
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	checkoutUtils.MustHandleError(err)
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	checkoutUtils.MustHandleError(err)
}


func initOrderClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	checkoutUtils.MustHandleError(err)
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	checkoutUtils.MustHandleError(err)
}