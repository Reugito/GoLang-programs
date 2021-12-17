package main

import "fmt"

func main() {
	var array [10]int

	for i := 0; i < 10; i++ {
		arrayEle := 0
		fmt.Printf("Enter ele at index: %v\n", i)
		fmt.Scan(&arrayEle)
		array[i] = arrayEle
	}
	fmt.Println(array)
}
