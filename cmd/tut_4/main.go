package main

import "fmt"


func main(){
        var intArr [3]int32

        // intArr := [3]int32{1,2,3}
        // intArr := [...]int32{1,2,3} // This inferes the size of array
        intArr[1] = 123
        fmt.Println(intArr)

        var intSlice []int32 = []int32{4,5,6}
        fmt.Printf("the length before append %v and capacity %v\n Arr: %v", len(intSlice), cap(intSlice),intSlice)
        fmt.Println()

        intSlice = append(intSlice,7)
        fmt.Printf("the length before append %v and capacity %v\n Arr: %v", len(intSlice), cap(intSlice),intSlice)
        fmt.Println()

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice,intSlice2...)
        fmt.Printf("the length before append %v and capacity %v\n Arr: %v", len(intSlice), cap(intSlice),intSlice)
        fmt.Println()

	var intSliceNew []int32 = make([]int32,3,8)
        fmt.Printf("the length before append %v and capacity %v\n Arr: %v", len(intSliceNew), cap(intSliceNew),intSliceNew)
        fmt.Println()


	// MAPS
	var myMap map[string]uint8 = make(map[string]uint8)

	fmt.Println(myMap)

	var myMap2 map[string]uint8 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2)
	
	var age, ok = myMap2["Jason"]	

	if ok{
		fmt.Printf("The age is %v\n", age)

	} else{
		fmt.Println("Didn't find it in the map")
	}

	
	// LOOPS

	for name, age := range myMap2{
		fmt.Printf("The Name and age is %v and %v\n", name,age)
	
	}

	for index, num := range intSlice{
		fmt.Printf("The index and num is %v -> %v\n", index,num)

	}

	for i:=0; i<10; i++{
		fmt.Printf("The num is %v\n", i)

		
	}
}    
