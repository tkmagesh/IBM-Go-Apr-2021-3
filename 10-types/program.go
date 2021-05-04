package main

import (
	"fmt"
	"strconv"
)

func main() {
	var val interface{}
	val = 10
	val = "abc"
	val = true
	fmt.Println(val)

	val = "abc"

	if no, ok := val.(int); ok {
		fmt.Println(no * 10)
	} else {
		fmt.Println("Value is not type integer")
	}

	val = true
	switch v := val.(type) {
	case int:
		fmt.Printf("Twice of %v is %v\n", v, v*2)
	case string:
		fmt.Printf("Len of %v is %v\n", v, len(v))
	default:
		fmt.Printf("Unknown type of %v = %T\n", v, v)

	}

	fmt.Println("Generic Sum fn")
	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30, 40))
	fmt.Println(sum(10))
	fmt.Println(sum())
	fmt.Println(sum(10, "20", 30, "40"))
	fmt.Println(sum(10, "20", 30, "40", "abc"))
	fmt.Println(sum(10, 20, []int{30, 40}))
	fmt.Println(sum(10, 20, []interface{}{30, 40, []int{10, 20}}))
	fmt.Println(sum(10, 20, []interface{}{30, 40, []interface{}{10, "20"}}))
}

func sum(nos ...interface{}) int {
	result := 0
	for _, no := range nos {
		switch no.(type) {
		case int:
			result += no.(int)
		case string:
			if val, err := strconv.Atoi(no.(string)); err == nil {
				result += val
			}
		case []int:
			noList := no.([]int)
			intfList := make([]interface{}, len(noList))
			for idx, x := range noList {
				intfList[idx] = x
			}
			result += sum(intfList...)
		case []interface{}:
			result += sum(no.([]interface{})...)
		}

	}
	return result
}
