package main

import (
	"fmt"
)

var println = fmt.Println
var printf = fmt.Printf
var scan = fmt.Scan

func main() {
	var intSlice []int
	for {
		var choice int

		println("\n1. Show Array")
		println("2. Show size of array")
		println("3. Append element")
		println("4. Append multiple elements")
		println("5. update Element by index")
		println("6. Remove Element by index")
		println("7. Remove element")
		println("8. check if element exist")
		println("9. count the given element")
		println("10. check capacity of slice")
		println("100. Quit")

		printf("\nEnter Your Choice: ")
		scan(&choice)

		switch choice {
		case 1:
			println("Int array : ", intSlice)
		case 2:
			println("Size of array :", len(intSlice))
		case 3:
			intSlice = appendElement(intSlice, intInput("element"))
			println("intArray after Appending Element: ", intSlice)
		case 4:
			intSlice = appendElements(intSlice, intSliceInput())
			println("intArray after Appending multiple Elements: ", intSlice)
		case 5:
			intSlice = updateElementByIndex(intSlice, intInput("Index")-1, intInput("element"))
			println("intArray after Updating the element By index: ", intSlice)
		case 6:
			intSlice = removeElementByIndex(intSlice, intInput("index to remove")-1)
			println("intArray after removing the element By index: ", intSlice)
		case 7:
			intSlice = removeElement(intSlice, intInput("element to remove"))
			println("intArray after removing the element : ", intSlice)
		case 8:
			ele := intInput("elemet ")
			printf("%v exists in %v : %v\n", ele, intSlice, checkIfItemExist(intSlice, ele))
		case 9:
			ele := intInput("element to count")
			printf("The count of %v in %v = %v\n", ele, intSlice, countTheElement(intSlice, ele))
		case 10:
			println("Capacity of array :", cap(intSlice))
		case 100:
			break
		default:
			println("Invalid choice")
		}

		if choice == 100 {
			break
		}

	}
}

// @note This method is to show how to declare different types of slice
func declareSlice() {

	intSlice1 := []int{1, 4, 6, 12, 11} // capacity
	intSlice2 := intSlice1[2:4]
	intSlice2 = append(intSlice2, 27, 100, 101)
	printf("int array: %v, size of slice is %v and capacity of slice is: %v\n", intSlice2, len(intSlice2), cap(intSlice2))

	strArr := []string{"Jan", "feb", "march"}
	printf("String slice: %v\n", strArr)

	complexArr := []complex128{1 + 2i}
	printf("Complex slice: %v\n", complexArr)

	float64Array := []float64{1.2, 3.45}
	printf("float64 slice: %v\n", float64Array)
}

/*
 @note This method is to take int input from console
 @return inputVal
*/
func intInput(printStatement string) int {
	var inputVal int
	printf("enter %v : ", printStatement)
	scan(&inputVal)
	return inputVal
}

/*
 @note This method is to take array of number as input from console
 @return intArray
*/
func intSliceInput() []int {
	var size int
	var arr []int
	printf("enter size of array: ")
	scan(&size)
	for i := 0; i < size; i++ {
		arr = append(arr, intInput("element"))
	}
	return arr
}

/*
 @note This method is to append the given element in given array
 @param intSlice, element
 @return array with appended value
*/
func appendElement(intSlice []int, element int) []int {
	return append(intSlice, element)
}

/*
 @note This method is to append the given elements in given array
 @param intSlice, elements
 @return array with appended values
*/
func appendElements(intSlice, elements []int) []int {
	return append(intSlice, elements...)
}

/*
 @note This method is to update a slice element at specific index
 @param intArr, index , element
 @return updated intSlice
*/
func updateElementByIndex(intSlice []int, index, element int) []int {
	intSlice[index] = element
	return intSlice
}

/*
 @note This method is to remove a slice element from given index
 @param intArr, index , element
 @return intSlice
*/
func removeElementByIndex(intSlice []int, index int) []int {
	return append(intSlice[:index], intSlice[index+1:]...)
}

/*
 @note This method is to remove the given element from Slice
 @param intArr, index , element
 @return intSlice
*/
func removeElement(intSlice []int, element int) []int {

	for index, val := range intSlice {
		if val == element {
			return append(intSlice[:index], intSlice[index+1:]...)
		}
	}
	return intSlice
}

/*
 @note This method is to check if given slice element is present or not
 @param intArr, element
 @return boolean value epending on element is present or not
*/
func checkIfItemExist(intSlice []int, element int) bool {
	for _, val := range intSlice {
		if val == element {
			return true
		}
	}
	return false
}

/*
 @note This method is to get the number of times given element present in Slice
 @param intArr, element
 @return count
*/
func countTheElement(intSlice []int, element int) int {
	count := 0
	for _, val := range intSlice {
		if val == element {
			count = count + 1
		}
	}
	return count
}
