package main

import (
	"fmt"

	"github.com/zinoing/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("zino")
	account.Deposit(100)

	account.ChangeOwner("nico")

	fmt.Println(account.String())
}
