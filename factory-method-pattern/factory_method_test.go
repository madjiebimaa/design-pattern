package factory_method__test

import (
	"fmt"
	"go-design-pattern/factory-method-pattern/models"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	// employee1 := models.Employee{
	// 	Name:   "Adjie",
	// 	Title:  "SE",
	// 	Salary: 10000000,
	// }
	// employee2 := models.Employee{
	// 	Name:   "HAKIM",
	// 	Title:  "SE",
	// 	Salary: 10000000,
	// }
	// employee3 := models.Employee{
	// 	Name:   "Maxdha",
	// 	Title:  "FE",
	// 	Salary: 9000000,
	// }
	// employee4 := models.Employee{
	// 	Name:   "Tio",
	// 	Title:  "FE",
	// 	Salary: 9000000,
	// }
	employee := models.NewEmployeeFactory()

	employee1 := employee.NewSE("Adjie")
	employee2 := employee.NewSE("Hakim")
	employee3 := employee.NewFE("Maxdha")
	employee4 := employee.NewFE("Tio")

	empoyees := []models.Employee{
		employee1,
		employee2,
		employee3,
		employee4,
	}

	for _, employee := range empoyees {
		fmt.Println(employee)
	}
}
