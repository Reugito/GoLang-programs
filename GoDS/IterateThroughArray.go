package main

import "fmt"

func main() {
	intArray := [5]int{10, 20, 30, 40, 50}

	// normal loop depending on lenth of array
	for i := 0; i < len(intArray); i++ {
		fmt.Printf("intArray[%v]: %v\n", i, intArray[i])
	}

	// loop using range  index and element both present
	for index, element := range intArray {
		fmt.Printf("intArra[%v]: %v\n", index, element)

	}

	// loop using range only value is present
	for _, value := range intArray {
		fmt.Println(value)
	}

	j := 0
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}
