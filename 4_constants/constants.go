package main

import "fmt"

// Declare outside the function, if use shorthand, then it's not possible outside the function

func main() {
	const name string = "GoLang"
	const age = 20

	// If sometimes use multiple const at a time

	const (
		port = 500
		host = "localhost"
	)

	fmt.Println(port, host)

}
