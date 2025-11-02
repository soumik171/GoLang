package main

import (
	"fmt"
)

// for->only construct  in go for looping
func main() {
	// If have to declare while loop(As it's not present ni go)

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	println("---")

	// If need infinite loop:

	// for {
	// 	println("1")
	// 	fmt.Println(2)
	// }

	// classic for loop

	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	println("---")

	// can also use break,continue

	for i := 0; i <= 4; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
	println("---")

	for i := 0; i <= 3; i++ {

		fmt.Println(i)
		break
	}
	println("---")

	// range(removes the last term)

	for i := range 5 {
		fmt.Println(i)
	}

}
