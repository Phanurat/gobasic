package main

import "fmt"

var productName [4]string
var price [4]float32

func main() {
	productName[0] = "Macbook"
	productName[1] = "Ipad"
	productName[2] = "Iphone"
	productName[3] = "Airpods"

	price := [4]float32{40000, 30000, 20000, 4000}

	fmt.Println(productName)
	fmt.Println(price)
}
