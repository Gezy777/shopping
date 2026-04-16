package main

import (
	"context"
	"fmt"
	"log"

	"cwgo_test/kitex_gen/api"
	"cwgo_test/kitex_gen/api/echo"

	"github.com/bytedance/gopkg/cloud/metainfo"
	//"github.com/cloudwego/kitex-examples/kitex_gen/pbapi"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
    r, err := consul.NewConsulResolver("127.0.0.1:8500")
    if err != nil {
        log.Fatal(err)
		fmt.Println("3")
    }
	fmt.Println("1")
    c, err := echo.NewClient("cwgo_test", client.WithResolver(r),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
	)
	fmt.Println("2")
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("4")
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "cwgo_test_client")
	res, err := c.Echo(ctx, &api.Request{Message: "hello"})
	fmt.Println("5")
	if err != nil {
		fmt.Println("7")
		log.Fatal(err)
	}
	fmt.Println("6")
	fmt.Printf("%v", res)
}
