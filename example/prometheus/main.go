package main

import (
	"fmt"
	"github.com/tonyjia87/tractorFrame/pkg/tools/Guzzle"
	"github.com/tonyjia87/tractorFrame/pkg/tools/Prometheus/health"
)

func main() {
	profile := &Guzzle.Config{
		Address: "promethues-vpc.zmops.cc",
		Scheme:  "http",
	}

	client, _ := Guzzle.NewClient(profile)
	pc := health.Prometheus{client}

	if err := pc.Health("172.17.121.128"); err != true {
		fmt.Println(err)
	}

}
