package register

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"su/common"
	"su/goods"
	"su/logic"
)

func InitRegister() {
	flag.Parse()
	g := grpc.NewServer()
	listen, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println(err)
	}
	client, err := common.CreateConsul()
	if err != nil {
		panic(err)
	}
	err = common.HealthConsul(client, "10.2.171.79", listen.Addr().(*net.TCPAddr).Port)
	if err != nil {
		panic(err)
	}
	grpc_health_v1.RegisterHealthServer(g, health.NewServer())
	goods.RegisterStreamGreeterServer(g, &logic.Server{})
	fmt.Println("started...", listen.Addr())
	err = g.Serve(listen)

}
