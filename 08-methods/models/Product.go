package models

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	Expiry int
}

func NewProduct(Id int, Name string, Cost float32, Units int, Category string, expiry int) *PerishableProduct {
	return &PerishableProduct{Product{Id, Name, Cost, Units, Category}, expiry}
}

func (product *PerishableProduct) ApplyDiscount(discount float32) {
	product.Cost = product.Cost * 0.9
}

func (product *PerishableProduct) ToString() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s, Expiry=%d", product.Id, product.Name, product.Cost, product.Units, product.Category, product.Expiry)
}

func (product Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%f, Units=%d, Category=%s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}
