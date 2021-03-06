package models

import (
	"fmt"
)

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

func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	(products)[i], (products)[j] = (products)[j], (products)[i]
}

type ByName struct{ Products }

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

type ByUnits struct{ Products }

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

type ByCost struct{ Products }

func (byCost ByCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

type ByCategory struct{ Products }

func (byCategory ByCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

/*
Pinky
	- a new type has to be created for each comparison
	+ adheres to Open Closed Principle

Magesh
	+ a new comparer need to added to the productComparer map
	- violation of Open Closed Princile
*/
