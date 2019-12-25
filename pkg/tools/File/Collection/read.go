package Collection

import (
	"bufio"
	"os"
)

func ReadLine(fileName string) []string {
	list := make([]string,0)
	if file, err := os.Open(fileName);err !=nil{
		panic(err)
	}else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan(){
			list = append(list, string(scanner.Text()))
		}
	}
	return list
}