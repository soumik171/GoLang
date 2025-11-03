package main

import (
	"fmt"
	"maps"
)

// maps->hash, onject, dict

// 2 ways:
// 1.Initialize or uninitialize [when we know the values, can also add later]
// 2.Using make function, just declare [when we don't know the value, add elements later]

func main() {
	// creating map

	m := make(map[string]string)

	// setting an element

	m["name"] = "golang"
	m["area"] = "backend"

	// get element

	fmt.Println(m["name"], m["area"])

	// If key doesnot exist in map, then it return zero value

	fmt.Println(m["phone"]) // return ""

	println("---")

	// For integer:

	n := make(map[string]int)
	n["age"] = 30
	n["price"] = 50

	fmt.Println(n["age"], n["phone"], n["price"])
	fmt.Println(len(n))

	// Delete map:

	delete(n, "price")
	fmt.Println(n)
	fmt.Println(len(n))

	// clear map:

	clear(n)
	fmt.Println(n)

	println("---")

	// Initialize at first :

	mapWOMake := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
	}
	mapWOMake["Four"] = 4

	fmt.Println(mapWOMake)

	// Checking element is present or not:
	// format: return type, bool(true,false):=mapName[key that have to check]

	r, ok := mapWOMake["Two"] // if don't use r, then no need to mention that

	if ok {
		fmt.Println("all ok", r)
	} else {
		fmt.Println("not ok")
	}

	println("---")

	// Equality check of map:

	m1 := map[string]int{"price": 40, "phones": 50}
	m2 := map[string]int{"price": 40, "phones": 50}

	fmt.Println("Is it true that m1 & m2 is true? ", maps.Equal(m1, m2))

}
