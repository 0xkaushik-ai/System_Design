// A value is sent into a channel using:
// channel <- value

package main

import "fmt"

func main() {
	ch := make(chan string) // here we made a channel of type string := symbol means "make a channel of type string"

	go func() { // here go means "run this function in a separate goroutine"
		ch <- "Hello, World!"
	}() // here () means "call this function immediately"
	message := <-ch      // here <-ch means "receive a value from the channel ch"
	fmt.Println(message) // here we print the message received from the channel

}

// A channel on only works between go routines inside the same running program . Put both the operation together .
