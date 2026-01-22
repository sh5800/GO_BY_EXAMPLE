package main

import "fmt"

type vehicle struct{
	wheels int
	engineModel string
}

// func(v *vehicle) setWheels(wheels int){
// 	v.wheels = wheels
// }

func(v *vehicle) getWheels() int{
	return v.wheels
}

type bmw struct{
	vehicle
	name string
}

func main(){
	car1 := bmw{
		vehicle: vehicle{
			wheels:4,
			engineModel:"2025",
		},
		name:"BMW-X4",
	}

	fmt.Println(car1.getWheels())
}