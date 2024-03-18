package common

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

func NacosRegister() {
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "2bdf0290-9626-41e8-821f-00c954bc107e", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}
	namingClient, _ := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	success, _ := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "10.2.171.791",
		Port:        8086,
		ServiceName: "demo.go",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "cluster-a", // default value is DEFAULT
		GroupName:   "group-a",   // default value is DEFAULT_GROUP
	})
	log.Printf("Register instance success: %v", success)
	// Create config client for dynamic configuration
	config, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	getConfig, _ := config.GetConfig(vo.ConfigParam{
		DataId: "user",
		Group:  "dev",
	})
	fmt.Println(getConfig)

}
func GetNacos() {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "2bdf0290-9626-41e8-821f-00c954bc107e", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			"127.0.0.1",
			8848,
		),
	}
	// Another way of create naming client for service discovery (recommend)
	namingClient, _ := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	service, _ := namingClient.GetService(vo.GetServiceParam{
		ServiceName: "user",
		GroupName:   "dev",
	})
	fmt.Println(service)

}

type Goods struct {
	Id         int
	GoodsName  string
	GoodsPrice float64
	GoodsNum   int
	State      int
	Category   string
}
