package Composite

import "fmt"

type Employee struct {
	name         string
	dept         string
	salary       int
	subordinates []*Employee
}

func newEmployee(name string, dept string, salary int) *Employee {
	return &Employee{
		name,
		dept,
		salary,
		make([]*Employee, 0),
	}
}

func (e Employee) Name() string {
	return e.name
}

func (e Employee) Salary() int {
	return e.salary
}

func (e *Employee) Add(employees ...*Employee) {
	e.subordinates = append(e.subordinates, employees...)
}

func (e *Employee) Remove(employee *Employee) {
	for i, item := range e.subordinates {
		if item.Name() == employee.name {
			e.subordinates = append(e.subordinates[:i], e.subordinates[i+1:]...)
		}
	}
}

func (e *Employee) GetSubordinates() []*Employee {
	return e.subordinates
}

func (e Employee) String() string {
	return fmt.Sprintf("Name{%s}, Dept{%s}, Salary{%d}", e.name, e.dept, e.salary)
}


