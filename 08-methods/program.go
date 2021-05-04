package main

import (
	"fmt"
	"methods-demo/models"
)

type myStr string

func (s myStr) length() int {
	return len(s)
}

func main() {
	/*
		grapes := models.NewProduct(200, "Grapes", 120, 50, "Food", 2)
		//fmt.Println(grapes)
		//fmt.Println(models.ToString(grapes))
		fmt.Println(grapes.ToString())
		//models.ApplyDiscount(grapes, 10)
		grapes.ApplyDiscount(10)
		//fmt.Println(models.ToString(grapes))
		fmt.Println(grapes.ToString())

		//var test myStr = "Some dummy string"
		test := myStr("Some dummy string")
		fmt.Println(test.length())

		var s string = "Some other dummy string"
		fmt.Println(myStr(s).length())
	*/

	/* pp1 := models.NewProduct(200, "Grapes", 120, 50, "Food", 2)
	pp2 := models.NewProduct(200, "Grapes", 120, 50, "Food", 2)
	fmt.Println(*pp1 == *pp2) */

	products := &models.Products{}
	products.AddProduct(models.Product{105, "Pen", 10, 100, "Stationary"})
	products.AddProduct(models.Product{103, "Pencil", 5, 200, "Stationary"})
	products.AddProduct(models.Product{102, "Mouse", 400, 10, "Stationary"})
	products.AddProduct(models.Product{104, "Pan", 500, 5, "Utencil"})
	products.AddProduct(models.Product{101, "Water Dispenser", 250, 10, "Utencil"})

	fmt.Println(products)
	fmt.Println(products.IndexOf(models.Product{101, "Pen", 10, 100, "Stationary"}))
	fmt.Println(products.Includes(models.Product{101, "Pen", 10, 100, "Stationary"}))

	anyStationaryProducts := products.Any(func(product models.Product) bool {
		return product.Category == "Stationary"
	})
	fmt.Println("Any stationary products ? :", anyStationaryProducts)

	allStationaryProducts := products.All(func(product models.Product) bool {
		return product.Category == "Stationary"
	})
	fmt.Println("All stationary products ? :", allStationaryProducts)

	stationaryProducts := products.Filter(func(product models.Product) bool {
		return product.Category == "Stationary"
	})

	fmt.Println("Stationary Products :")
	fmt.Println(stationaryProducts)

	fmt.Println("Sorting")
	/* names := []string{"Magesh", "Ramesh", "Ganesh", "Rajesh", "Suresh"}
	sort.Strings(names)
	fmt.Println(names) */

	//sort by Id
	fmt.Println("Sort Products by Id")
	products.Sort("Id")
	fmt.Println(products)

	//sort by Name
	fmt.Println("Sort Products by Name")
	products.Sort("Name")
	fmt.Println(products)

	//sort by Cost
	fmt.Println("Sort Products by Cost")
	products.Sort("Cost")
	fmt.Println(products)

	//sort by Units
	fmt.Println("Sort Products by Units")
	products.Sort("Units")
	fmt.Println(products)

	//sort by Category
	fmt.Println("Sort Products by Category")
	products.Sort("Category")
	fmt.Println(products)
}
