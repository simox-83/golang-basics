package main

import (
	"fmt"
)

type Employee struct {
	name   string
	salary int
	age    int
}

type Address struct {
	city    string
	country string
}

type Person struct {
	name    string
	address Address
}

func main() {

	e := Employee{name: "Simone",
		salary: 1000,
		age:    38,
	}
	fmt.Printf("%v\n", e)

	e = Employee{"Simone", 1500, 38}
	fmt.Printf("%v\n", e)

	//e = Employee{"Simone"} Non si puo'fare - con il literal bisogna specificare tutti i campi

	as := struct {
		kind  string
		value int
	}{
		kind:  "anonymous struct",
		value: 1,
	}
	fmt.Println("Anonymous struct printed", as)
	fmt.Printf("%T\t%T\n", e, as)

	fmt.Printf("Salary of employee %s is %d\n", e.name, e.salary)
	e.salary = 2000
	fmt.Printf("Salary of employee %s is %d\n", e.name, e.salary)

	var e1 Employee
	fmt.Println(e1)
	e2 := Employee{name: "Nuovo employee"}
	fmt.Println(e2)

	e3 := &Employee{"Puntatore Employee", 3000, 20}
	fmt.Println(*e3)
	e3.name = "Dereferenzio il puntatore"
	fmt.Println(e3.name)

	p := Person{name: "Simone",
		address: Address{
			city:    "Rome",
			country: "Italy",
		},
	}
	fmt.Println(p)
	p.address.city = "Prague"
	p.address.country = "Czechia"
	fmt.Println(p)

	if e1 == e2 {
		fmt.Println("e1 e2 are equal")
	} else {
		fmt.Println("e1 e2 are not equal")
	}

	modifyStruct(&e1)
	fmt.Println(e1)

}

func modifyStruct(e *Employee) {
	*e = Employee{"from the function", 5000, 45}

}
