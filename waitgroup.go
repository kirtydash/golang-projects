package main

import (
	"fmt"
	"sync"
)

/*

Demonstation of how waitgroups work 

*/


func fetchValuesFromChannel(ch chan []int,wg *sync.WaitGroup) {
	fmt.Println(<-ch)
	wg.Done()
}

func main(){

	sliceChannel := make(chan []int)
	a := []int{1,2,3,4,5}

	var wg sync.WaitGroup

	wg.Add(1)

	go func(){
		defer close(sliceChannel)
		sliceChannel <- a
	}()

	go fetchValuesFromChannel(sliceChannel,&wg)
	
	wg.Wait()
}
