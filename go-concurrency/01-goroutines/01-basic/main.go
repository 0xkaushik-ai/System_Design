package main

import (
	"fmt"
	//"sync"
	"time"
)

func printMsg() {
	fmt.Println("Hello from go routine")
}

func main() {
	go printMsg()               //  here we called  the function  in go routine and used the keyword go to run the function concurrently ,
	time.Sleep(1 * time.Second) // here we used the sleep function to wait for 1 second before the main function exits, so that the go routine has time to complete its execution

	fmt.Println("Hello from main function") // here we printed the message from the main function
}
