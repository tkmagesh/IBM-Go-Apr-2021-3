package models

import "fmt"

type Products []Product

func (products *Products) AddProduct(product Product) {
	*products = append(*products, product)
}

func (products *Products) IndexOf(product Product) int {
	for idx, p := range *products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products *Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products *Products) Any(predicate func(Product) bool) bool {
	for _, product := range *products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func (products *Products) All(predicate func(Product) bool) bool {
	for _, product := range *products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func (products *Products) Filter(predicate func(Product) bool) *Products {
	result := &Products{}
	for _, product := range *products {
		if predicate(product) {
			*result = append(*result, product)
		}
	}
	return result
}

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%v\n", product)
	}
	return result
}

func (products *Products) Sort() {
	/*  */
}
