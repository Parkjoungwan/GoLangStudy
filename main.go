package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 10
	*b = 20
	fmt.Println(a, *b) //Low-level programing
}
