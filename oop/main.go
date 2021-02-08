package main

import "oop/employee"

func main() {

	//e := employee.employee{Name: "Simone", TotalLeaves: 25, TakenLeaves: 10}
	//e.RemainingLeaves()

	//var n employee.employee
	//n.RemainingLeaves() // restituisce 0.. sintatticamente e' corretto, ma semanticamente no

	n1 := employee.NewEmployee("Ciccio", 25, 22)
	n1.RemainingLeaves()
}
