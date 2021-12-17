package main

import "fmt"

func main() {
	months := map[interface{}]interface{}{}

	var ker, val interface{}
	fmt.Scan(&ker)
	fmt.Scan(&val)
	months[ker] = val
	fmt.Println(months)
}
