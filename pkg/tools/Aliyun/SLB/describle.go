package SLB

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

// 获取所有slb信息
//{
//"Address": "",
//"ResourceGroupId": "",
//"VSwitchId": "",
//"CreateTime": "",
//"AddressIPVersion": "ipv4",
//"LoadBalancerId": "",
//"PayType": "",
//"SlaveZoneId": "",
//"BusinessStatus": "Normal",
//"InternetChargeType": "",
//"RegionIdAlias": "",
//"LoadBalancerName": "",
//"VpcId": "",
//"NetworkType": "",
//"RegionId": "",
//"AddressType": "",
//"LoadBalancerStatus": "active",
//"MasterZoneId": "",
//"CreateTimeStamp": 1587124584000,
//"Tags": {
//"Tag": []
//}
//},
func (d *DefaultSet) CreateDescribeLoadBalancersRequest() (*CreateDescribeLoadBalancersResponse, error) {
	request := slb.CreateDescribeLoadBalancersRequest()
	request.Scheme = d.Scheme
	request.PageSize = d.RequestPageSize
	response, err := d.Client.Client.DescribeLoadBalancers(request)
	if err != nil {
		fmt.Println("CreateDescribeLoadBalancersRequest error:", err.Error())
		return nil, err
	}
	return (*CreateDescribeLoadBalancersResponse)(response), nil
}

// 调用DescribeHealthStatus查询后端服务器的健康状态。
// 			{
//				"ListenerPort": 80,
//				"ServerId": "",
//				"Port": 80,
//				"ServerIp": "",
//				"ServerHealthStatus": "normal",
//				"Protocol": "http"
//			},
func (d *DefaultSet) CreateDescribeHealthStatusRequest(LoadBalancerId string) (*CreateDescribeHealthStatusResponse, error) {
	request := slb.CreateDescribeHealthStatusRequest()
	request.Scheme = d.Scheme
	request.LoadBalancerId = LoadBalancerId
	response, err := d.Client.Client.DescribeHealthStatus(request)
	if err != nil {
		fmt.Println("CreateDescribeHealthStatusRequest error:", err.Error())
		return nil, err
	}
	return (*CreateDescribeHealthStatusResponse)(response), nil
}

func (d *DefaultSet) CreateDescribeLoadBalancerAttributeRequest(LoadBalancerId string) (*CreateDescribeLoadBalancerAttributeResponse, error) {
	request := slb.CreateDescribeLoadBalancerAttributeRequest()
	request.Scheme = d.Scheme
	request.LoadBalancerId = LoadBalancerId
	response, err := d.Client.Client.DescribeLoadBalancerAttribute(request)
	if err != nil {
		fmt.Println("CreateDescribeLoadBalancerAttributeRequest:", err.Error())
		return nil, err
	}
	return (*CreateDescribeLoadBalancerAttributeResponse)(response), nil
}

func (d *DefaultSet) CreateDescribeMetricDataRequest() {
	request := cms.CreateDescribeMetricDataRequest()
	request.Scheme = d.Scheme
}
