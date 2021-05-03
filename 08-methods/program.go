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
}
