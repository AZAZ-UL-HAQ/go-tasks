package main

import (
	"fmt"
)

type employee struct {
	name string

	salary int

	position string
}

type company struct {
	companyName string

	employees []employee
}

func main() {

	Employee1 := employee{
		name:     "AZAZ",
		salary:   1000,
		position: "Army-Cheif",
	}

	Employee2 := employee{
		name:     "Noman",
		salary:   1400,
		position: "Sec.Lt",
	}
	Employee3 := employee{
		name:     "Sohaib",
		salary:   5000,
		position: "HOD",
	}

	emplys := []employee{Employee1, Employee2, Employee3}
	Company := company{
		companyName: "Haq-Brothers",
		employees:   emplys,
	}
	fmt.Printf("Compnay Name is : %s  \n", Company.companyName)
	for i := range Company.employees {
		fmt.Printf("Name: %s\n", Company.employees[i].name)
		fmt.Printf("Salary: %d\n", Company.employees[i].salary)
		fmt.Printf("Position: %s\n\n", Company.employees[i].position)
	}

}
