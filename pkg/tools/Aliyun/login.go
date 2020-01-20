package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

var err error

type SLBClient struct {
	IsInsecure    bool
	Connect       *slb.Client
	Request       *slb.DescribeLoadBalancersRequest
	PageSize      string
	PageNumber    int
	Response      *slb.DescribeLoadBalancersResponse
	HealthRequest *slb.DescribeHealthStatusRequest
}

type DnsClient struct {
	IsInsecure bool
	Connect    *alidns.Client
	Request    *alidns.DescribeDomainsRequest
	PageSize   string
	PageNumber int
	Response   *alidns.DescribeDomainsResponse
}

type EcsClient struct {
	IsInsecure bool
	Connect    *ecs.Client
	Request    *ecs.DescribeInstancesRequest
	PageSize   string
	PageNumber int
	Response   *ecs.DescribeInstancesResponse
}

func (login *SLBClient) NewLogin(regionId, accessKeyId, accessKeySecret string) *SLBClient {
	if login.IsInsecure == false {
		login.Connect, err = slb.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
		if err != nil {
			panic(err)
		}
		login.IsInsecure = true
		login.Request = slb.CreateDescribeLoadBalancersRequest()
		login.HealthRequest = slb.CreateDescribeHealthStatusRequest()
	}
	return login
}

func (login *DnsClient) NewLogin(regionId, accessKeyId, accessKeySecret string) *DnsClient {
	if login.IsInsecure == false {
		login.Connect, err = alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
		if err != nil {
			panic(err)
		}
		login.IsInsecure = true
		login.Request = alidns.CreateDescribeDomainsRequest()
	}

	return login
}

func (login *EcsClient) NewLogin(regionId, accessKeyId, accessKeySecret string) *EcsClient {
	if login.IsInsecure == false {
		login.Connect, err = ecs.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
		if err != nil {
			panic(err)
		}
		login.IsInsecure = true
		login.Request = ecs.CreateDescribeInstancesRequest()
	}

	return login
}
