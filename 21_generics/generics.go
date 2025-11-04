package main

import "fmt"

// Generics:removes redundancy, improve efficiency
// can use any/interface{}

// Takes slice as a parameter

// func printSlice[T interface{}](items []T){...} ->same as any

// func printSlice[T any](items []T) {
// 	for _, elem := range items {
// 		fmt.Println(elem)
// 	}
// }

// If try to fix certain type, not all types:

// func printSlice[T int | string | bool](items []T) {
// 	for _, elem := range items {
// 		fmt.Println(elem)
// 	}
// }

// If more than certain types: then use "comparable"

// func printSlice[T comparable](items []T) {
// 	for _, elem := range items {
// 		fmt.Println(elem)
// 	}
// }

// Can take multiple generic types:

func printSlice[T comparable, V string](items []T, names V) {
	for _, elem := range items {
		fmt.Println(elem, names)
	}

}

// Can use generic in structs:
type genStruct[T any] struct {
	elements []T
}

func main() {

	printSlice([]int{1, 2, 3}, "number")

	printSlice([]string{"js", "ts", "go"}, "language")

	myGenStruct := genStruct[int]{
		elements: []int{5},
		// elements: []string{"golang"},
	}

	fmt.Println(myGenStruct)

}
