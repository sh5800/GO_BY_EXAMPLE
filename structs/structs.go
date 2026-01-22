package main

import "fmt"

type person struct{  //structs are used to wrap different kinds of data together
	name string
	age int
}

func newPerson(name string, age int) *person{ 
	return &person{name:name,age:age}
}

func main(){
	p1 := person{name:"Shreyash",age:25} //initialization of a struct

	fmt.Println(p1.name)
	fmt.Println(p1.age)

	p2 := newPerson("Aryan",25)
	fmt.Println(p2.name)
	fmt.Println(p2.age)

	dog := struct{  //struct can be declared and initialized explicity
		name string
		isGood bool
	}{
		name:"Rex",
		isGood: true,
	}

	fmt.Println(dog)
}