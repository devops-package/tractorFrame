package main

import (
	"bufio"
	"fmt"
	"github.com/tonyjia87/tractorFrame/pkg/tools/Guzzle"
	"github.com/tonyjia87/tractorFrame/pkg/tools/Prometheus/health"
	"os"
)

func main() {


	profile := &Guzzle.Config{
		Address: "promethues-vpc.zmops.cc",
		Scheme:  "http",
	}

	client, _ := Guzzle.NewClient(profile)
	pc := health.Prometheus{client}

	d := ReadLineFile("/Users/tonyjia/zm/vpc/vpc")
	for _, ip := range d{
		if err := pc.Health(ip); err != true {

			fmt.Println(ip, err)
		}
	}

}


func ReadLineFile(fileName string) []string {
	list := make([]string,0)
	if file, err := os.Open(fileName);err !=nil{
		panic(err)
	}else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan(){
			//fmt.Println(scanner.Text())
			list = append(list, string(scanner.Text()))
		}
	}
	return list
}
