package composite_test

import (
	"fmt"
	"go-design-pattern/composite-pattern/models"
	"testing"
)

func TestComposite(t *testing.T) {
	// categories := []models.Category{
	// 	models.CompositeCategory{Category: },
	// 	{Name: "Fashion"},
	// 	{Name: "Automotive"},
	// }

	categories := []models.Category{
		*models.NewCategory("adjie"),
		*models.NewCategory("bunga"),

		// 	*models.NewCompositeCategory("sakha", []models.Category{
		// 		*models.NewCategory("amar"),
		// 	}),
	}

	for _, category := range categories {
		fmt.Println(category.Name)
	}
}
