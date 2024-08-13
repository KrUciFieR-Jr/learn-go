package main

import "fmt"

func main() {

	var p *int32 = new(int32)
	var i int32

	fmt.Println("The value p points to is: %v", *p)
	fmt.Println("The value i is: %v", i)

	// P point to i

	p = &i
	*p = 1
	fmt.Println("The value p points to is: %v", *p)
	fmt.Println("The value i is: %v", i)

	var thing = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The location of thing array is : %p", &thing)

	var result [5]float64 = square(&thing)

	fmt.Printf("\nThe result is : %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing)
	fmt.Printf("\nThe location of result array is : %p", &result)

}

func square(thing *[5]float64) [5]float64 {

	for i := range thing {
		thing[i] = thing[i] * thing[i]
	}
	return *thing
}