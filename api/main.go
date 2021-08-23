package main

import (
	"flag"
	"fmt"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/micro/cli/v2"

	"api/api"
	"api/app"
	"api/client"
	"api/config"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug mode")

}
func main() {
	defer panicRecover()
	flag.Parse()
	config.LoadConfig(debug)
	app.InitRedis()
	app.InitLogger()

	svrRun()
}

func svrRun() {
	svrConfig := config.GetHttpServer()
	srv := httpServer.NewServer(
		server.Name(svrConfig.Name),
		server.Version(svrConfig.Version),
		server.Address(":"+svrConfig.Port),
	)

	consulDns := config.GetConsul().Host + ":" + config.GetConsul().Port
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{consulDns}
	})
	r := api.StartApi()
	hd := srv.NewHandler(r)

	if err := srv.Handle(hd); err != nil {
		panic("api 注册失败")
	}

	apiService := micro.NewService(
		micro.Server(srv),
		// 注册consul中心
		micro.Registry(reg),
		micro.Flags(
			&cli.StringFlag{
				Name: "debug",
			},
		),
	)
	apiService.Init()
	client.MicroService = apiService

	if err := apiService.Run(); err != nil {
		app.ZapLog.Error("server error", err.Error())
		panic(err)
	}
}
func panicRecover() {
	if err := recover(); err != nil {
		msg := fmt.Sprintf("panic %v \n", err)
		app.ZapLog.Error("panicRecover", msg)
	}
}
