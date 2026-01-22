package main

import "fmt"

func main(){
	msgChan := make(chan string, 2)  //by default channels have a property that once data is sent to a channel by a goroutine it must be received by another goroutine. This is the unbuffered concept of channels. when channels are buffered a specific number of data items can be sent onto the channel without having a corresponding receiver for them.

	// go func(){
	// 	msgChan <- "Hello"
	// }()

	// go func(){
	// 	msgChan <- "World"
	// }()
	msgChan <- "Hello"
	msgChan <- "World"

	readVal1 := <- msgChan
	readVal2 := <- msgChan

	fmt.Println(readVal1)
	fmt.Println(readVal2)
}