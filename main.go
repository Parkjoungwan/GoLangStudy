package main

import (
	"fmt"

	"github.com/park/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("Second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
