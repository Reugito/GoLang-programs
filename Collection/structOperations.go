package main

import (
	"encoding/json"
	"fmt"
)

var println, printf, scan = fmt.Println, fmt.Printf, fmt.Scan

type empDetails struct {
	EmpId    int    `json:"empId"`
	EmpName  string `json:"empName"`
	JobTitle string `json:"Job"`
}

// This struct is to to display the struct containt in json format
type Employee struct { // the struct name and all attributes should be start with capital letter
	EmpId    int    `json:"empId"`
	Name     string `json:"name"`
	JobTitle string `json:"Job"`
}

var empDB []empDetails

func main() {

	var choice int
	for {
		println("\n1. show array")
		println("13. show array in JsonFormat")
		println("2. show size of DB")
		println("3. add new Employee")
		println("4. update Employee Name by empID")
		println("5. update Employee JobTitle by empID")
		println("6. delete employee by empId")
		println("7. sort employee by id in Ascending order")
		println("8. sort employee by id in Descending order")
		println("9. sort employee by name in Ascending order")
		println("10. sort employee by name in Descending order")
		println("11. get employeeDetails by name")
		println("12. get employeeDetails by JobTitle")
		println("100. Quit")

		printf("\nEnter the choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			displayEmployees(empDB)
		case 13:
			jsonData := convertToJSONFormat(empDB)
			printf("%s \n", jsonData)
			displayEmployees(converBackFromJSONFormat(jsonData))
		case 2:
			printf("The size of DB = %v\n", len(empDB))
		case 3:
			empID := intInput(" Employee ID")
			empName := strInput("new Employee name")
			empJobTitle := strInput("new Employee job Title")
			empDB = addEmployee(empDB, empID, empName, empJobTitle)
		case 4:
			empID := intInput(" Employee ID to update")
			empName := strInput("new Employee name")
			empDB = uppdateEmployeeNameByEmpID(empDB, empID, empName)
		case 5:
			empID := intInput(" Employee ID to update")
			jobTitle := strInput("new Job title")
			empDB = updateEmployeeJobTitleByEmpID(empDB, empID, jobTitle)
		case 6:
			empID := intInput(" Employee ID to delete")
			empDB = deleteEmployeeById(empDB, empID)
		case 7:
			empDB = sortEmployeesByEmpIdASC(empDB)
			displayEmployees(empDB)
		case 8:
			empDB = sortEmployeesByEmpIdDESC(empDB)
			displayEmployees(empDB)
		case 9:
			empDB = sortEmployeesByNameASC(empDB)
			displayEmployees(empDB)
		case 10:
			empDB = sortEmployeesByNameDESC(empDB)
			displayEmployees(empDB)
		case 11:
			displayEmployees(getEmployeeDetailsByName(empDB, strInput("empName")))
		case 12:
			displayEmployees(getEmployeeDetailsByJobTitle(empDB, strInput("empJobTitle")))
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

func waysOfCreatingStructInstance() {
	var emp1 empDetails
	emp1.EmpId = 1001
	emp1.EmpName = "Rao"
	emp1.JobTitle = "Developer"
	println("1st Emp Instance: ", emp1)

	var emp2 = empDetails{EmpId: 1002, EmpName: "Shankar"}
	println("2nd Emp Instance: ", emp2)

	emp3 := empDetails{EmpName: "Shankar", JobTitle: "HR"}
	println("3rd Emp Instance: ", emp3)

	var emp4 = new(empDetails) //emp5 pointer to the instance
	emp4.EmpId = 1003
	println("4th Emp Instance: ", emp4)

	emp5 := new(empDetails) //emp5 pointer to the instance
	emp5.EmpId = 1004
	emp5.EmpName = "kalyan"
	println("5th Emp Instance: ", emp5)

	emp6 := &empDetails{EmpId: 1002, EmpName: "Shankar"}
	println("6th Emp Instance: ", emp6)

	var emp7 = Employee{1009, "Raghav", "Raj"}
	jsonStr, _ := json.Marshal(emp7)
	printf("%s", jsonStr)
}

func (emp empDetails) greetingMSG() {
	printf(" --------- Welcome  %v --------- ", emp.EmpName)
}

/*
 @note This method is to take string input from user
 @param  printStatement
 @return inputValue
*/
func strInput(printStatement string) string {
	var strVal string
	printf("Enter %v: ", printStatement)
	scan(&strVal)
	return strVal
}

/*
 @note This method is to take int input from user
 @param  printStatement
 @return inputValue
*/
func intInput(printStatement string) int {
	intVal := -1
	for {
		printf("Enter %v: ", printStatement)
		_, err := scan(&intVal)
		if err != nil {
			println("You Entered the Wrong input")
		} else {
			return intVal
		}
	}
}

/*
 @note This method is do display the the content of empDB on console
 @param empDB :- List of empId, empName, jobTitle
*/
func displayEmployees(empDB []empDetails) {
	for k, v := range empDB {
		printf("\n%v. \n empId: %v\n empName: %v\n jobTitle: %v\n", k+1, v.EmpId, v.EmpName, v.JobTitle)
	}
}

/*
 @note This method is to add given employee details in empDB
 @param empDB, empID, empName, JobTitle
 @return empDB List of employee details i.e. empId, empName, jobTitle
*/
func addEmployee(empDB []empDetails, empId int, empName, jobTitle string) []empDetails {
	ispresent := true
	for _, v := range empDB {
		if v.EmpId == empId {
			ispresent = false
			break
		}
	}
	if ispresent {
		newEmp := empDetails{empId, empName, jobTitle}
		newEmp.responseToAddEmployeeCall()
		return append(empDB, newEmp)
	}
	return empDB
}

/*
 @note This method is to update the employee name with given employee ID
 @param empDB, empID, empName
 @return empDB List of employee details i.e. empId, empName, jobTitle
*/
func uppdateEmployeeNameByEmpID(empDB []empDetails, empId int, empName string) []empDetails {
	var index int
	isEmployeeIDPresent := false
	var employee empDetails
	for i, v := range empDB {
		if v.EmpId == empId {
			index = i
			employee = v
			isEmployeeIDPresent = true
			break
		}
	}
	if isEmployeeIDPresent {
		employee.EmpName = empName
		empDB[index] = employee
	}
	return empDB
}

/*
 @note This method is to update the employee Job title with given employee ID
 @param empDB, empID, jobTitle
 @return empDB List of employee details i.e. empId, empName, jobTitle
*/
func updateEmployeeJobTitleByEmpID(empDB []empDetails, empId int, jobTitle string) []empDetails {

	var index int
	isEmployeeIDPresent := false
	var employee empDetails
	for i, v := range empDB {
		if v.EmpId == empId {
			index = i
			employee = v
			isEmployeeIDPresent = true
			break
		}
	}
	if isEmployeeIDPresent {
		employee.JobTitle = jobTitle
		empDB[index] = employee
	}
	return empDB
}

/*
 @note This method is to delete the employee details with given employee ID
 @param empDB, empID
 @return empDB List of employee details i.e. empId, empName, jobTitle
*/
func deleteEmployeeById(empDB []empDetails, empId int) []empDetails {

	var index int
	isEmployeeIDPresent := false
	for i, v := range empDB {
		if v.EmpId == empId {
			index = i
			isEmployeeIDPresent = true
			break
		}
	}
	if isEmployeeIDPresent {
		empDB[index].responseToDeletedEmployeeCall()
		return append(empDB[:index], empDB[index+1:]...)
	}
	return empDB
}

/*
 @note This method is to sort the employee details from empDB in ascending order
 		with respect to empID
 @param empDB
 @return sorted empDB List of employee details i.e. empId, empName, jobTitle
*/
func sortEmployeesByEmpIdASC(empDB []empDetails) []empDetails {
	for i := 0; i < len(empDB); i++ {
		for j := i + 1; j < len(empDB); j++ {
			if empDB[i].EmpId > empDB[j].EmpId {
				empDB[i], empDB[j] = empDB[j], empDB[i]
			}
		}
	}
	return empDB
}

/*
 @note This method is to sort the employee details from empDB in descending order
 		with respect to empID
 @param empDB
 @return sorted empDB List of employee details i.e. empId, empName, jobTitle
*/
func sortEmployeesByEmpIdDESC(empDB []empDetails) []empDetails {
	for i := 0; i < len(empDB); i++ {
		for j := i + 1; j < len(empDB); j++ {
			if empDB[i].EmpId < empDB[j].EmpId {
				empDB[i], empDB[j] = empDB[j], empDB[i]
			}
		}
	}
	return empDB
}

/*
 @note This method is to sort the employee details from empDB in ascending order
 		with respect to empName
 @param empDB
 @return sorted empDB List of employee details i.e. empId, empName, jobTitle
*/
func sortEmployeesByNameASC(empDB []empDetails) []empDetails {
	for i := 0; i < len(empDB); i++ {
		for j := i + 1; j < len(empDB); j++ {
			if empDB[i].EmpName > empDB[j].EmpName {
				empDB[i], empDB[j] = empDB[j], empDB[i]
			}
		}
	}
	return empDB
}

/*
 @note This method is to sort the employee details from empDB in descending order
 		with respect to empName
 @param empDB
 @return sorted empDB List of employee details i.e. empId, empName, jobTitle
*/
func sortEmployeesByNameDESC(empDB []empDetails) []empDetails {
	for i := 0; i < len(empDB); i++ {
		for j := i + 1; j < len(empDB); j++ {
			if empDB[i].EmpName < empDB[j].EmpName {
				empDB[i], empDB[j] = empDB[j], empDB[i]
			}
		}
	}
	return empDB
}

func getEmployeeDetailsByName(empDB []empDetails, empName string) []empDetails {
	var empWithSameName []empDetails

	for _, v := range empDB {
		if v.EmpName == empName {
			empWithSameName = append(empWithSameName, v)
		}
	}
	return empWithSameName
}

func getEmployeeDetailsByJobTitle(empDB []empDetails, jobTitle string) []empDetails {
	var empWithSameJobTitle []empDetails

	for _, v := range empDB {
		if v.JobTitle == jobTitle {
			empWithSameJobTitle = append(empWithSameJobTitle, v)
		}
	}
	return empWithSameJobTitle
}

//@note This method if for printing msg on console after adding new employee details in DB
// func =  keyword  receiver identifire (params) return type
func (emp empDetails) responseToAddEmployeeCall() {
	printf("\n %v added to DB with %v empId\n", emp.EmpName, emp.EmpId)
}

//@note This method if for printing msg on console after deleting employee details from DB
func (emp empDetails) responseToDeletedEmployeeCall() {
	printf("\n %v deleted from DB with %v empId\n", emp.EmpName, emp.EmpId)
}

func convertToJSONFormat(empDB []empDetails) []byte {
	var jsonData []byte
	jsonData, err := json.Marshal(empDB)
	if err != nil {
		return jsonData
	} else {
		return jsonData
	}
}

func converBackFromJSONFormat(jsonData []byte) []empDetails {
	var empDB []empDetails
	err := json.Unmarshal(jsonData, &empDB)
	if err != nil {
		return empDB
	} else {
		return empDB
	}
}
