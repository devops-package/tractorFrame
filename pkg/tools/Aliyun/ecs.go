package aliyun

import (
	"fmt"
	"math"
	"regexp"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// EcsGatherData 直接查询后的所有服务器结构图
type EcsGatherData struct {
	Data []ecs.Instance `json:"Instance" xml:"Instance"`
}

// WhereKV 批次条件 目前没有实现
type WhereKV struct {
	mArgs map[string]string
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

// Reset 清理client
func (E *EcsClient) Reset() {
	E.IsInsecure = false
	E.PageNumber = 1
}

// BatchSetParam 批次的方法
func (E *EcsClient) BatchSetParam(args *WhereKV) {
}

// whereInstances 通过填写条件开始拉数据
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

func (E *EcsClient) MatchByName(i *EcsGatherData, r string) *EcsGatherData {
	Matched := &EcsGatherData{}
	for _, d := range i.Data {
		matched, _ := regexp.MatchString(r, d.InstanceName)
		if matched {
			Matched.Data = append(Matched.Data, d)
		}
	}
	return Matched
}

func (E *EcsGatherData) MatchOsType(r string) *EcsGatherData {
	Matched := &EcsGatherData{}
	for _, d := range E.Data {
		matched, _ := regexp.MatchString(r, d.OSType)
		if matched {
			Matched.Data = append(Matched.Data, d)
		}
	}
	return Matched
}

func (E *EcsGatherData) MatchStatus(r string) *EcsGatherData {
	Matched := &EcsGatherData{}
	for _, d := range E.Data {
		matched, _ := regexp.MatchString(r, d.Status)
		if matched {
			Matched.Data = append(Matched.Data, d)
		}
	}
	return Matched
}
