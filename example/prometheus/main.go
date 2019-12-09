package main

import (
	"fmt"
	"github.com/tonyjia87/tractorFrame/pkg/tools/Guzzle"
	"time"
)

// Health can be used to query the Health endpoints


//// Health returns a handle to the health endpoints
//func (c *Guzzle.Client) Prometheus() *Prometheus {
//	return &Prometheus{c}
//}

func (p *Prometheus) UP() {
	fmt.Println("here")
}

//func (c *Guzzle.Client) Prometheus() *Prometheus {
//	return &Prometheus{c}
//}



func main() {
	profile := &Guzzle.Config{
		Address: "promethues-vpc.zmops.cc",
		Scheme: "http",
	}

	client, _ := Guzzle.NewClient(profile)

	pc := &Prometheus{client}
	rtt, rsp , err := pc.Node()
	fmt.Printf("%+v\n",rtt)
	fmt.Printf("%+v\n",*rsp)
	fmt.Printf("%+v\n",err)


}

type Prometheus struct {
	c  *Guzzle.Client
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

func (p *Prometheus) Node() (time.Duration, *RspPrometheus, error, ){
	r := p.c.DoNewRequest("GET","/api/v1/query")
	r.SetParam("Query","query","up{instance='172.17.121.128:9100',job='consul'}")
	rtt, resp, err := Guzzle.RequireOK(p.c.NewDoRequest(r))
	defer resp.Body.Close()
	fmt.Println(rtt,resp, err)
	out := &RspPrometheus{}
	if err := Guzzle.DecodeBody(resp, &out); err != nil {
		return rtt, nil, err
	}

	return rtt, out, nil

}

