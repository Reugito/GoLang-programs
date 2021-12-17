package main

import "fmt"

func main() {
	num := 0
	fmt.Print("Enter the No: ")
	fmt.Scan(&num)
	fmt.Printf("%v is palindrom No: %v", num, isPalindromeNo(num))

}

func isPalindromeNo(num int) bool {
	reverse, remainder, temp := 0, 0, 0
	temp = num
	for {
		remainder = temp % 10
		reverse = reverse*10 + remainder
		temp /= 10
		if temp == 0 {
			break
		}
	}
	if num == reverse {
		return true
	} else {
		return false
	}
}
