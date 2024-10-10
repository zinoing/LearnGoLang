package main

import (
	"fmt"

	"github.com/zinoing/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("zino")

	fmt.Println(*account)
}
