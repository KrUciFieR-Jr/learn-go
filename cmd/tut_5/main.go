package main

import "fmt"

type gasEngine struct {
	mpg     uint16
	gallons uint16
	owner
}

type electricEngine struct {
	mpkwh uint16
	kwh   uint16
	owner
}

func (e gasEngine) milesLeft() uint16 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint16 {
	return e.mpkwh * e.kwh
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() uint16
}

func canMakeIt(eng engine, expectedMiles uint16) {
	fmt.Println(eng.milesLeft())
	if expectedMiles <= eng.milesLeft() {
		fmt.Println("You can make it there")
	} else {
		fmt.Println("You can't make it there")
	}
}
func main() {

	var myGasEngine = gasEngine{25, 15, owner{"Karthik"}}
	var myElectricEngine = electricEngine{15, 15, owner{"Karthik"}}
	fmt.Println(myGasEngine.mpg, myGasEngine.gallons, myGasEngine.name)
	fmt.Println(myElectricEngine.mpkwh, myElectricEngine.kwh, myElectricEngine.name)

	canMakeIt(myElectricEngine, 250)
	canMakeIt(myGasEngine, 250)

}
