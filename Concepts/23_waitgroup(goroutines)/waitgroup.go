package main

import (
	"fmt"
	"sync"
)

// waitgroup:run concurrently but print when all are done

func display(i int, w *sync.WaitGroup) {
	defer w.Done() //defer->run at the end of the function, wg.Done()->decrement the waitgroup counter by 1
	fmt.Println("doing it for waitgroup", i)
}
func main() {
	var wg sync.WaitGroup // creates wg of type sync.Waitgroup,keep track how many goroutines are running, so that main can wait for them to finish

	for i := 0; i <= 10; i++ {
		wg.Add(1)          // Increment the counter by 1, then have to call Done() to considered complete
		go display(i, &wg) // starts the new goroutine
	}

	wg.Wait() // wait until waitgroup gets zero

}
