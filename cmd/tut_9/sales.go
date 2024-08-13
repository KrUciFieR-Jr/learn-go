package main

import (
	"fmt"
	"math/rand"
	"time"
)

type saleDetails struct {
	saleItem string
	website  string
	price    float32
}

var MAX_CHICKEN_PRICE float32 = 10
var MAX_TOFU_PRICE float32 = 50

func main() {
	var channel = make(chan saleDetails)

	var websites = []string{"walmart.com", "apnimandi.com", "safeway.com"}

	for _, website := range websites {
		go checkChickenPrices(website, channel)
		go checkTofuPrices(website, channel)
	}

	sendMessage(channel)

}

func checkChickenPrices(website string, chickenChannel chan saleDetails) {
	for {
		fmt.Println("Checking Chicken price ! on:", website)
		time.Sleep(1 * time.Second)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice < MAX_CHICKEN_PRICE {
			chickenChannel <- saleDetails{saleItem: "chicken", website: website, price: chickenPrice}
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan saleDetails) {
	for {
		fmt.Println("Checking Tofu price ! on:", website)
		time.Sleep(1 * time.Second)
		var tofuPrice = rand.Float32() * 50
		if tofuPrice < MAX_TOFU_PRICE {
			tofuChannel <- saleDetails{saleItem: "tofu", website: website, price: tofuPrice}
			break
		}
	}
}

func sendMessage(channel chan saleDetails) {
	select {
	case sale := <-channel:
		switch sale.saleItem {
		case "tofu":
			sendEmail(sale)
		case "chicken":
			sendText(sale)
		default:
			fmt.Println("The code broke !")

		}
	}
}

func sendEmail(sale saleDetails) {
	fmt.Printf("Sending email: %s at %s with price %.2f\n", sale.saleItem, sale.website, sale.price)
}

func sendText(sale saleDetails) {
	fmt.Printf("Sending text: %s at %s with price %.2f\n", sale.saleItem, sale.website, sale.price)
}

/*
Checking Tofu price ! on: apnimandi.com
Checking Chicken price ! on: apnimandi.com
Checking Chicken price ! on: walmart.com
Checking Tofu price ! on: safeway.com
Checking Tofu price ! on: walmart.com
Checking Chicken price ! on: safeway.com
Sending text: chicken at safeway.com with price 7.99

*/
