package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"cwgo_test/biz/dal"
	"cwgo_test/conf"
	"cwgo_test/kitex_gen/api/echo"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	opts := kitexInit()
	fmt.Println("5")
	svr := echo.NewServer(new(EchoImpl), opts...)
	fmt.Println("6")
	err = svr.Run()
	fmt.Println("7")
	if err != nil {
		fmt.Println("8")
		klog.Error(err.Error())
		fmt.Println(err)
	}
	fmt.Println("9")
}

func kitexInit() (opts []server.Option) {

	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))
	// thrift meta handler
	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	fmt.Println("1")
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
    fmt.Println("2")
	if err != nil {
		fmt.Println("3")
        log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println("4")
	opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
