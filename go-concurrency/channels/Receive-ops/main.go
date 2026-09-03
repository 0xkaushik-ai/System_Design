package main

import "fmt"

func sendMessage(ch chan<- string) {
	ch <- "Hello from send operation!" // send a message to the channel
}

func receiveMessage(ch <-chan string) {
	message := <-ch // receive a message from the channel
	fmt.Println(message)
}

func producer(ch chan<- string) {
	ch <- "Message  1"
	ch <- "Message 2"
	close(ch) // this closes the channel after sending all messages, so that the receiver knows there are no more messages to receive
}

func main() {
	// we have received a value from a Send operation using the receive operator <-. The receive operator is used to receive a value from a channel. The syntax for receiving a value from a channel is:
	// <-channel

	ch := make(chan string)

	go sendMessage(ch) // start a goroutine to send a message
	receiveMessage(ch) // receive the message in the main goroutine

	// now receive multiple messages from the producer function
	// here we called the function producer in a separate goroutine and passed the channel ch2 to it. The producer function sends multiple messages to the channel and then closes it. In the main function, we use a range loop to receive messages from the channel until it is closed. The range loop will exit when it reaches the end of the channel, which is indicated by the close operation in the producer function.
	ch2 := make(chan string)
	go producer(ch2)

	for msg := range ch2 { // receive messages until the channel is closed , here range  keyword is used to iterate over the values received from the channel until it is closed
		fmt.Println(msg)
	} // := means "receive a value from the channel and assign it to the variable msg" and range means "iterate over the values received from the channel until it is closed" but how its know ths value is closed ?  the channel is closed in the producer function after sending all messages, so the range loop will exit when it reaches the end of the channel.

}
