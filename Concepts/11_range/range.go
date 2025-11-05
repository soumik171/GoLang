package main

import "fmt"

func main() {
	nums := []int{5, 6, 7}

	// if use for:
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	println("--Print only element--")

	// if using range:

	for _, elem := range nums { // here elem refers the elements of the slices
		fmt.Println(elem)
	}

	println("--Print Index--")

	// can use the index also:

	for i, val := range nums {
		fmt.Println(val, i)
	}

	println("--Iter map(key,value) pair--")

	// Can also iterate map as well:

	m := map[string]string{"fruit": "apple", "flower": "lily"}

	for k, v := range m {
		fmt.Println(k, v)

	}

	println("--Only key--")

	// if req only key, then also use that but convention is to use key & value at once

	for k := range m {
		fmt.Println(k)
	}

	println("--Range applied in string--")

	// Rnage in string:
	// i refers starting byte of the rune
	// c refers unicode/code point rune

	for i, c := range "golang" {
		fmt.Println(i, c)
	}

	println("--Print with character---")

	// If require string, then use string(c)

	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
