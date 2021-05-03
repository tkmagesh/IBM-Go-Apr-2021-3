package main

import "fmt"

func main() {

	//Arrays
	//var nos [5]int

	//var nos [5]int = [5]int{10, 20, 30, 40, 50}
	var nos [5]int = [...]int{10, 20, 30, 40, 50}
	fmt.Println(nos)

	/* for i := 0; i < len(nos); i++ {
		fmt.Println(nos[i])
	} */

	/*
		for idx, value := range nos {
			fmt.Printf("%d at %d\n", value, idx)
		}
	*/
	for _, value := range nos {
		fmt.Printf("value = %d\n", value)
	}

	//Slice
	var names []string = []string{"Magesh", "Suresh", "Ramesh"}
	fmt.Println(names)
	fmt.Printf("No of names = %d\n", len(names))

	//append
	//names = append(names, "Ganesh", "Rajesh")
	names = append(append([]string{}, "Ganesh", "Rajesh"), names...)
	fmt.Printf("After appending 2 more names, names  = %v\n", names)

	//slicing
	/*
		list[lo : hi] => from lo to hi-1
		list[:hi] => from 0 to hi-1
		list[lo :] => all the values from lo
		list[lo : lo] => empty
		list[lo : lo+1] => one element (at lo)
	*/
	fmt.Printf("name[1:4] = %v\n", names[1:4])
	fmt.Printf("name[3:] = %v\n", names[3:])
	fmt.Printf("name[:3] = %v\n", names[:3])
	fmt.Printf("After the above slice operations, names = %v\n", names)

	//create a new slice with all the values except 1 to 3

	//create a list of 50 random numbers (with range 0 to 200)
	//print the number of prime numbers generated in the 50 random numbers

	//use the math/rand package - rand.Intn(200)

	//r := rand.New(40)
	//r.Intn(100)

	//map
	cityRanks := map[string]int{
		"Mysuru":    2,
		"Mangaluru": 3,
		"Udupi":     1,
	}
	fmt.Printf("City Ranks = %v\n", cityRanks)
	fmt.Printf("Rank of Mangaluru = %d\n", cityRanks["Mangaluru"])

	fmt.Println("After adding a new city")
	cityRanks["Bengaluru"] = 5
	fmt.Printf("City Ranks = %v\n", cityRanks)

	fmt.Println("Iterating through a map")
	for key, value := range cityRanks {
		fmt.Printf("Key = %s, Value = %d\n", key, value)
	}

	fmt.Println("Check if Chennai exists")
	//cityRanks["Chennai"] = 7
	if rankOfChennai, exists := cityRanks["Chennai"]; exists {
		fmt.Println("Rank of Chennai = ", rankOfChennai)
	} else {
		fmt.Println("Chennai doesnot exist")
	}

}
