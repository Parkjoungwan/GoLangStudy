package main

import (
	"fmt"

	"github.com/park/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("park")
	fmt.Println(account)
}
