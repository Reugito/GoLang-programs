package main

import "fmt"

var printf = fmt.Printf
var println = fmt.Println
var print = fmt.Print

const size = 9

func main() {
	var intArray = [size]int{1, 2, 8, 9, 11, 3}
	var choice int
	for {
		println("\n1. show array")
		println("2. show lenth array")
		println("3. add element in array")
		println("4. remove array element by index")
		println("5. remove array element by value")
		println("6. check index of array element")
		println("7. reverse array")
		println("8. sort array in ascending order")
		println("9. sort array in descending order")
		println("10. count a element in array")
		println("11. get sum of array")
		println("12. get Avaragr of array")
		println("13. get max element of array")
		println("14. get min element of array")
		println("100. quit")
		printf("\nEnter the choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			println("array = ", intArray)
		case 2:
			println("lenth of array is: ", len(intArray))
		case 3:
			var ele, index int
			print("Enter index : ")
			fmt.Scan(&index)
			print("Enter element to add: ")
			fmt.Scan(&ele)
			intArray = appendIntElement(intArray, index-1, ele)
			println("New array ", intArray)
		case 4:
			var index int
			printf("Enter index to remove: ")
			fmt.Scan(&index)
			intArray = removeIntElementByIndex(intArray, index-1)
			println("New array ", intArray)
		case 5:
			var ele int
			printf("Enter element to remove: ")
			fmt.Scan(&ele)
			intArray = removeIntElementByValue(intArray, ele)
			println("New array ", intArray)
		case 6:
			var ele int
			printf("Enter element to get index: ")
			fmt.Scan(&ele)
			printf("Index of %v is %v\n", ele, indexOf(intArray, ele))
		case 7:
			fmt.Println("array berfore reverse: ", intArray)
			intArray = reverseIntArray(intArray)
			println("reversed array: ", intArray)
		case 8:
			fmt.Println("array berfore sort: ", intArray)
			intArray = sortIntArrayInAsc(intArray)
			println("array after sort: ", intArray)
		case 9:
			fmt.Println("array berfore sort: ", intArray)
			intArray = sortIntArrayInDesc(intArray)
			println("array after sort: ", intArray)
		case 10:
			var ele int
			print("Enter the element to count: ")
			fmt.Scan(&ele)
			printf("count of %v in array is: %v", ele, getCountOfElement(intArray, ele))
		case 11:
			printf("Sum of %v is = %v", intArray, getSum(intArray))
		case 12:
			printf("Avarage of %v is = %v", intArray, getAVG(intArray))
		case 13:
			printf("max element in %v is = %v", intArray, getMax(intArray))
		case 14:
			printf("min element in %v is = %v", intArray, getMin(intArray))
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

// @note This method is to show how to declare different types of array
func declareArrays() {
	var intArr [5]int // way of declaring array
	intArr[2] = 10
	intArray := [6]int{1, 4, 6, 12, 11} // or
	printf("int array: %v, size of array is %v\n", intArray, len(intArray))

	strArr := [5]string{"Jan", "feb", "march"}
	printf("String array: %v\n", strArr)

	complexArr := [5]complex128{1 + 2i}
	printf("Complex Array: %v\n", complexArr)

	float64Array := [5]float64{1.2, 3.45}
	printf("float64 Array: %v\n", float64Array)
}

/*
 @note This method is to add element at specific index of array
 @param arr, index, element
 @return array
*/
func appendIntElement(arr [size]int, index, element int) [size]int {
	arr[index] = element
	return arr
}

/*
 @note This method is to remove element from given index of array
 @param arr, index
 @return array
*/
func removeIntElementByIndex(arr [size]int, index int) [size]int {
	arr[index] = 0
	return arr
}

/*
 @note This method is to remove the given element from given array
 @param arr, element
 @return array
*/
func removeIntElementByValue(arr [size]int, element int) [size]int {
	for index, val := range arr {
		if val == element {
			arr[index] = 0
			return arr
		}
	}
	return arr
}

/*
 @note This method is to get the index of given element from the given array
 @param arr, element
 @return index of element
*/
func indexOf(arr [size]int, element int) int {
	for index, val := range arr {
		if val == element {
			return index
		}
	}
	return -1
}

/*
 @note This method is to reverse the given array
 @param arr
 @return reversed array
*/
func reverseIntArray(arr [size]int) [size]int {

	for i := 0; i < size/2; i++ {
		arr[i], arr[size-1-i] = arr[size-1-i], arr[i]
	}
	return arr
}

/*
 @note This method is to sort the given array in ascending order
 @param arr
 @return sorted array
*/
func sortIntArrayInAsc(arr [size]int) [size]int {
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size-1; j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

/*
 @note This method is to sort the given array in descending order
 @param arr
 @return sorted array
*/
func sortIntArrayInDesc(arr [size]int) [size]int {
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size-1; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

/*
 @note This method is to get the count of given element from given array
 @param arr, elemet
 @return count of given element
*/
func getCountOfElement(arr [size]int, element int) int {
	count := 0
	for _, v := range arr {
		if v == element {
			count = count + 1
		}
	}
	return count
}

/*
 @note This method is to get sum of given int array
 @param arr
 @return sum of integers in array
*/
func getSum(arr [size]int) int {
	sum := 0
	for i := 0; i < len(arr)-1; i++ {
		sum = sum + arr[i]
	}
	return sum
}

/*
 @note This method is to get avarage of given int array
 @param arr
 @return avarage of integers in array
*/
func getAVG(arr [size]int) int {
	sum := 0
	for i := 0; i < len(arr)-1; i++ {
		sum = sum + arr[i]
	}
	return sum / len(arr)
}

/*
 @note This method is to get maximum element of given int array
 @param arr
 @return max value of integers in array
*/
func getMax(arr [size]int) int {
	max := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

/*
 @note This method is to get minimum element of given int array
 @param arr
 @return min value of integers in array
*/
func getMin(arr [size]int) int {
	min := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
