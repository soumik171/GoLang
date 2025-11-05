package main

import "fmt"

func counter() func() int {
	count := 5

	return func() int {
		count++
		return count
	}
}
func main() {

	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())

}
