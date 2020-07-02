package SLB

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/tonyjia87/tractorFrame/pkg/tools/Aliyun"
)

type DefaultSet struct {
	Client          *aliyun.SLBClient
	Scheme          string
	RequestPageSize requests.Integer
}

type CreateDescribeLoadBalancersResponse slb.DescribeLoadBalancersResponse

func SLB_Init(client *aliyun.SLBClient) *DefaultSet {
	setting := new(DefaultSet)
	setting.Client = client
	setting.Scheme = "https"
	setting.RequestPageSize = "100"
	return setting
}

func (d *DefaultSet) SetParam(request *slb.DescribeLoadBalancersRequest) {
	request.Scheme = d.Scheme
	request.PageSize = d.RequestPageSize
}
