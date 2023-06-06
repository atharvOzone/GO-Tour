package main

import "fmt"

type employee struct {
	salary int
	country string
}

func main() {
	emp1 := employee{
		salary: 1200,
		country: "USA",
	}
	emp2 := employee {
		salary: 1400,
		country: "India",
	}

	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Jamie": emp2,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary: $%d Country: %s\n", name, info.salary, info.country)
	}
}