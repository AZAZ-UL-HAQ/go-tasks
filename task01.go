package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func getName(Pers1 Person, Pers2 Person) {
	fmt.Print(Pers1)
	fmt.Print(Pers2)
}
func main() {
	Pers1 := Person{
		name:   "Hege",
		age:    45,
		job:    "Teacher",
		salary: 6000,
	}

	Pers2 := Person{
		name:   "AZAZ",
		age:    22,
		job:    "NASA",
		salary: 000,
	}
	getName(Pers1, Pers2)
}