package aliyun

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"reflect"
)

var err error

type DnsClient struct {
	IsInsecure bool
	Connect    *alidns.Client

}

type EcsClient struct {
	IsInsecure bool
	Connect	   *ecs.Client
	Request    *ecs.DescribeInstancesRequest
}

func (login *DnsClient ) NewLogin(regionId, accessKeyId, accessKeySecret string) *DnsClient {
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

type Login interface {
	NewLogin(regionId, accessKeyId, accessKeySecret string) *DnsClient
}

func (E *EcsClient) SetParam(key, value string) {
	request := E.Request
	x :=reflect.ValueOf(&request).Elem().FieldByName("InstanceNetworkType")
	//reflect.ValueOf(&request).FieldByName(key).SetString(value)
	fmt.Printf("%+v\n",x)
	fmt.Printf("%+v",key)
	fmt.Printf("%+v",value)
}

//func (E *EcsClient) Instances() {
//	response, err := E.Connect.DescribeInstances(E.Request)
//	if err != nil {
//		fmt.Print(err.Error())
//	}
//	fmt.Printf("response is %#v\n", response)
//}