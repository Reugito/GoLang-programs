package main

import "fmt"

func main() {
	months := map[int]string{1: "JAN", 2: "FEB", 3: "MARCH", 4: "APREL", 5: "MAY",
		6: "JUNE", 7: "JULY", 8: "AUG", 9: "SEPT", 10: "OCT", 11: "NOV", 12: "DEC"}

	for i := 1; i <= len(months); i++ {
		fmt.Printf("%v: %v ", i, months[i])
	}
	println()
	j := 1
	for range months {
		fmt.Printf("%v: %v ", j, months[j])
		j++
	}

	println()
	for key, val := range months {
		fmt.Printf("%v: %v ", key, val)
	}

}
