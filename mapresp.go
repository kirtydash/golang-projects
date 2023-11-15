package main

import (
	"fmt"
)

/**

check if key exists in dict

**/



func main(){

	carDict := map[string]string{
		"model":"maruti",
		"batch":"2.1.2",
		"size":"51",
	}

	if val,ok := carDict["model"]; ok{
            fmt.Println(val)
	}

}
