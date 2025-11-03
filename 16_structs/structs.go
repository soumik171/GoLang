package main

import (
	"fmt"
	"time"
)

// As class is not available in go, we can use struct instead of class
// Can also use for "the collection of datatypes"
// Struct is like a blueprint
// From struct, we can create multiple instances

// If you don't set any feild,default value is zero
// int=>0,float=>0,string=>"",bool=>false

// order struct

type order struct {
	id        string
	ammount   float32
	status    string
	createdAt time.Time // nanosecond precision
}

// If pass the struct in a function(receiver type)
// convention: consider the starting letter of the struct

// When change anything, must use pointer but no need to 
func (o *order) changeStatus(status string) {
	o.status = status
}

// when just get the values, no need to use the pointer, if use, then also it's ok
func (o *order) getAmmount() float32 {
	return o.ammount
}

func main() {

	myOrder := order{
		id:      "171",
		ammount: 50.00,
		status:  "received",
	}
	fmt.Println("Order Struct", myOrder)

	// add value after
	myOrder.createdAt = time.Now()

	fmt.Println("Order Struct", myOrder)

	// get feilds:

	fmt.Println("Ammount is", myOrder.ammount)

	println("--Create 2nd instance--")
	// Created a 2nd instance:

	myOrder2 := order{
		id:        "2",
		ammount:   100,
		status:    "delivered",
		createdAt: time.Now(),
	}
	fmt.Println(myOrder2)

	myOrder2.status = "paid"

	fmt.Println(myOrder2)

	println("--Pass struct into function--")
	// pass struct into function
	myOrder2.changeStatus("confirmed")
	fmt.Println(myOrder2)

	fmt.Println(myOrder2.getAmmount())

	println("--Shorthand--")
	// If sometimes just use one time struct, then can use shorthand:

	about := struct {
		id     int
		name   string
		goodAt string
	}{12, "Soumik", "JS"}

	fmt.Println(about)

}
