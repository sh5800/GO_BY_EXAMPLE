package main 

import(
	"fmt"
	"time"
)

func main(){
	i := 5
	fmt.Println("Write ",i," as: ")

	switch i{   //switch looks the same as other OOP languages, just here we do not have a break statement after every case
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")		
	}

	switch time.Now().Weekday(){ //multiple cases can be clubbed together with ,
	case time.Saturday,time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")	
	}

	t := time.Now()
	switch{ //switch without condition is equivalent to if else-if else blocks case -> if/else-if: default -> else
	case t.Hour() < 12:
		fmt.Println("Before Noon")
	default:
		fmt.Println("After noon")	
	}

	whatAmi := func(i interface{}){
		switch i.(type){
		case int:
			fmt.Println("Integer")
		case float64:
			fmt.Println("Float")
		default:
			fmt.Println("Boolean")		
		}
	}

	whatAmi(12)
	whatAmi(45.00)
	whatAmi(true)
}