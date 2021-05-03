package main

import (
	"fmt"
	"methods-demo/models"
)

func main() {
	grapes := models.NewProduct(200, "Grapes", 120, 50, "Food", 2)
	//fmt.Println(grapes)
	//fmt.Println(models.ToString(grapes))
	fmt.Println(grapes.ToString())
	//models.ApplyDiscount(grapes, 10)
	grapes.ApplyDiscount(10)
	//fmt.Println(models.ToString(grapes))
	fmt.Println(grapes.ToString())
}
