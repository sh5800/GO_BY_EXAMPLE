package main

import (
	"fmt"
	"time"
)

func f(s string){
	for i := range 3{
		fmt.Println("From: ",s, " value:",i)
	}
}

func main(){
	f("Direct")

	go f("goroutine")  //goroutines are light weight threads that can execute. The main function in itself is a goroutine.
	                   //goroutines execute concurrently i.e one after the other making it impossible to predict in which manner the output of an executing goroutine would be displayed

	go func(msg string){
		fmt.Println(msg)
	}("Going")

	time.Sleep(time.Second)  //to wait for a goroutine to finish we should use a waitgroup,rather than time.Sleep
	fmt.Println("Done")
}