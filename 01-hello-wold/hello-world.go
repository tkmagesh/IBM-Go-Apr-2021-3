//package declaration
package main

//package imports
import "fmt"

//package level variables

//package init

//main func
func main() {
	//fmt.Println("Hello there! Welcome to Golang!")

	/*
		var name string
		name = "Magesh"
	*/

	/*
		var name string = "Magesh"
	*/

	/*
		var name = "Magesh"
	*/

	name := "Magesh" /* applicable only in a function scope variable */
	//fmt.Println("Hello " + name + "! Welcome to Golang!")
	fmt.Printf("Hello %s!, Welcome to Golang!\n", name)

	/*
		var username string
		var message string
	*/

	/*
		var (
			username string
			message  string
		)
		username = "Philip"
		message = "Welcome to Golang!"
	*/

	/*
		var (
			username, message string
		)
	*/
	/*
		var username, message string
		username = "Philip"
		message = "Welcome to Golang!"
	*/
	/*
		var (
			username string = "Philip"
			message  string = "Welcome to Golang!"
		)
	*/

	/*
		var username, message string = "Philip", "Welcome to Golang!"
	*/

	/*
		var username, message = "Philip", "Have a nice day!"
	*/
	username, message := "Philip", "Have a nice day!"
	greetMsg := fmt.Sprintf("Hello %s!, %s\n", username, message)
	fmt.Println(greetMsg)

}

//other functions
