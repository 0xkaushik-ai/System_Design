package main

import (
	"fmt"
	"sync"
)

func printMsg(message string, wg *sync.WaitGroup) {
	defer wg.Done() // here we used the defer keyword to call the Done method of the WaitGroup when the function exits, so that the WaitGroup knows that this go routine has completed its execution
	fmt.Println("Hello from go routine: " + message)
}

func main() {

	var wg sync.WaitGroup // here we created a WaitGroup variable to keep track of the number of go routines that are running
	wg.Add(3)             // here we added 3 to the WaitGroup to indicate that we are going to run 3 go routines

	go printMsg("Message 1", &wg) // here we called the function in go routine and used the keyword go to run the function concurrently , and passed the WaitGroup variable as a pointer to the function
	go printMsg("Message 2", &wg) // here we called the function in go routine and used the keyword go to run the function concurrently , and passed the WaitGroup variable as a pointer to the function
	go printMsg("Message 3", &wg) // here we called the function in go routine and used the keyword go to run the function concurrently , and passed the WaitGroup variable as a pointer to the function

	wg.Wait() // here we used the Wait method of the WaitGroup to wait for all the go routines to complete their execution before exiting the main function
	//wg.Done() // here we used the Done method of the WaitGroup to indicate that the main function has completed its execution and the WaitGroup can be decremented by 1
	fmt.Println("Hello from main function")

}

// Here is the whole concept of using all three concept is to manage the lifecycle of multiple go routines and ensure that the main function waits for all of them to complete before exiting.
