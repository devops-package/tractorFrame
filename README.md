# tractorFrame


使用 

```
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/devops-package/tractorFrame/pkg/tools/Guzzle"
	"github.com/devops-package/tractorFrame/pkg/tools/Prometheus/health"
)

func main() {

	profile := &Guzzle.Config{
		Address: "ip",
		Scheme:  "http",
	}

	client, _ := Guzzle.NewClient(profile)
	pc := health.Prometheus{client}

	if err := pc.Health(ip, 9100); err != true {
        fmt.Println(ip, err)
	}

}

```
