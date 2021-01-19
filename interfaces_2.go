package main

import (
	"fmt"
)

type Describer interface {
	Describe()
}

type Citizener interface {
	PrintCountry()
}

type SalaryCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft()
}

type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}

type Person struct {
	name string
	age  int
}

type Address struct {
	country string
	city    string
}

type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func main() {
	p := Person{"Simone", 38}
	a := Address{"Italy", "Rome"}
	var d Describer
	d = p
	d.Describe()
	//d= a  Non compila perche' a e' di tipo Address e Address non implementa Describe (ma *Address si)
	d = &a
	d.Describe()
	var c Citizener = a
	c.PrintCountry()
	e := Employee{"Simone", "D'Andreta", 2000, 500, 25, 12}
	var eo EmployeeOperations = e
	eo.DisplaySalary()
	eo.CalculateLeavesLeft()

}

func (p Person) Describe() {
	fmt.Println(p)
}

func (a *Address) Describe() {
	fmt.Println(a)
}

func (a Address) PrintCountry() {
	fmt.Println("Country is", a.country)
}

func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d\n", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() {
	fmt.Printf("%s %s has %d days left", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
