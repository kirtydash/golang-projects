package main

import "fmt"

func main() {
    a := make([]int, 2, 3)

    if len(a) > 3 {
        if a[3] != 0 {
            panic("A try block")
        }
    }

    defer func() {
        fmt.Println("Exit gracefully")
    }()
}

