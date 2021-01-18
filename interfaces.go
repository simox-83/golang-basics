package main

import (
	"fmt"
)

type VowelsFinder interface {
	FindVowels() []rune
}

type myString string

type SalaryCalculator interface {
	calculateSalary() int
}

type Permanent struct {
	empId    int
	basicPay int
	bonus    int
}

type Contractor struct {
	empId    int
	basicPay int
}

type Freelancer struct {
	hours int
	rate  int
}

func main() {

	s := myString("ciao")
	var v VowelsFinder
	describe(v)
	v = s
	describe(v)
	assert(v)
	fmt.Printf("Le vocali sono %c\n", v.FindVowels())
	p1 := Permanent{1, 2000, 500}
	p2 := Permanent{2, 1500, 300}
	c1 := Contractor{50, 2700}
	c2 := Contractor{51, 3000}
	f1 := Freelancer{120, 30}
	totalsalaries := []SalaryCalculator{p1, p2, c1, c2, f1}
	describe(totalsalaries)
	sum := totalExpense(totalsalaries)
	fmt.Println(sum)
	describe(sum)
	assert(sum)

	var i interface{} = 123
	assert(i)


}

func (s myString) FindVowels() []rune {
	var r []rune
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			r = append(r, c)
		}
	}

	return r
}

func (p Permanent) calculateSalary() int {
	return p.basicPay + p.bonus
}

func (c Contractor) calculateSalary() int {
	return c.basicPay
}

func (f Freelancer) calculateSalary() int {
	return f.rate * f.hours
}

func totalExpense(s []SalaryCalculator) int {
	amount := 0
	for _, n := range s {
		amount += n.calculateSalary()
	}
	return amount
}

func describe(i interface{}) {

	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{}) {

	v, ok := i.(myString)
	if !ok {
		fmt.Println("Tipo di dato non myString")
	} else {

		fmt.Println(v, ok)
	}
}
