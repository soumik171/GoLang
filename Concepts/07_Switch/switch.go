package main

import (
	"fmt"
	"time"
)

func main() {

	// Basic Switch:
	i := 5

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Default value")
	}
	println("---")

	// Multiple Condition switch:

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's workday")
	}
	println("---")

	// Type-Switch

	whoAmI := func(i interface{}) { // interface{}-> it can take any typed values or, we can write whoAmI:=func(i any)
		switch i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a boolean")
		default:
			fmt.Println("Other")
		}
	}

	whoAmI(12)
	whoAmI(55.3)
	whoAmI("soumik")

}
