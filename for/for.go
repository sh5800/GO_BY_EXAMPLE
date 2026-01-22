package main

import(
	"fmt"
)

//To make a statement execute multiple times we use the for loop

func main(){
	i := 1
    for i <= 3 {     //resembles while loop in other OOP languages, but here we use the for keyword to check the condition
        fmt.Println(i)
        i = i + 1
    }

    for j := 0; j < 3; j++ {  //a regular for loop with initialization, condition check and increment statement
        fmt.Println(j)
    }


    for i := range 3 {  //with range keyword we can iterate till a given range of numbers.
        fmt.Println("range", i)
    }


    for {   //for loop without initialization and condition check will run forever until terminated explicitly
        fmt.Println("loop")
        break
    }


    for n := range 6 { //continue statements are used to skip an iteration and move to the next iteration value
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}