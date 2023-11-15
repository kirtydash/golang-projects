package main

import "fmt"

func permute(input []rune, start int) {
	if start == len(input)-1 {
		fmt.Println(string(input))
		return
	}

	for i := start; i < len(input); i++ {
		input[start], input[i] = input[i], input[start]
		permute(input, start+1)
		input[start], input[i] = input[i], input[start]
	}
}

func main() {
	str := "abc"
	chars := []rune(str)

	fmt.Println("Permutations of", str, ":")
	permute(chars, 0)
}

