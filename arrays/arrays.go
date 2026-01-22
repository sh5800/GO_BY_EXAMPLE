package main

import "fmt"

func main(){
	var a [5]int   //Array Declaration syntax 1: var arrayName [size]dataType
	fmt.Println("Array A: ",a)

	a[3] = 250 //Assigning value to an index of an array
	fmt.Println("Array A: ",a)

	b := [5]int{1,2,3,4,5}  //Array Declaration syntax 2: arrayName := [size]dataType{val1,val2,val3,....}
	fmt.Println("Array B: ",b)

	b = [...]int{1,2,3,4,5} //Array Declaration syntax 3: arrayName := [...]dataType{val1,val2,val3,...}. Here ... represents spread operator (maybe)
	fmt.Println("Array B: ",b)

	var twoD [2][3] int //2D Array Declaration syntax 1: var arrayName [rows][cols] dataType
	
	for i := range len(twoD){  
		for j := range len(twoD[i]){
			twoD[i][j] = i + j //2D Array initialization
		}
	}

	fmt.Println("2D array: ",twoD)

	twoD = [2][3]int{
		{1,2,3},
		{4,5,6},
	}

	fmt.Println("2D array: ",twoD)

	var threeD [1][2][3] int

	for i := range len(threeD){
		for j := range len(threeD[i]){
			for k := range len(threeD[i][j]){
				threeD[i][j][k] = i + j + k
			}
		}
	}

	fmt.Println("3D Array: " ,threeD)
}