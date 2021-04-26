package conf

import (
	gokitConsul "github.com/go-kit/kit/sd/consul"
	consulAPI "github.com/hashicorp/consul/api"
)

// ConnectConsul ...
func ConnectConsul() (client gokitConsul.Client, err error) {
	consulConfig := consulAPI.DefaultConfig()
	consulConfig.Address = "http://consul-serv:8500"

	consulClient, err := consulAPI.NewClient(consulConfig)
	if err != nil {
		logger.Fatalf("Connect consul failed: " + err.Error())
		return
	}
	client = gokitConsul.NewClient(consulClient)
	return
}
