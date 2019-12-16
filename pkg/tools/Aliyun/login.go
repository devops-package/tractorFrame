package aliyun

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

var err error

type LoginClientStruct struct {
	IsInsecure bool
	Connect    *alidns.Client
}

func (login *LoginClientStruct) NewLogin(regionId, accessKeyId, accessKeySecret string) *LoginClientStruct {
	if login.IsInsecure == false {
		login.Connect, err = alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
		if err != nil {
			fmt.Print(err.Error())
		}
		login.IsInsecure = true
	}

	return login
}
