
## usage

## 搜索ip
```shell script
	
package main

import (
	"fmt"
	"github.com/devops-package/tractorFrame/pkg/tools/CD"
	"github.com/devops-package/tractorFrame/pkg/tools/Guzzle"
	"log"
)



func main()  {
	CMDBAddress := "cmdb02.zmops.cc"
	CDProfile := &Guzzle.Config{
		Address: CMDBAddress,
		Scheme:  "http",
	}
	CDHandler, err := Guzzle.NewClient(CDProfile)
	if err != nil {
		log.Fatal(err)
	}
	CDClient := CD.CDGuzzle{CDHandler}

	o, _ := CDClient.ServerSearch("172.19.19.791", true)
	fmt.Println(o)
}
```
