package main

import "fmt"

func nextVal() func() int{
	i := 0
	return func() int{
		i++
		return i
	}
}

//other than primitives that can be assigned to a variable, functions can also be assigned to a variable like this funcName := func(parameters dataTypes) (returnTypes){ //function body }

func main(){
	nextInt := func(i int) int{
		return i + 1
	}

	fmt.Println(nextInt(0))
	fmt.Println(nextInt(1))
	fmt.Println(nextInt(2))

	upcomingVal := nextVal()
	fmt.Println(upcomingVal()) //0
	fmt.Println(upcomingVal()) //1
	fmt.Println(upcomingVal()) //2
}