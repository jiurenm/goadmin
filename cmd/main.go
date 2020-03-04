package main

import (
	"admin/internal/di"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func register() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "localhost",
			ContextPath: "/nacos",
			Port:        8848,
		},
	}
	namingClient, _ := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
	})
	success, _ := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "localhost",
		Port:        8083,
		ServiceName: "tool",
		Weight:      1,
		ClusterName: "",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata: map[string]string{
			"preserved.register.source": "SPRING_CLOUD",
		},
	})
	if !success {
		fmt.Println("register error")
	}
}

func main() {
	register()

	_, _, _ = di.InitApp()
}
