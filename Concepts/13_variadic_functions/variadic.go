package main

import "fmt"

// If receive data at as a similar types
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// If receive data at any types, then use interface{}

func takeDiffTypes(val ...interface{}) interface{} {
	return val
}

func main() {

	// pass individually in functions
	result := sum(1, 5, 7, 9, 6, 2)
	fmt.Println(result)

	// pass from slice

	slice := []interface{}{1, 5, "soumik"}
	fmt.Println(takeDiffTypes(slice))

}
