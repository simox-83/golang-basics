package employee

import "fmt"

//Employee defines an employee
type employee struct {
	name        string
	totalLeaves int
	takenLeaves int
}

//RemainingLeaves returns the number of free days left
func (e employee) RemainingLeaves() {
	fmt.Println(e.name, "has", e.totalLeaves-e.takenLeaves, "days remaining")
}

//Class constructor to avoid empty structs to be initialized uncorrectly
func NewEmployee(name string, total int, taken int) employee {
	return employee{
		name,
		total,
		taken,
	}
}
