package Consul

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/tonyjia87/tractorFrame/pkg/tools/Guzzle"
)

type Consul struct {
	*Guzzle.Client
}

type Check struct {
	HTTP     string `json:"http"`
	Interval string `json:"interval"`
	Timeout  string `json:"timeout"`
}
type Exporter struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Port    int      `json:"port"`
	Tags    []string `json:"tags"`
	Checks  []Check
}

func genBody(exporter, node string, port int, tags []string, interval string, timeout string) Exporter {

	body := Exporter{
		ID:      fmt.Sprintf("%s_%d", node, port),
		Name:    exporter,
		Address: node,
		Port:    port,
		Tags:    tags,
		Checks: []Check{
			{HTTP: "HTTP", Interval: interval, Timeout: timeout},
		},
	}

	return body
}

func (c *Consul) GenBody(exporter, node string, port int, tags []string, interval string, timeout string) *bytes.Buffer {
	body := genBody(exporter, node, port, tags, interval, timeout)
	json, _ := json.Marshal(body)
	return bytes.NewBuffer(json)
}

//
func (c *Consul) Register(exporter, node string, port int) (string, bool) {
	r := c.DoNewRequest("PUT", "/v1/agent/service/register")
	tag := make([]string, 1)
	tag[0] = "pro"
	body := c.GenBody(exporter, node, port, tag, "5s", "5s")
	r.SetBody(body)
	out := Guzzle.RequireOK(c.NewDoRequest(r))
	defer out.Close()
	if out.StatusCode == 400 {
		return fmt.Sprintf("%s register %s fales %d\n", node, exporter, out.StatusCode), false
	}
	return fmt.Sprintf("%s server %s registered %d\n", node, exporter, out.StatusCode), true
}
