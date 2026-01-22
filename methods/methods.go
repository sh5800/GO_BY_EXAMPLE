package main

import "fmt"

type person struct{
	name string
	age int
}

func(p1 person) getName() string{ //when a function has a receiver type of a struct it is called a method. A function that is associated with a struct is called as a method
	return p1.name
}

func(p1 person) getAge() int{
	return p1.age
}
func main(){
	p1 := person{
		name:"Shreyash",
		age:25,
	}

	fmt.Println(p1.getName())
	fmt.Println(p1.getAge())
}