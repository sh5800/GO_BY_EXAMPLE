package main

import (
	"fmt"
	"strconv"
	"maps"
)

func main(){
	var m1 map[string]int   //Syntax to declare an empty map
	fmt.Println("m1: ",m1) //m1:  map[]

	m1 = make(map[string]int) //To actually declare a map without nil length and capacity we use the make function as make(map[keyDataType]valDataType)
	m1["k1"] = 1 //initialization of map
	m1["k2"] = 2

	fmt.Println("m1: ",m1) //m1:  map[k1:1 k2:2]

	v1 := m1["k2"] //to get the value for a given key 
	fmt.Println("v1:",v1)

	m1["l1"] = 4
	m1["l2"] = 5

	fmt.Println("m1:",m1)

	v2 := m1["l2"]
	fmt.Println("v2:",v2)

	delete(m1,"l2")  //to delete a key we use the delete function as delete(mapName,key), key must be of the same type as defined when creating the map.
	fmt.Println("m1:",m1)

	//clear(m1) //to completely clear the contents of a map we use the clear function as clear(mapName)

	val,exists := m1["k2"]  //to check if a key exists in a map we use the syntax, valueAtKey,exists := mapName[keyName]
	if exists{
		fmt.Println("They key: k2 with value: ",strconv.Itoa(val)," exists in map m1")
	}

	map1 := map[string]int{"foo":10,"bar":20}
	map2 := map[string]int{"foo":100,"bar":200}

	fmt.Println(maps.Equal(map2,map1))  //to check if two maps are equal we use the Equal function of the map package as maps.Equal(map1,map2)
}