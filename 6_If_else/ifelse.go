package main

import "fmt"

// Ternary is not supported in go as far as i learn
func main() {

	// if-else

	age := 16

	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Under-aged")
	}

	println("---")

	// if-else-if

	number := 15

	if number >= 20 {
		fmt.Println("Greater than 20")
	} else if number > 10 && number < 20 {
		fmt.Println("In between 10 to 20")
	} else {
		fmt.Println("Below 10")
	}

	println("---")

	var role = "admin"
	var hasPermission = true

	if role == "admin" || hasPermission {
		fmt.Println("yes")
	}
	println("---")

	if age := 15; age > 18 {
		fmt.Println("Adult", age)
	} else if age >= 12 {
		fmt.Println("Teenager", age)
	}
	println("---")

}
