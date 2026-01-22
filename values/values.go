package main 

import(
	"fmt"
	"reflect"
)

func main(){
	fmt.Println("Shreyash" + "Kashyap")  //Strings in go can be appended with + operator
	fmt.Println("3+6= ",3+6)             // integers can be represented in go in three different formats int, int32(covers unicode as well) and int64(for long long integer values)
	fmt.Println("7.00/3.00 = ",7.00/3.00) //floats can be represented in go in three ways float,float32 and float64
	fmt.Println(7 + 3i)  //complex numbers of the form a + ib where i = sqrt(-1) also called as iota

	fmt.Println(true || false)  //booleans true or false
	fmt.Println(false && true)
	fmt.Println(!false)

	a := 3
	fmt.Printf("The type of a is :%T\n",a) //%T is used to print the type of a variable
	fmt.Println("Determining the type of variable using reflect",reflect.TypeOf(a))       //reflect.TypeOf is used to print the type of a variable
	fmt.Println("Determining the type of variable using reflect",reflect.TypeOf(a).Kind()) //reflect.TypeOf(var).Kind() is used to determine the fundamental data type of the variable
}