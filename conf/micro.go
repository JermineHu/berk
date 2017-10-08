// This file "micro.go" is created by Jermine Hu at 6/22/16.
// Copyright © 2016 - Jermine Hu. All rights reserved
package conf

import (
	dc "git.vdo.space/foy/config"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-os/config"
	config_consul "github.com/micro/go-os/config/source/consul"
	"github.com/micro/go-os/trace"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

const (
	PACKAGE_NAME      = "service"
	HTTP_SERVICE_NAME = "tech-ngs-icarus-http"
	HTTP_SERVICE_PORT = ":9701"
	RPC_SERVICE_NAME  = "tech-ngs-icarus"
	RPC_SERVICE_PORT  = ":19701"
	SERVICE_VERSION   = "1.0"
)

const (
	CONSUL_ADDRESSES_KEY = "CONSUL_ADDRS"
)

func Init() {
	reg = consul.NewRegistry(
		registry.Addrs(os.Getenv(CONSUL_ADDRESSES_KEY)),
	)

	cmd.Init(
		cmd.Registry(&reg),
	)
}

var (
	reg    registry.Registry
	conf   config.Config
	tracer trace.Trace
)

//regConf 注册配置
func regConf() {

	dc.SetMGOUserName(Conf_GetValue(MGO_USERNAME))
	dc.SetMGOPassword(Conf_GetValue(MGO_PASSWORD))
	dc.SetMGOAddress(Conf_GetValue(MGO_ADDRESSES))
	dc.SetMGOPort(Conf_GetValue(MGO_PORT))
	dc.SetMGODatabase(Conf_GetValue(MGO_DATABASE))
	dc.SetResourceAddr(Conf_GetValue(ENV_RESOURCE_ADDR))
	dc.SetResourcePath(Conf_GetValue(ENV_RESOURCE_PATH))
}

func InitConfig() config.Config {
	conf = config.NewConfig(
		config.PollInterval(time.Second*10),
		config.WithSource(config_consul.NewSource(config.SourceHosts(os.Getenv(CONSUL_ADDRESSES_KEY)))),
	)

	regConf()

	return conf
}

func InitTracer() trace.Trace {
	if conf == nil {
		log.Fatal("[Loading] Config must be inited before init tracer")
	}

	tracer = trace.NewTrace(
		trace.Collectors(Conf_GetValue(ZIPKIN_ADDRESS)),
	)
	return tracer
}

func App(rpcPort string) micro.Service {
	if tracer == nil {
		log.Fatal("[Loading] Tracer must be inited before init tracer")
	}

	//b := broker.NewBroker(
	//	broker.Addrs(rpcPort+"1"),
	//)

	srv := server.NewServer(
		server.Address(rpcPort),
		server.Name(RPC_SERVICE_NAME),
		//server.Broker(b),
	)

	service := micro.NewService(
		micro.Registry(reg),
		micro.Server(srv),
		//micro.RegisterInterval(time.Second*5),
		micro.Name(RPC_SERVICE_NAME),
		micro.Version(SERVICE_VERSION),
		micro.WrapHandler(MDBHandler),
	)

	regConf()

	// proto.RegisterIcarusServicesHandler(service.Server(), new(services.Icarus)) //注册服务处理

	return service
}

func MDBHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		//mdb := dc.GetMongo()
		//db_ctx := context.WithValue(ctx, services.MDB_CONTEXT, mdb)
		//err := fn(db_ctx, req, rsp)
		//db_ctx.Done()
		//return err
		return nil
	}
}

func ControllerHandler(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		//c_ctx := context.WithValue(ctx, echo.C_CONTEXT, controller)
		//err := fn(c_ctx, req, rsp)
		//c_ctx.Done()
		//return err
		return nil
	}
}
