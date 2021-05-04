package models

import (
	"fmt"
	"sort"
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

var currentComparer LessFn = productComparers["id"]

func (products *Products) Sort(attrName string) {
	currentComparer = productComparers[attrName]
	sort.Sort(products)
}

func (products Products) Less(i, j int) bool {
	return currentComparer((products)[i], (products)[j])
}

func (products Products) Swap(i, j int) {
	(products)[i], (products)[j] = (products)[j], (products)[i]
}

type LessFn func(p1, p2 Product) bool

var productComparers map[string]LessFn = map[string]LessFn{
	"Id": func(p1, p2 Product) bool {
		return p1.Id < p2.Id
	},
	"Name": func(p1, p2 Product) bool {
		return p1.Name < p2.Name
	},
	"Cost": func(p1, p2 Product) bool {
		return p1.Cost < p2.Cost
	},
	"Units": func(p1, p2 Product) bool {
		return p1.Units < p2.Units
	},
	"Category": func(p1, p2 Product) bool {
		return p1.Category < p2.Category
	},
}
