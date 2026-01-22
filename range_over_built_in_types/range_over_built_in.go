package main

import "fmt"

func main() {
	arr1 := [5]int{1,2,3,4,5}

	for index,value := range arr1{
		fmt.Println("Index: ",index, " Value: ",value)
	}

	slice1 := []int{4,5,6,7,8}

	for index,value := range slice1{
		fmt.Println("Index: ",index, " Value: ",value)
	}

	map1 := map[string]int{"foo":1,"bar":2}

	for key,val := range map1{
		fmt.Println("Key: ",key, " Value: ",val)
	}

	for key := range map1{
		fmt.Println("Key: ",key)
	}

	for _,val := range map1{
		fmt.Println("Value: ",val)
	}
}