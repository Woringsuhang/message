package common

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

func CreateConsul() (*api.Client, error) {
	client, err := api.NewClient(&api.Config{Address: "127.0.0.1:8500"})
	if err != nil {
		return client, err
	}
	return client, nil
}
func HealthConsul(client *api.Client, ip string, port int) error {
	return client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    "servers",
		Tags:    []string{"GRPC"},
		Port:    port,
		Address: ip,
		Check: &api.AgentServiceCheck{
			DeregisterCriticalServiceAfter: "30s",
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           fmt.Sprintf("%v:%v", ip, port),
		},
	})

}
