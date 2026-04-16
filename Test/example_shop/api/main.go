package main

import (
	"context"
	"log"
	"time"

	"example_shop/kitex_gen/example/shop/item"
	"example_shop/kitex_gen/example/shop/item/itemservice"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	cli itemservice.Client
)

func main() {
	// 使用时请传入真实 etcd 的服务地址，本例中为 127.0.0.1:2379
  resolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
  if err != nil {
	  log.Fatal(err)
  }

  c, err := itemservice.NewClient("example.shop.item", client.WithResolver(resolver)) // 指定 Resolver
  if err != nil {
	  log.Fatal(err)
  }
	cli = c

  hz := server.New(server.WithHostPorts("localhost:8889"))

  hz.GET("/api/item", Handler)

  if err := hz.Run(); err != nil {
	  log.Fatal(err)
  }
}

func Handler(ctx context.Context, c *app.RequestContext) {
	req := item.NewGetItemReq()
	req.Id = 1024
	resp, err := cli.GetItem(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	c.String(200, resp.String())
}
