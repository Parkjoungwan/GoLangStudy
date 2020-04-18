package main

import (
	"fmt"

	"github.com/park/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("park")
	account.Deposit(10)

	fmt.Println(account)

}
