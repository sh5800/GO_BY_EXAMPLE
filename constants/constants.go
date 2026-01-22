package main

import(
	"fmt"
)

const name string = "Shreyash"  //const keyword declares a constant. The value of a constant cannot be changed. Syntax const varName datatype = value

const(
	dob = "05 Aug 2000"
	residence = "Mumbai"
)

func main(){
	const age = 25  //syntax const varName = value

	
	
	fmt.Println("Name is: ",name)

	fmt.Println("Age is: ",age)

	fmt.Println("Date of Birth is: ",dob)

	fmt.Println("Residence is: ",residence)


}