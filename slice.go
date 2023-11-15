package main

import (
	"fmt"
)

func main(){

	//slice with len 5 and capacity 10
	slc := make([]int,5,10)
	//array of size 5 immutable
	slc1 := [5]int{}
	fmt.Println(slc,slc1)
}
