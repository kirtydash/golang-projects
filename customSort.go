package main

import (
	"fmt"
	"sort"
)

type Descending []string

func (d Descending) Len() int           { return len(d) }
func (d Descending) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d Descending) Less(i, j int) bool { return d[i] > d[j] }

func main() {
	strArray := []string{"apple", "orange", "banana", "grape", "kiwi"}

	sort.Sort(Descending(strArray))

	fmt.Println("Sorted Array of Strings (Descending):", strArray)
}

