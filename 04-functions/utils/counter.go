package utils

//prevent this varibale to be accessble from other functions in the same package

var Counter = func() func() int {
	var count = 0
	return func() int {
		count += 1
		return count
	}
}()

/*
var Increment func() int
var Decrement func() int
*/

/*
var Spinner = func() (func() int, func() int) {
	count := 0
	var increment = func() int {
		count += 1
		return count
	}
	var decrement = func() int {
		count -= 1
		return count
	}
	return increment, decrement
}
*/

var Increment, Decrement func() int = func() (func() int, func() int) {
	count := 0
	var increment = func() int {
		count += 1
		return count
	}
	var decrement = func() int {
		count -= 1
		return count
	}
	return increment, decrement
}()
