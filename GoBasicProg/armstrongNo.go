package main

import "fmt"

func main() {
	checkArmstrongNo(153)

}

func checkArmstrongNo(num int) {
	var tempNumber, remainder int

	result := 0

	tempNumber = num

	for {
		remainder = tempNumber % 10
		result += remainder * remainder * remainder
		tempNumber /= 10

		if tempNumber == 0 {
			break
		}
	}

	if result == num {
		fmt.Printf("%d is an Armstrong number.", num)
	} else {
		fmt.Printf("%d is not an Armstrong number.", num)
	}
}
