package Consul

import (
	"fmt"
	"github.com/tonyjia87/tractorFrame/pkg/tools/Guzzle"
)

type StateStruct []struct {
	Node        string   `json:"Node"`
	CheckID     string   `json:"CheckID"`
	Name        string   `json:"Name"`
	Status      string   `json:"Status"`
	Notes       string   `json:"Notes"`
	Output      string   `json:"Output"`
	ServiceID   string   `json:"ServiceID"`
	ServiceName string   `json:"ServiceName"`
	ServiceTags []string `json:"ServiceTags"`
	Definition  struct {
	} `json:"Definition"`
	CreateIndex int `json:"CreateIndex"`
	ModifyIndex int `json:"ModifyIndex"`
}

func (c *Consul) state(state string) (*StateStruct, error) {
	url := fmt.Sprintf("/v1/health/state/%s", state)
	r := c.DoNewRequest("GET", url)
	out := Guzzle.RequireOK(c.NewDoRequest(r))

	f := &StateStruct{}
	if err := out.Json(f); err != nil {
		return nil, err
	}

	return f, nil
}

func (c *Consul) State(state string, result *StateStruct) *StateStruct {
	out, err := c.state(state)
	if err != nil {
		return nil
	}
	result = out
	return result
}
