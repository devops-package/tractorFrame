package health

import (
	"fmt"

	"github.com/tonyjia87/tractorFrame/pkg/tools/Guzzle"
)

type Prometheus struct {
	*Guzzle.Client
}

type RspPrometheus struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Name     string `json:"__name__"`
				Instance string `json:"instance"`
				Job      string `json:"job"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func (p *Prometheus) Health(node string, port int) bool {
	//
	// n := fmt.Sprintf("up{instance='%s:9100',job='consul'}", node)
	n := fmt.Sprintf("up{instance='%s:%d',job='consul'}", node, port)
	r := p.DoNewRequest("GET", "/api/v1/query")
	r.SetParam("Query", "query", n)
	out := Guzzle.RequireOK(p.NewDoRequest(r))
	f := &RspPrometheus{}
	if err := out.Json(f); err != nil {
		return false
	}
	if len(f.Data.Result) > 0 {
		value := f.Data.Result[0].Value[1]
		if value == "1" {
			return true
		}
	}

	return false
}

func (p *Prometheus) JvmHealth(node string, port int) bool {
	//
	// n := fmt.Sprintf("up{instance='%s:9100',job='consul'}", node)
	n := fmt.Sprintf("jmx_exporter_build_info{instance='%s:%d',job='consul'}", node, port)
	r := p.DoNewRequest("GET", "/api/v1/query")
	r.SetParam("Query", "query", n)
	out := Guzzle.RequireOK(p.NewDoRequest(r))
	f := &RspPrometheus{}
	if err := out.Json(f); err != nil {
		return false
	}
	if len(f.Data.Result) > 0 {
		value := f.Data.Result[0].Value[1]
		if value == "1" {
			return true
		}
	}

	return false
}
