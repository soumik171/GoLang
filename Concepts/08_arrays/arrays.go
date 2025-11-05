package main

import (
	"fmt"
)

// Numbered sequence of specific length
// If not declared first, then bydefault int->0, string->"", bool->false

// When to use:
// -fixed size, that is predictable
// -Memory Optimization
// -Constant time access

func main() {
	var nums [4]int

	// return length
	fmt.Println(len(nums))

	// set & return values
	nums[0] = 1
	fmt.Println(nums[0])

	// return array
	fmt.Println(nums)

	println("---")

	// For boolean:

	var vals [5]bool
	vals[2] = true
	fmt.Println(vals)

	println("---")

	// For String:

	var names [3]string
	names[1] = "soumik"
	fmt.Println(names)

	println("---")

	// Declared in a single line:

	values := [3]int{1, 2, 3}

	fmt.Println(values)

	println("---")

	// 2d arrays
	nums2d := [2][2]int{{3, 4}, {5, 6}}

	fmt.Println((nums2d))
}
