package main

import "fmt"

type order struct {
	id      int
	name    string
	address string
	isGood  bool
}

// Although constructor is not present in go, but in a bit diff way, we can inittialize them to apply OOP concept

func newOrder(id int, name string, address string, isGood bool) *order { //as change so use pointer
	myOrder := order{
		id:      id,
		name:    name,
		address: address,
		isGood:  isGood,
	}

	return &myOrder

}

func main() {

	fmt.Println(newOrder(1, "alif", "Mohammadpur", true))

}
