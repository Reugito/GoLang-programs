package main

import "fmt"

func main() {
	var num int
	fmt.Print("enter the number: ")
	fmt.Scan(&num)
	fmt.Printf("%v is Duck No. %v", num, isDuckNo(num))
}

func isDuckNo(no int) bool {
	for {
		if no%10 == 0 {
			return true
		}
		no /= 10
		println(no)
		if no <= 1 {
			break
		}
	}
	return false
}
