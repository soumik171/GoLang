package main

import "fmt"

func main() {

	i := 5

	// can also pass the value, if previously declared, no need to declare again.

	// But we can use that in the function as a parameter

	func() {
		fmt.Println("In-line functions 1:", i)
	}()

	func(i int) {
		fmt.Println("In-line functions 2:", i)
	}(i)

}
