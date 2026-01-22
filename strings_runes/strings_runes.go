package main

import(
	"fmt"
)

func main(){
	const s = "Shreyash" // a string in golang is represented as varName := "string". Strings in golang are byte slices []byte

	for i := range s{
		fmt.Println(s[i])   //this will print the ascii value of each character
	}

	for i := range s{
		fmt.Printf("%c\n",s[i]) //this will print the actual character
	}
}