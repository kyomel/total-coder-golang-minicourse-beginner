package main

import "fmt"

// func eatCookie(messageCh chan string) {
// 	result := <-messageCh
// 	fmt.Println(result)
// }

func eatCookies(messageCh chan string) {
	for message := range messageCh {
		fmt.Println("reading from channel:", message)
	}
}

// A channel ALWAYS blocks if its full
// An unbuffered channel is always full
func main() {
	messageCh := make(chan string, 128) // buffered channel
	// messageCh := make(chan string) // unbuffered channel
	// 1. Write to a channel
	// 2. Read from a channel

	// unbuffered channel init
	// go eatCookies(messageCh)

	// for i := 0; i < 100; i++ {
	// 	messageCh <- "Hello" // blocking here
	// }

	// channel with close
	for i := 0; i < 100; i++ {
		messageCh <- fmt.Sprintf("Hello %d\n", i) // blocking here
	}
	close(messageCh)

	eatCookies(messageCh)

	// buffered channel init
	// result := <-messageCh // unblocking here
	// fmt.Println(result)
}
