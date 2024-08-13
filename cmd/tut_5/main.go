package main

import (
	"fmt"
	"strings"
)

func main() {

	var myString string = "résumé"

	for index, char := range myString {
		fmt.Println(index, char)
	}
	/*

		0 114	0 - r
		1 233	1,2 - `e
		3 115	3 - s
		4 117	4 - u
		5 109	5 - m
		6 233	6,7 - `e

	*/
	var myString2 = []rune("résumé")

	for index, char := range myString2 {
		fmt.Println(index, char)
	}

	/* Proper index matching
	0 114
	1 233
	2 115
	3 117
	4 109
	5 233
	*/

	// String building

	var strSlice = []string{"K", "a", "r", "t", "h", "i", "k"}
	var strBuilder strings.Builder

	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	var catStr = strBuilder.String()
	fmt.Println(catStr)

}
