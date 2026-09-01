package main

import (
	"fmt"
	//"sync"
	"time"
)

func printMsg(message string) {
	fmt.Println("Hello from go routine: " + message)
}

func main() {
	go printMsg("Message 1") //  here we called  the function  in go routine and used the keyword go to run the function concurrently ,
	go printMsg("Message 2") //  here we called  the function  in go routine and used the keyword go to run the function concurrently ,
	go printMsg("Message 3") //  here we called  the function  in go routine and used the keyword go to run the function concurrently ,

	time.Sleep(1 * time.Second) // here we used the sleep function to wait for 1 second before the main function exits, so that the go routine has time to complete its execution

	fmt.Println("Hello from main function") // here we printed the message from the main function
}

// Here the order of the go routine output may change each time . THis is because the go schedular decides when each go routine will run and the order of execution is not guaranteed.
