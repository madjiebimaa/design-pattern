package models

type EmployeeFactory struct {
}

func NewEmployeeFactory() EmployeeFactory {
	return EmployeeFactory{}
}

func (f *EmployeeFactory) NewSE(name string) Employee {
	return Employee{
		Name:   name,
		Title:  "SE",
		Salary: 10000000,
	}
}

func (f *EmployeeFactory) NewFE(name string) Employee {
	return Employee{
		Name:   name,
		Title:  "FE",
		Salary: 9000000,
	}
}
