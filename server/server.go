package main

import (
	"flag"
	"fmt"

	"server/app"
	"server/config"
	"server/handler"
	"server/proto/hello"
	"server/repository"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
)

var debug bool

func init() {
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
}
func main() {
	defer panicRecover()
	flag.Parse()
	config.LoadConfig(debug)
	app.InitLogger()
	app.InitMysqlDb()

	svrRun()
}

func svrRun() {
	consulDsn := config.GetConsul().Host + ":" + config.GetConsul().Port
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{consulDsn}
	})

	svrCfg := config.GetServer()
	srv := service.NewService(
		service.Name(svrCfg.Name),
		service.Version(svrCfg.Version),
		// 注册consul中心
		micro.Registry(reg),
	)

	//hello
	if err := hello.RegisterHelloHandler(srv.Server(), &handler.HelloHandler{
		HelloRepository: &repository.HelloRepository{},
	}); err != nil {
		panic("server run error")
	}

	// Run server
	app.ZapLog.Info(svrCfg.Name, svrCfg.Name+" 服务启动成功")
	if err := srv.Run(); err != nil {
		panic("server run error")
	}
}

func panicRecover() {
	if err := recover(); err != nil {
		msg := fmt.Sprintf("panic %v \n", err)
		app.ZapLog.Error("panicRecover", msg)
	}
}
