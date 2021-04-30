package utils

//prevent this varibale to be accessble from other functions in the same package
var count = 0

func Counter() int {
	count += 1
	return count
}
