package main

import "fmt"

func main() {
	fmt.Println("Enter 3 numbers :")
	var first int
	fmt.Scanln(&first)
	var second int
	fmt.Scanln(&second)
	var third int
	fmt.Scanln(&third)

	if first >= second && first >= third {
		fmt.Println(first, "is Largest among three numbers.")
	}
	if second >= first && second >= third {
		fmt.Println(second, "is Largest among three numbers.")
	}
	if third >= first && third >= second {
		fmt.Println(third, "is Largest among three numbers")
	}
}
