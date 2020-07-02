package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
)

type LoadBalance []slb.LoadBalancer

// SlbGatherData 返回结构图
type SlbGatherData struct {
	Data LoadBalance
}

//func (S *SLBClient) Instances() (*SlbGatherData, error) {
//	slb := &SlbGatherData{}
//	instances, err := S.whereInstances(1, slb)
//	if err != nil {
//		return nil, err
//	}
//	return instances, nil
//}

//func (S *SLBClient) whereInstances(number int, slb *SlbGatherData) (*SlbGatherData, error) {
//	// E.SetPage("PageNumber", number)
//	S.SetParam("PageNumber", number)
//
//	current, err := S.Connect.DescribeLoadBalancers(S.Request)
//	if err != nil {
//		fmt.Print(err.Error())
//		return nil, err
//	}
//	slb.Data = current.LoadBalancers.LoadBalancer
//	pages := int(math.Ceil(float64(current.TotalCount) / float64(current.PageSize)))
//	if pages > 1 {
//		for PageNumber := current.PageNumber + 1; PageNumber <= pages; PageNumber++ {
//			// E.SetPage("PageNumber", requests.NewInteger(PageNumber))
//			S.SetParam("PageNumber", PageNumber)
//			response, err := S.Connect.DescribeLoadBalancers(S.Request)
//			if err != nil {
//				fmt.Print(err.Error())
//				return nil, err
//			}
//			for _, d := range response.LoadBalancers.LoadBalancer {
//				slb.Data = append(slb.Data, d)
//			}
//		}
//	}
//	return slb, nil
//}
