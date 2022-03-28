package main

import (
	"fmt"
)

var list = []int{1, 2, 3, 5, 8}

func main() {
	fmt.Println("Hello, World!\n Here's a list")
	for item := range list {
		fmt.Print(list[item], " ")
	}
	fmt.Print("\n")
}
