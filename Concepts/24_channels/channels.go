package main

import "fmt"

// channel:used to communicate between the goroutines, send data from one to another

// ele<-10  :: refers send to ele
// <-ele    :: refers received from ele

// Receive data from main:

// func processNum1(numChan chan int) {
// 	fmt.Println("processing number", <-numChan)
// }

// receive data through loop from main: if use loop then no need to use the arrow sign

// func processNum2(numChanL chan int) {
// 	for num := range numChanL {
// 		fmt.Println("Processing number(loop)", num)
// 		time.Sleep(time.Second) // dealy 1 s for better visibility
// 	}
// }

// data send to main:

// func sum(result chan int, num1 int, num2 int) {
// 	newRes := num1 + num2
// 	result <- newRes
// }

// goroutines synchronizer: waitgroup alternatives[when just deal with 1 goRoutines, then can use channel but when we have to use multiple doRoutines, then it's better to use waitGroup cz, we can manually add, done, sleep each goroutines]

func test(done chan bool) {
	defer func() { done <- true }()

	fmt.Println("processing..")
}

func main() {

	// chanMessage := make(chan string) // create channel
	// chanMessage <- "abc" // set goRoutines in channels
	// <-chanMessage // receive goRoutines from channels
	// fmt.Println(chanMessage) //at that time, issue is found deadlock

	// data send from main:

	// numChanN := make(chan int)

	// go processNum1(numChanN)

	// numChanN <- 5

	// println("--Pass loop with delay--")

	// normlly not pass a single data,  pass as a loop and also receive through the loop

	// data send from main:

	// numChanL := make(chan int)

	// go processNum2(numChanL)

	// for {
	// 	numChanL <- rand.Intn(100)
	// }

	// receive data from function:

	// result := make(chan int)

	// go sum(result, 5, 8)

	// receiveRes := <-result

	// fmt.Println(receiveRes)

	done := make(chan bool)
	go test(done)

	<-done // block, here no need to use the sleep

}
