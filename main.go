package main

import (
	"fmt"

	mydict "github.com/zinoing/learngo/dict"
)

func main() {
	dictionary := mydict.Dictionary{}
	err := dictionary.Add("first", "First word")
	if err != nil {
		fmt.Println(err)
	}

	err = dictionary.Delete("first")
	if err != nil {
		fmt.Println(err)
	}

	def, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		println(def)
	}
}
