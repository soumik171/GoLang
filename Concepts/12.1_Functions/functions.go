package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// if multiple values are same type can declare last at once
func addS(a, b int) int {
	return a + b
}

// can return multiple values at once

func getLanguages() (string, string, bool) {
	return "GoLang", "JavaScript", true
}

// Function receives function as a parameter:

func functionReceivesFunction(fn func(a int) int) int {
	return fn(5)
}

// Function returns function:

func functionReturnsFunction() func(a int) int {
	return func(a int) int {
		return a
	}

}

func main() {

	// store in result then call
	result := add(5, 3)
	fmt.Println(result)

	// direct call
	fmt.Println(addS(6, 8))

	// store into diff variables[if anything get but not want to return then go for _]
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)

	// receives function as a parameter
	fn1 := func(a int) int {
		return 2 * a
	}
	fmt.Println(functionReceivesFunction(fn1))

	// return function
	fn2 := functionReturnsFunction() //fn2 receives function
	fmt.Println(fn2(3))

}
