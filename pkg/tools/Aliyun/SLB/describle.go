package SLB

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

// 获取所有slb信息
func (d *DefaultSet) CreateDescribeLoadBalancersRequest() (*CreateDescribeLoadBalancersResponse, error) {
	request := slb.CreateDescribeLoadBalancersRequest()
	d.SetParam(request)
	response, err := d.Client.Client.DescribeLoadBalancers(request)
	if err != nil {
		fmt.Println("CreateDescribeLoadBalancersReques error:", err.Error())
		return nil, err
	}
	return (*CreateDescribeLoadBalancersResponse)(response), nil
}
