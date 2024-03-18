package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"net/http"
	"su/common"
	"su/goods"
)

func main() {
	dial, _ := grpc.Dial(fmt.Sprintf("consul://%v:%v/", "10.2.171.79", 8500)+"servers"+"?wait=14s", grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`))
	Client := goods.NewStreamGreeterClient(dial)
	fmt.Println(123)
	fmt.Println(Client)
	g := gin.Default()

	common.GetNacos()
	g.POST("/login", func(c *gin.Context) {
		Client.GoodsCreated(context.Background(), &goods.CreateGoodsReq{
			Image: "djask",
			Name:  "小鱼er",
		})
		token, _ := common.GenToken(1)

		c.JSON(http.StatusOK, gin.H{
			"messages": token,
		})
	})
	err := g.Run(":8888")
	if err != nil {
		panic(fmt.Sprintf("gin启动失败%v", err))
	}
}
