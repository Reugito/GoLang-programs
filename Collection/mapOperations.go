package main

import (
	"fmt"
)

var println, printf, scan = fmt.Println, fmt.Printf, fmt.Scan

func main() {
	animalsMap := map[interface{}]interface{}{}

	for {
		var choice int

		println("\n1. show animalsMap map")
		println("2. lenth of animalNameAndTypes map")
		println("3. Add Animal and Its type ")
		println("4. Add multiple Animal and Its type in ")
		println("5. update animal type by animal name")
		println("6. remove animal by name")
		println("7. sort animalMap ")
		println("8. count animal by animal type")
		println("9. Get animal By name")
		println("10. Get animals by type")
		println("100. Quit")

		printf("\nEnter Your Choice: ")
		scan(&choice)

		switch choice {
		case 1:
			displayMap(animalsMap)
		case 2:
			println("lenth of animalNameAndTypes map: ", len(animalsMap))
		case 3:
			animalsMap = addAnimalAndType(animalsMap, strInput("animal to add"), strInput("type of animal"))
		case 4:
			animalsMap = addMultipleAnimals(animalsMap, getMultipleAnimal())
		case 5:
			animalsMap = updateAnimalType(animalsMap, strInput("Animal name"), strInput("animal Type"))
		case 6:
			animalsMap = removeAnimalByName(animalsMap, strInput("Animal name to remove"))
		case 7:
			// sortAnimalMapByName(animalsMap)
		case 8:
			animalType := strInput("animal type to count")
			printf("\nThere are %v animals of type %v in animalMap\n", countAnimalsByType(animalsMap, animalType), animalType)
		case 9:
			aName, aType := getAnimalByName(animalsMap, strInput("Animal name"))
			printf("\n1.%v --- %v\n", aName, aType)
		case 10:
			newMap := getAllAnimalsByType(animalsMap, strInput("Animal Type"))
			displayMap(newMap)
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

// This method is to show ways to declare map
func waysToDeclareMaps() {
	dict := map[string]string{} // empty mapp
	println("Empty declaration : ", dict)

	dict1 := make(map[string]string) // by using make function
	println("Map declaration with make function : ", dict1)

	var dict2 = map[string]string{"rao": "dhotre"}
	println("map : ", dict2)
}

/*
 @note This method is to take string input from user
 @param  printStatement
 @return inputValue
*/
func strInput(printStatement string) string {
	inputVal := ""
	printf("Enter %v : ", printStatement)
	_, err := fmt.Scan(&inputVal)
	if err != nil {
		return inputVal
	}
	return inputVal
}

/*
 @note This method is to take animal name and animal type as input and convert it into map
 @return map containing animal name and animal type
*/
func getMultipleAnimal() map[interface{}]interface{} {
	size := 0
	printf("Enter no of animals want to add: ")
	scan(&size)

	newMap := map[interface{}]interface{}{}
	for i := 0; i < size; i++ {
		newMap[strInput("Animal Name")] = strInput("Animal Type")
	}
	return newMap
}

/*
 @note This method is do display the animal name with type on console
 @param animalMap :- contains animal name and type
*/
func displayMap(animalsMap map[interface{}]interface{}) {
	srNo := 0
	println()
	for k, v := range animalsMap {
		srNo = srNo + 1
		printf(" %v. %v --- %v\n", srNo, k, v)
	}
}

/*
 @note This method is to add given animal name and type in animalsMap dictionary
 @param animalsMap, animalName, animalType
 @return animalsMap containing animal name and types
*/
func addAnimalAndType(animalsMap map[interface{}]interface{}, animal string, animalType string) map[interface{}]interface{} {
	animalsMap[animal] = animalType
	return animalsMap
}

/*
 @note This method is to add multiple animal name and types in animalsMap dictionary at once
 @param animalsMap, newAnimals
 @return animalsMap containing animal name and types
*/
func addMultipleAnimals(animalsMap map[interface{}]interface{}, newAnimals map[interface{}]interface{}) map[interface{}]interface{} {
	for animalName, animalType := range newAnimals {
		animalsMap[animalName] = animalType
	}
	return animalsMap
}

/*
 @note This method is to update given animal name  with animal type in animalsMap dictionary
 @param animalsMap, animalName, animalType
 @return animalsMap containing animal name and types
*/
func updateAnimalType(animalsMap map[interface{}]interface{}, animalName, animalType string) map[interface{}]interface{} {
	animalsMap[animalName] = animalType
	return animalsMap
}

/*
 @note This method is to remove given animal name from animalsMap dictionary
 @param animalsMap, animalName
 @return animalsMap containing animal name and types
*/
func removeAnimalByName(animalsMap map[interface{}]interface{}, animalName string) map[interface{}]interface{} {
	delete(animalsMap, animalName)
	return animalsMap
}

/*
 @note This method is to count the given animal types from animalsMap dictionary
 @param animalsMap, animalType
 @return count of given animal type from animalMap
*/
func countAnimalsByType(animalsMap map[interface{}]interface{}, animalType string) int {
	count := 0
	for _, t := range animalsMap {
		if t == animalType {
			count = count + 1
		}
	}
	return count
}

/*
 @note This method is to get animal details from animalsMap dictionary by animal name
 @param animalsMap, animalName
 @return animal name and animal type
*/
func getAnimalByName(animalsMap map[interface{}]interface{}, animalName string) (interface{}, interface{}) {

	for k, v := range animalsMap {
		if k == animalName {
			return k, v
		}
	}
	return "", ""
}

/*
 @note This method is to get animal details from animalsMap dictionary which contains given animal type
 @param animalsMap, animalType
 @return map containing animal details with same animal types
*/
func getAllAnimalsByType(animalsMap map[interface{}]interface{}, animalType string) map[interface{}]interface{} {
	newMap := map[interface{}]interface{}{}

	for k, v := range animalsMap {
		if v == animalType {
			newMap[k] = v
		}
	}
	return newMap
}
