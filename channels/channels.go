package main

import "fmt"

func main(){

	msgChan := make(chan string)   //channels in go are named pipes used to pass data between goroutines. To declare a channel we use chanName := make(chan chanType)

	go func(){
		msgChan <- "Hello"   // to send data to a channel we use the syntax chanName <- data
	}()

	message := <- msgChan   // to receive data from a channel we use the syntax varName := <- chanName
	fmt.Println(message)    //here data is being passed from the anonymous goroutine to the func main goroutine.
}