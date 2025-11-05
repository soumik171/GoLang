package main

import "fmt"

type customer struct {
	name  string
	phone string
}

// embed customer struct into order struct

type order struct {
	id     int
	name   string
	status string
	isGood bool
	customer
}

func main() {

	currOrder := order{
		id:     1,
		name:   "shapna",
		status: "hw",
		isGood: true,
	}

	fmt.Println(currOrder)
	fmt.Println(currOrder.customer)

	println("--Set customer values--")

	// set customer values:

	newCustomer := customer{
		name:  "AB",
		phone: "12547893",
	}

	newOrder := order{
		id:       5,
		name:     "soumik",
		status:   "apl",
		isGood:   false,
		customer: newCustomer,
	}

	fmt.Println(newOrder.customer)
	fmt.Println(newOrder)

	println("--Change the value of the attributes--")

	// change the attributes of struct:

	newOrder.customer.phone = "Srabanty"

	fmt.Println(newOrder.customer)
	fmt.Println(newOrder)

}
