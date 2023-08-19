package main

import "fmt"

type Person interface {
	GetName() string
	GetAge() int
}

type Employee struct {
	Name    string
	Age     int
	Address Address
	Salary  float64
}

type Address struct {
	Street  string
	City    string
	Country string
}

func (e Employee) GetName() string {
	return e.Name
}

func (e Employee) GetAge() int {
	return e.Age
}

func main() {
	p := Employee{
		Name: "John",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			Country: "USA",
		},
		Salary: 50000.0,
	}

	PrintPerson(p)
}

func PrintPerson(p Person) {
	fmt.Println(p.GetName(), "is", p.GetAge(), "years old and lives at", p.(Employee).Address.Street+",", p.(Employee).Address.City+",", p.(Employee).Address.Country+".")
}
