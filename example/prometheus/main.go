package main

import (
	"fmt"
	"github.com/tonyjia87/tractorFrame/pkg/tools/Guzzle"
	"github.com/tonyjia87/tractorFrame/pkg/tools/Prometheus/health"
)

func main() {
	profile := &Guzzle.Config{
		Address: "you address",
		Scheme:  "http",
	}

	client, _ := Guzzle.NewClient(profile)
	pc := health.Prometheus{client}

	if err := pc.Health("you ip "); err != true {
		fmt.Println(err)
	}

}
