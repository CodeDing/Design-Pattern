package Composite

import "testing"

func TestComposite(t *testing.T) {
	CEO := newEmployee("John", "CEO", 30000)
	headSales := newEmployee("Robert", "Head Sales", 20000)
	headMarketing := newEmployee("Michael", "Head Marketing", 20000)
	CEO.Add(headSales, headMarketing)
	clerk1 := newEmployee("Laura", "Marketing", 10000)
	clerk2 := newEmployee("Bob", "Marketing", 10000)
	headMarketing.Add(clerk1, clerk2)
	sales1 := newEmployee("Rob", "Sales", 10000)
	sales2 := newEmployee("Richard", "Sales", 10000)
	headSales.Add(sales1, sales2)

	t.Logf("CEO: %s", CEO)
	for _, headEmployee := range CEO.GetSubordinates() {
		t.Logf("HeadEmployee: %s", headEmployee)
		for _, employee := range headEmployee.GetSubordinates() {
			t.Logf("Employee: %s", employee)
		}
	}
}
