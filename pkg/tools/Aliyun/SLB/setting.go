package SLB

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/devops-package/tractorFrame/pkg/tools/Aliyun"
)

type DefaultSet struct {
	Client          *aliyun.SLBClient
	Scheme          string
	RequestPageSize requests.Integer
}

type CreateDescribeLoadBalancersResponse slb.DescribeLoadBalancersResponse
type CreateDescribeHealthStatusResponse slb.DescribeHealthStatusResponse
type CreateDescribeLoadBalancerAttributeResponse slb.DescribeLoadBalancerAttributeResponse

func SLB_Init(client *aliyun.SLBClient) *DefaultSet {
	setting := new(DefaultSet)
	setting.Client = client
	setting.Scheme = "https"
	setting.RequestPageSize = "100"
	return setting
}
