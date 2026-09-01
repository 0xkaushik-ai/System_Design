package main

import (
	"fmt"
)

func main() {
	messageChannel := make(chan string, 2) // here we created a buffered channel with a capacity of 2, which means that the channel can hold up to 2 messages before it blocks the sender
	messageChannel <- "Hello from go routine 1"
	messageChannel <- "Hello from go routine 2" // here we sent the message to the channel using the send operator <- , and the message is sent to the channel and the channel is blocked until the message is received by the main function
	fmt.Println(<-messageChannel)
	fmt.Println(<-messageChannel) // here we received the message from the channel using the receive operator <- , and the message is received from the channel and the channel is unblocked
}

// buffered channels let go routines communicate with each other and synchronize their execution, and they are used to send and receive messages between go routines. The difference between buffered and unbuffered channels is that buffered channels can hold a certain number of messages before they block the sender, while unbuffered channels block the sender until the message is received by the receiver.
// buffer channel is useful when you want to send multiple messages to a channel without blocking the sender, and you want to receive the messages in a different order than they were sent.
// unbuffered channel is useful when you want to send a message to a channel and wait for the receiver to receive the message before continuing execution.
// in distibuted systems, buffered channels can be used to implement a message queue, where multiple producers can send messages to the queue and multiple consumers can receive messages from the queue. This allows for decoupling of the producers and consumers, and allows for better scalability and fault tolerance.
//
