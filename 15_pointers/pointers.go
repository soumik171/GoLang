package main

import "fmt"

// when we pass the value,it creates a copy and we use them, it doesn't distinctly change the value

func changeNumbyVal(num int) {
	num = 5
	fmt.Println("In changeNum by value", num)

}
func changeNumbyRef(num *int) {
	*num = 5
	fmt.Println("In changeNum by reference", *num)

}
func main() {
	num := 1

	changeNumbyVal(num)

	fmt.Println("After changing by value in main:", num)

	changeNumbyRef(&num)

	fmt.Println("After changing by reference in main", num)
}
