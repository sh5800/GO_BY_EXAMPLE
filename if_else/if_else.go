package main

import "fmt"

func main(){
	a,b := 20,10

	if a > b{  //typical if-else code block
		fmt.Println("A is greater")
	}else{
		fmt.Println("B is greater")
	}

	if a == 20{ //if can be used alone
		fmt.Println("Value of A is: ",a)
	}

	if x := 33; x < 10{ //if can be clubbed with multiple else if's 
		fmt.Println("X is less than 10")
	}else if(x % 11 == 0){
		fmt.Println("X is divisible by 11")
	}else{
		fmt.Println("X is greater than 10")
	}

	//go doesn't have a ternary operator and to implement it we have to rely on if-else statements
}