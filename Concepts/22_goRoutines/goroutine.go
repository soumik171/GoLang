package main

import (
	"fmt"
	"time"
)

// why use goroutine:runs concurrently,lightweight, managed by Go's scheduler
// runs difficult task as a same time where other languages might fail

// just go func()-> run concurrently and exit anytime
func main() {

	for i := 0; i <= 10; i++ {

		go func() {
			fmt.Println("doing it", i)

		}()

	}

	// delay the process for goroutines: but this is not reliable, see in channel & waitgroup

	time.Sleep(time.Second * 2)

}
