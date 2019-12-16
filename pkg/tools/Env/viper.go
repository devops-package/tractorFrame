package Env

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
)

type Consul struct {
	Address string
	Port    int
	Error   error

}

func NewConsul(c *Consul) *Consul{

	return c
}

func (c *Consul) Get(key string) interface{} {
	host := fmt.Sprintf("%s:%s",c.Address,c.Port)
	viper.SetConfigType("yaml")
	if err := viper.AddRemoteProvider("consul", host, key); err != nil {
		log.Fatal(err)
	}
	if err := viper.ReadRemoteConfig(); err != nil {
		log.Fatal(err)
	}
	return viper.Get(key)
}