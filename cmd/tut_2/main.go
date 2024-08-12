package main

import "fmt"
import "unicode/utf8"

func main(){
	
	var intNum int16 = 12
	var intNum1 int16 = 17
	fmt.Println(intNum/intNum1)
	fmt.Println(intNum%intNum1)

	var myString string = "Hello \n World"
	var myString2 string = "Hello" + " "+ "Karthik"
	var myString3 string = `Hellowwww  
		World `

	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(myString3)
	
	var tryThisLength string = "漢"
	var oneCharacter string = "A"
	
	fmt.Println(len(oneCharacter))
	fmt.Println(len(tryThisLength))
	
	
	fmt.Println(utf8.RuneCountInString(oneCharacter))
	fmt.Println(utf8.RuneCountInString(tryThisLength))
	
	var MyRune rune = 'a'
	fmt.Println(MyRune)

	var myBollean bool = false
	fmt.Println(myBollean)

	var var1, var2 int = 1,2
	ivar1, ivar2 := 1,2
 
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(ivar1)
	fmt.Println(ivar2)

	// You cannot change this since it is a constant
	const myConst string = "const value"


	fmt.Println(myConst)
}
