package main

import (
	"fmt"
)

// Declare outside the function, if use shorthand, then it's not possible outside the function

func main() {
	// If declare any variables, then must have ti use that or it shows errors

	// If type is not defined then auto define that

	// string

	// var name string = "golang"
	var name = "string"

	fmt.Println(name)

	// int

	var age = 25
	fmt.Println(age)

	// bool

	var isNum = true
	fmt.Println(isNum)

	// shorthand:declare with only name

	myName := "Soumik"
	fmt.Println(myName)

	// Some times, we need to declare first, then use them, at that cases

	var ammount float32
	ammount = 25

	fmt.Println(ammount)

}
