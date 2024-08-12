package main

import (
	"errors"
	"fmt"
)

func main(){

	var printValue string = "Hellow World"
	printMe(printValue)

	var num int = 11
	var den int = 5
	
	var result, remain, err = intDiv(num,den)

	switch {
		case err != nil:
			fmt.Printf(err.Error())
		case remain == 0:
			fmt.Println("The result of the division is %v", result)
		default:
			fmt.Println("The result and remainder of the division is %v, %v respectively", result, remain)
	}
	
	switch remain {
		case 0:
			fmt.Println("The remainder is 0 from conditional switch")
		default:
			fmt.Println("The remainder is not 0 from conditional switch")
	
	}	
}

func printMe(printValue string){
	
	fmt.Println(printValue)
}


func intDiv(num int, den int)(int, int, error){
	
	var err error

	if den == 0 {
		err = errors.New("Cannot divide by 0")
		return 0,0, err
	}
	var result int = num/den
	var remainder int = num%den

	return result, remainder, err
}
