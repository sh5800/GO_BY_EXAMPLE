package main

import(
	"fmt"
)

func sum(x,y int) int{   //function declaration func funcName(parameters dataTypes) returnType{ //function body }
	return x + y
}

func sum3(x,y,z int)int{
	return x + y + z
}

func main(){
	fmt.Println(sum(10,20))  //call the function with actual values like result := functionName(val1,val2,val3 ......)
	fmt.Println(sum3(10,20,30))
}