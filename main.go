package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"Ramen"}
	zino := person{name: "zino", age: 25, favFood: favFood}
	fmt.Println(zino.favFood)
}
