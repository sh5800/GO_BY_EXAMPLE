package main 

import(
	"fmt"
	"slices"
)

func main(){
	var s1 []int  //Slice declaration syntax 1: var sliceName []dataType. //Slices can be considered as dynamic lists/arrays

	fmt.Println("s1: ",s1," length of s1: ",len(s1), " Capacity of s1: ",cap(s1))

	var s2 = make([]int,3,5) //Slice declaration syntax 2: var sliceName = make([]datatype,length,capacity)
	s2[0] = 1
	s2[1] = 2
	s2[2] = 3 //in the first go slice can be initialized using sliceName[index] = value 

	fmt.Println("s2: ",s2," length of s2: ",len(s2), " Capacity of s2: ",cap(s2))

	s2 = append(s2, 4) //Once slice has elements equivalent to the defined length (2nd arg in make function). Then to add more elements we should we use the append function. append can also be used to append elements to a slice after it is defined and not necessarily after it has reached it's defined length
	s2 = append(s2,5)

	fmt.Println("s2: ",s2," length of s2: ",len(s2), " Capacity of s2: ",cap(s2))

	s2 = append(s2,6) //If we append elements in a slice beyond it's capacity, then the slice will automatically double it's capacity to fit the other elements
	fmt.Println("s2: ",s2," length of s2: ",len(s2), " Capacity of s2: ",cap(s2))

	s3 := make([]int,6,10)
	copy(s3,s2)  //to copy elements of one slice to another slice till it's defined length (2nd arg in make function) we use the copy operator

	fmt.Println("s3: ",s3," length of s3: ",len(s3), " Capacity of s3: ",cap(s3))

	if slices.Equal(s3,s2){ //to compare if two slices are equal we use the 'Equal' function present in 'slices' package
		fmt.Println("s2 == s3")
	}

	l := s3[2:5] //to extract specific part of slice we use the : operator like sliceName[startIndex:endIndex]. here the element of the startIndex will be the part of the newSlice, whereas the element just before the endIndex will be the part of the new slice not the element at the end index. in short element of startIndex is inclusive and element at endIndex is exclusive.
	fmt.Println("l: ",l)

	m := s3[:3] //extract everything from the start till endIndex-1
	fmt.Println("m: ",m)

	n := s3[0:] //extract everything from the start
	fmt.Println("n: ",n)

	p := s3[2:] //extract everything from the the startIndex which is 2
	fmt.Println("p: ",p)

	twoD := make([][]int, 3)
    for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}