package main

import "fmt"

func sum(nums ... int) int{  // A variadic function contains an argument of type slice so that it can accept any number of arguments like func funcName(param1 dataType1,param2 dataType2,..,variadicParam ...dataType){ //function body }
	fmt.Println(nums)
	total := 0

	for _,val := range nums{
		total += val
	}

	return total
}

func main(){
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))

	nums := []int{1,2,3,4}  //if the variadic parameter is a slice then we can use it like this res := funcName(sliceName...)

	fmt.Println(sum(nums...))
}