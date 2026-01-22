package main

import "fmt"

// a function in go can return multiple values at once like func funcName(parameter1 dataType1, parameter2 dataType2, ....) (returnType1, returnType2,....){
	//return returnType1, returnType2,.....}
func vals() (int, string){ 

	return 3, "Shreyash"
}

func main(){
	a,b := vals() //to call a multi return function we use the res1,res2,res3,... := functionName(param1,param2,param3,...)

	fmt.Println("a: ",a," b: ",b)
}