package main

import "fmt"

// Employee struct represents an employee record
type Employee struct {
	Name     string
	Position string
	Salary   float64
}

// UpdateSalary updates the salary of an employee
func UpdateSalary(emp *Employee, newSalary float64) {
	fmt.Println(emp)
	emp.Salary = newSalary
}

func main() {
	// Creating an instance of Employee
	emp := Employee{
		Name:     "Raghul M",
		Position: "Software Quality Engineer",
		Salary:   25000,
	}

	// Printing initial state
	fmt.Println("Before update:", emp)

	// Updating salary
	UpdateSalary(&emp, 70000)

	// Printing updated state
	fmt.Println("After update:", emp)
}