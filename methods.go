package main

import (
	"fmt"
)

type Employee struct {
	name   string
	salary int
	age    int
	Address
}

type Address struct {
	city    string
	country string
}

type Rectangle struct {
	length float64
	width  float64
}

type myType int

func main() {
	e1 := Employee{"Simone", 2000, 38,
		Address{"Rome", "Italy"},
	}
	e1.displaySalary()
	displaySalary(e1)

	e1.changeName("Tereza")
	fmt.Println("Contenuto di e1 dopo la chiamata senza puntatore", e1)
	(&e1).changeNamePointer("Simon") // e1.changeNamePointer va bene ugualmente
	fmt.Println("Contenuto di e1 dopo la chiamata con puntatore", e1)

	e1.displayAddress() // posso chiamare displayAddress su Employee, anche se il metodo e' definito sul tipo Address

	r := Rectangle{5.0, 4.2}
	fmt.Println("Area of r is", r.area())
	fmt.Println("Area of r is", area(r))
	var p *Rectangle
	p = &r

	fmt.Println("Area of p is", p.area()) // compila comunque, anche se l'oggetto chiamante e' *p, mentre il metodo e' definito su p
	//fmt.Println("Area of p is", area(p))  non compila perche' l'argomento del chiamante e *p, mentre la funzione vuole p

	a := myType(10)
	b := myType(5)
	fmt.Println("La somma e'", a.add(b))
}

func (e Employee) displaySalary() {
	fmt.Println("Salary of employee is", e.salary)
}

func displaySalary(e Employee) {
	fmt.Println("Salary of employee is", e.salary)
}

func (e Employee) changeName(newName string) {
	e.name = newName
}

func (e *Employee) changeNamePointer(newName string) {
	e.name = newName
}

func (a Address) displayAddress() {
	fmt.Println("L'indirizzo e'", a)
}

func (r Rectangle) area() float64 {
	return r.width * r.length
}

func area(r Rectangle) float64 {
	return r.width * r.length
}

func (t myType) add(b myType) myType {
	return t + b
}
