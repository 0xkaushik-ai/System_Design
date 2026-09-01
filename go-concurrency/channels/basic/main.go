package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string) {
	ch <- "Hello from go routine" // here we sent the message to the channel using the send operator <- , and the message is sent to the channel and the channel is blocked until the message is received by the main function
}
func main() {
	ch := make(chan string)
	go sendMessage(ch) // here we called the function in go routine and used the keyword go to run the function concurrently , and passed the channel variable as a parameter to the function
	message := <-ch    // here we received the message from the channel using the receive operator <- , and the message is received from the channel and the channel is unblocked
	fmt.Println(message)
	time.Sleep(1 * time.Second) // here we used the sleep function to wait for 1 second before the main function exits, so that the go routine has time to complete its execution

	fmt.Println("Hello from main function") // here we printed the message from the main function
}

// channels let go routne communicate with each other and synchronize their execution, and they are used to send and receive messages between go routines.
 ch := make(chan string)

//  Creates a channel that carries strings.

 // ch <- "message"

 // Sends a value into the channel.

  //message := <-ch
	  // Receives a value from the channel and assigns it to the variable message.
	  // The receving operation waits until the go routine sends a message . This means you dont need time.sleep .
	  