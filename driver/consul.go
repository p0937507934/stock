package driver

import (
	"github.com/hashicorp/consul/api"
)

func Register(id string, name string, address string, tags []string) {
	cfg := api.DefaultConfig()
	cfg.Address = "localhost:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	service := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: address,
		Tags:    tags,
	}
	err = client.Agent().ServiceRegister(service)
	if err != nil {
		panic(err)
	}
}

func DeRegister(id string) {
	cfg := api.DefaultConfig()
	cfg.Address = "localhost:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	err = client.Agent().ServiceDeregister(id)
	if err != nil {
		panic(err)
	}
}
