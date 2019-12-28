package aliyun

import (
	"fmt"
	"math"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// EcsGatherData 直接查询后的所有服务器结构图
type EcsGatherData struct {
	Data []ecs.Instance `json:"Instance" xml:"Instance"`
}

func (E *EcsClient) defaultParam() {
	request := E.Request
	request.PageSize = "100"
}

// SetParam 设置请求参数
func (E *EcsClient) SetParam(key string, value interface{}) {
	request := E.Request
	E.defaultParam()
	switch key {
	case "InstanceNetworkType":
		request.InstanceNetworkType = value.(string)
	case "PageNumber":
		val := value.(int)
		request.PageNumber = requests.NewInteger(val)
	}
}

// whereInstances 从第n页开始拉数据
func (E *EcsClient) whereInstances(number int, ecs *EcsGatherData) (*EcsGatherData, error) {
	// E.SetPage("PageNumber", number)
	E.SetParam("PageNumber", number)
	current, err := E.Connect.DescribeInstances(E.Request)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
	ecs.Data = current.Instances.Instance
	pages := int(math.Ceil(float64(current.TotalCount) / float64(current.PageSize)))
	if pages > 1 {
		for PageNumber := current.PageNumber + 1; PageNumber <= pages; PageNumber++ {
			// E.SetPage("PageNumber", requests.NewInteger(PageNumber))
			E.SetParam("PageNumber", PageNumber)
			response, err := E.Connect.DescribeInstances(E.Request)
			if err != nil {
				fmt.Print(err.Error())
				return nil, err
			}
			for _, d := range response.Instances.Instance {
				ecs.Data = append(ecs.Data, d)
			}
		}
	}
	return ecs, nil
	// E.Response = response

}

// Instances 所有ecs
func (E *EcsClient) Instances() (*EcsGatherData, error) {
	ecs := &EcsGatherData{}
	instances, err := E.whereInstances(1, ecs)
	if err != nil {
		return nil, err
	}
	return instances, nil
}
