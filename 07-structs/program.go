package main

import "fmt"

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

func main() {
	/*
		product := Product{}
		product.id = 100
		product.name = "Pen"
		product.cost = 10
		product.units = 100
		product.category = "Stationary"
	*/
	product := Product{id: 100, name: "Pen", cost: 10, units: 100, category: "stationary"}
	fmt.Println(product)
	fmt.Println("After applying 10% discount")
	applyDiscount(&product, 10)
	fmt.Println(product)
}

func applyDiscount(product *Product, discount float32) {
	product.cost = product.cost * 0.9
}
