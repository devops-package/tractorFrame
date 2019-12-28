package aliyun

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

var err error

type DnsClient struct {
	IsInsecure bool
	Connect    *alidns.Client
}

type EcsClient struct {
	IsInsecure bool
	Connect    *ecs.Client
	Request    *ecs.DescribeInstancesRequest
	PageSize   string
	PageNumber int
	Response   *ecs.DescribeInstancesResponse
}

func (login *DnsClient) NewLogin(regionId, accessKeyId, accessKeySecret string) *DnsClient {
	if login.IsInsecure == false {
		login.Connect, err = alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
		if err != nil {
			fmt.Print(err.Error())
		}
		login.IsInsecure = true
	}

	return login
}

func (login *EcsClient) NewLogin(regionId, accessKeyId, accessKeySecret string) *EcsClient {
	if login.IsInsecure == false {
		login.Connect, err = ecs.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
		if err != nil {
			fmt.Print(err.Error())
		}
		login.IsInsecure = true
		login.Request = ecs.CreateDescribeInstancesRequest()
	}

	return login
}

//type Login intererface regionId, accessKeyId, accessKeySecret string
// regionId, accessKeyId, accessKeySecret string
type Login interface {
	NewLogin(regionId, accessKeyId, accessKeySecret string) *DnsClient
}
