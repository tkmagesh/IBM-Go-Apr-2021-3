package main

import "fmt"

type Product struct {
	id       int
	name     string
	cost     float32
	units    int
	category string
}

/*
type PerishableProduct struct {
	product Product
	expiry  int
}
*/

type PerishableProduct struct {
	Product
	expiry int
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

	products := []Product{}
	addProduct(&products, 102, "Pencil", 5, 200, "Stationary")
	addProduct(&products, 103, "Mouse", 400, 10, "Stationary")
	addProduct(&products, 104, "Pan", 500, 5, "Utencil")
	addProduct(&products, 105, "Water Dispenser", 250, 10, "Utencil")
	fmt.Println(products)

	fmt.Println("Composition")
	/*
		grapes := PerishableProduct{
			product: Product{200, "Grapes", 120, 20, "Food"},
			expiry:  3,
		}
	*/
	/* grapes := PerishableProduct{Product{200, "Grapes", 120, 20, "Food"}, 2} */
	//grapes := PerishableProduct{Product: Product{200, "Grapes", 120, 20, "Food"}, expiry: 2}
	grapes := newPerishableProduct(200, "Grapes", 120, 20, "Food", 3)
	fmt.Println(grapes)
	//fmt.Println("Cost of grapes = ", grapes.Product.cost)
	fmt.Println("Cost of grapes = ", grapes.cost)
}

//an utilify function to hide the complexity of constructing a complex struct
func newPerishableProduct(id int, name string, cost float32, units int, category string, expiry int) *PerishableProduct {
	return &PerishableProduct{Product{id, name, cost, units, category}, expiry}
}

func addProduct(products *[]Product, id int, name string, cost float32, units int, category string) {
	*products = append(*products, Product{id, name, cost, units, category})
}

func applyDiscount(product *Product, discount float32) {
	product.cost = product.cost * 0.9
}

// create a utility function to add a product to the products slice
