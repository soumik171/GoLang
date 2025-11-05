package main

import (
	"fmt"
	"time"
)

// Buffered channel: If need to send data 1 by 1(send,receive), then can use normal channel, but if have to deal with multiple  data without blocking, then have to use buffered channel

// Here just visible, one channel(either send or receive) can work at a time
// If wanna see multiple channel both send or both received at a time, then visit 25.1_ad_buffer

// A channel can both send and receive data like:

// func emailsender(emailChan chan string, done chan bool) {...

// done<-true
// <-done

// }

// but normally we define that with at the paramterer of the function, like:

func emailsender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second) // normally reduce the process in function
	}
}

func main() {
	emailChan := make(chan string, 100) //as here define 100, that means can run upto 100 without blocking

	done := make(chan bool)

	// Can check manually, but mainly use loop to send & receive

	// emailChan <- "1@email.com"
	// emailChan <- "2@email.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	for i := 1; i <= 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	go emailsender(emailChan, done)

	fmt.Println("done sending")

	//close channel:important to close
	close(emailChan)

	<-done

}
