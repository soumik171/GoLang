package main

import (
	"fmt"
	"slices"
)

// when we don't know the actual size of the array, then we can use that
// It's dynamic
// most use construct in go
// useful methods

func main() {

	// Uninitilaize slice is null
	var name []int

	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println(name == nil)

	println("---")

	// 1st approch:

	nums := []int{}

	// Add elements type1:
	nums = append(nums, 5)
	nums = append(nums, 2)

	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))

	println("---")

	// 2nd approch:using make function, take 2 to 3 parameters
	// ...=make([]int,length/capacity)
	// ...=make([]int,length, capacity) if length exceeds, then capacity automatically doubles

	var slicess = make([]int, 2, 5)

	// Add elements type 2:
	slicess[0] = 3
	slicess[1] = 2

	slicess = append(slicess, 12)

	fmt.Println((slicess))
	fmt.Println((len(slicess)))
	fmt.Println((cap(slicess)))

	println("---")

	// Copy function:

	var copy1 = make([]int, 0, 5)
	var copy2 = make([]int, 2)

	copy1 = append(copy1, 3)

	copy(copy2, copy1)

	fmt.Println(copy1, copy2)

	println("---")

	// slice operator:[a,b] ->[a->b)-> a<=slice<b

	var sOp = []int{1, 2, 3}

	fmt.Println(sOp[0:2]) // [1,2]
	fmt.Println(sOp[:1])  // [1] if starting not mentioned, then go from the start
	fmt.Println(sOp[1:3]) // [1,2]
	fmt.Println(sOp[1:])  //[2,3] start from 1 index

	println("---")

	// Equality check:

	var equalityCheck1 = []int{1, 2, 3}
	var equalityCheck2 = []int{1, 2, 3}

	fmt.Println(slices.Equal(equalityCheck1, equalityCheck2))

	// 2d slices:

	var Slices2d = [][]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println(Slices2d)

}
