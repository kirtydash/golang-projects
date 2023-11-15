package main


import (
	"fmt"
)

func main(){

	var anything interface{}

	anything = 12

	if val,ok := anything.(int);ok{
		fmt.Println("everything well;",val)
	}

	anything = "kirty"

	if val,ok := anything.(string);ok {

		fmt.Println("string type here",val)
	}
}
