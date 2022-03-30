package main

import (
	"fmt"
	"learnpackage/simpleinterest"
	"math/rand"
)

func main() {
	rand.Seed(86)
	fmt.Println("Simple interest calculation")
	p := rand.Float64()*5000.0
	fmt.Println("Principal: ", p)
	r := 10.0
	t := 1.0
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
