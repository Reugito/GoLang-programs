package main

import (
	"fmt"
)

var printf = fmt.Printf
var println = fmt.Println
var scan = fmt.Scan

func main() {
	var choice int
	for {
		println("\n1. check if given string contains substring")
		println("2. check if given string contains a character")
		println("3. check if given string contains any character of substring")
		println("4. get count of substring from string")
		println("5. add two strings")
		println("6. join array of string with given character")
		println("7. Repeat the givent string")
		println("8. Replace the char from string")
		println("9. get index of given substring from string")
		println("10. split the givent string")
		println("11. convert the givent string in uppercase")
		println("12. convert the givent string in lowercase")
		println("13. reverse string")
		println("100. quit")
		printf("\nEnter the choice: ")
		scan(&choice)

		switch choice {
		case 1:
			str := getInputStr("string")
			pattern := getInputStr("pattern")
			printf("%v contains %v : %v\n", str, pattern, containsSubString(str, pattern))
		case 2:
			str := getInputStr("string")
			char := getInputChar("character")
			printf("%v contains char %v : %v\n", str, string(char), isContainsChar(str, char))
		case 3:
			str := getInputStr("string")
			subStr := getInputStr("SubString")
			printf("%v contains atleat one char of %v : %v\n", str, subStr, isContainsAny(str, subStr))
		case 4:
			str := getInputStr("string")
			subStr := getInputStr("SubString")
			printf("count of %v in %v = %v\n", str, subStr, getCount(str, subStr))
		case 5:
			str1 := getInputStr("string1")
			str2 := getInputStr("string2")
			printf("concatination of %v and %v = %v\n", str1, str2, addStrings(str1, str2))
		case 6:
			var strArr []string
			var size int
			printf("Entr size of int array: ")
			scan(&size)
			for i := 0; i < size-1; i++ {
				str := getInputStr("Enter string array element")
				strArr = append(strArr, str)
			}
			char := getInputStr("Joint with string or char")
			printf("%v join with %v = %v\n", strArr, char, joinStrings(strArr, char))
		case 7:
			str := getInputStr("string")
			var times int
			printf("Entr size of int array: ")
			scan(&times)
			printf("repeat %v %v times = %v\n", str, times, repeatString(str, times))
		case 8:
			str := getInputStr("String")
			char1 := getInputChar("old char")
			char2 := getInputChar("new char")
			printf("After replacing %v with %v in %v = %v\n", string(char1), string(char2), str, replaceString(str, char1, char2))
		case 9:
			str := getInputStr("string")
			subStr := getInputStr("SubString")
			printf("Index of %v in %v = %v\n", subStr, str, getIndex(str, subStr))
		case 10:
			str := getInputStr("string to split")
			char := getInputChar("char to split")
			printf("After splitting %v with %v we get %v\n", str, string(char), splitString(str, char))
		case 11:
			str := getInputStr("String")
			printf("%v in upperCase : %v", str, toUpper(str))
		case 12:
			str := getInputStr("String")
			printf("%v in lowerCase : %v\n", str, toLower(str))
		case 13:
			str := getInputStr("string to reverse")
			printf("reverse of %v = %v\n", str, reverseString(str))
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

/*
 @note this method is to take character input
 @param printStatement  :- statement to display in console
 @return input character
*/
func getInputChar(printStatement string) rune {
	var char rune
	printf("Enter %v: ", printStatement)
	scan(&char)
	return char
}

/*
 @note this method is to take string input
 @param printStatement  :- statement to display in console
 @return input string
*/
func getInputStr(printStatement string) string {
	var str string
	printf("Enter %v: ", printStatement)
	scan(&str)
	return str
}

/*
 @note this method is to check whether the given pattern is present in a given string
 @param  str, pattern
 @return boolean value depending on pattern, If pattern is present the true else false respectively
*/
func containsSubString(str, pattern string) bool {
	refIndex := 0
	for _, val := range str {

		if string(val) == string(pattern[refIndex]) {
			if refIndex < len(pattern)-1 {
				refIndex = refIndex + 1
			} else if refIndex == len(pattern)-1 {
				return true
			}
		} else {
			refIndex = 0
		}
	}
	return false
}

/*
 @note this method is to check whether the given character is present in a given string
 @param  str, pattern
 @return boolean value depending on character, If character is present the true else false respectively
*/
func isContainsChar(str string, char rune) bool {
	// return s.Contains(str, string(char))
	for _, val := range str {
		if val == char {
			return true
		}
	}
	return false
}

/*
 @note this method is to check whether the given pattern or characters of pattern are present in a given string
 @param  str, chars
 @return boolean value depending on pattern characters, If one of char is present the true else false respectively
*/
func isContainsAny(str, chars string) bool {
	// return s.ContainsAny(str, chars)
	for _, val := range chars {
		if isContainsChar(str, val) {
			return true
		}
	}
	return false
}

/*
 @note This method to get the count of a pattern present in a string
 @param str, pattern
 @return this method will return the count of pattern present in a string

*/
func getCount(str, pattern string) int {
	// return s.Count(str, pattern)
	refIndex := 0
	count := 0
	for _, val := range str {
		if string(val) == string(pattern[refIndex]) {
			if refIndex < len(pattern)-1 {
				refIndex = refIndex + 1
			} else if refIndex == len(pattern)-1 {
				count++
			}
		} else {
			refIndex = 0
		}
	}
	if count <= 0 {
		return -1
	}
	return count
}

/*
 @note this method is to concatinate two strings
 @param str1 , str1
 @return the concatenated string
*/
func addStrings(str1, str2 string) string {
	return str1 + str2
}

/*
 @note This method will join the array of string in one string with join character
 @param str , joinWith
 @return the concatenated string
*/
func joinStrings(str []string, joinWith string) string {
	// return s.Join(str, joinWith)
	joinedString := ""

	for index, val := range str {
		if index == len(str)-1 {
			joinedString = joinedString + val
		} else {
			joinedString = joinedString + val + joinWith
		}
	}
	return joinedString
}

/*
 @note this method will repeat a string n times depending on count
 @param str, count
 @return n times repeated string

*/
func repeatString(Str string, count int) string {
	// return s.Repeat(Str, count)
	repeatedString := ""

	for i := 0; i < count; i++ {
		repeatedString = repeatedString + Str
	}
	return repeatedString
}

/*
 @note   this method will replace a string with a new string
 @param  str, oldChar, newChar
 @return characters replaced string

*/
func replaceString(str string, oldchar, newchar rune) string {
	// return s.Replace(str, string(oldchar), string(newchar), 1)
	newString := ""
	for i := 0; i < len(str); i++ {
		if string(str[i]) == string(oldchar) {
			newString = newString + string(newchar)
		} else {
			newString = newString + string(str[i])
		}
	}
	return newString
}

/*
 @note This method will return a index of given string
 @param str, pattern
 @return  index of char

*/
func getIndex(str string, pattern string) int {
	// return s.Index(str, string(char))
	refIndex := 0
	for index, val := range str {

		if string(val) == string(pattern[refIndex]) {
			if refIndex < len(pattern)-1 {
				refIndex = refIndex + 1
			} else if refIndex == len(pattern)-1 {
				return index - refIndex
			}
		} else {
			refIndex = 0
		}
	}
	return -1
}

/*
 @note This method will split a string into a array of strings
 @param str, char
 @return  array of strings

*/
func splitString(str string, char rune) []string {
	// return s.Split(str, string(char))
	subString := ""
	var splitStringArr []string
	for _, val := range str {
		if val == char {
			splitStringArr = append(splitStringArr, subString)
			subString = ""
		} else {
			subString = subString + string(val)
		}
	}
	splitStringArr = append(splitStringArr, subString)
	return splitStringArr
}

/*
 @note This method will replace lower case alphabates to upper case
 @param str
 @return  uppercase string

*/
func toUpper(str string) string {
	// return s.ToUpper(str)
	newString := ""
	for _, val := range str {
		if val <= 122 && val >= 97 {
			newString = newString + string(val-32)
		} else {
			newString = newString + string(val)
		}
	}
	return newString
}

/*
 @note This method will replace upper case alphabates to lower case
 @param str
 @return  lowercase string

*/
func toLower(str string) string {
	// return s.ToUpper(str)
	newString := ""
	for _, val := range str {
		if val <= 90 && val >= 65 {
			newString = newString + string(val+32)
		} else {
			newString = newString + string(val)
		}
	}
	return newString
}

/*
 @note This method is to reverse the given string
 @param str
 @return reversed string
*/
func reverseString(str string) string {
	newString := ""

	for _, val := range str {
		newString = string(val) + newString
	}
	return newString
}
