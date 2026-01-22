package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)   //In Go, the select statement allows you to wait on multiple channel operations, such as sending or receiving values. Similar to a switch statement, select enables you to proceed with whichever channel case is ready, making it ideal for handling asynchronous tasks efficiently.
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for range 2 {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}