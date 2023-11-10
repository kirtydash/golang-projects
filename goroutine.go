package main


import (
	"fmt"
)

/*

simple function illustrating how channels work

*/

func fetchValuesFromChannel(ch chan []int){
	fmt.Println(<-ch)
}


func main(){

	slicechannel := make(chan []int)
	values := []int{1,2,3,4,5}

	go func(){
	defer close(slicechannel)
        slicechannel <- values
        }()

	fetchValuesFromChannel(slicechannel)
}
