package main

import(
	"fmt"
)

func main(){
	var a = 3    //var keyword is used to declare a variable. //Syntax var var1,var2,var3 ... = val1,val2,val3
	fmt.Println("Value of a is: ",a)
	fmt.Printf("Type of a is: %T\n",a)

	var x,y,z int = 1,3,5 //Another syntax var var1,var2,var3 ... datatype = value1,value2,value3.
	fmt.Println("a,b,c: ", x,y,z)

	var e int  //var variablename1,variablename2,variablename3 .... datatype
	fmt.Println(e)

	n := 3 + 4i  //To declare and use a variable directly we can use the := operator. The variable on the left of := operator is then auto assigned the data type based on what value is present on the right side of := operator
	fmt.Printf("Type of n is: %T\n",n)

	u := 'c' //int32 or rune data type
	l := "apple"

	fmt.Printf("Type of u is: %T\n",u)
	fmt.Println("l is: ",l)
}