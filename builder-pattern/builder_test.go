package builder_test

import (
	"fmt"
	"go-design-pattern/builder-pattern/models"
	"testing"
)

func TestBuilder(t *testing.T) {
	user1 := models.NewUser().SetFirstName("Adjie").SetLastName("Bima").Build()
	user2 := models.NewUser().SetFirstName("Raisa").SetLastName("Putri").SetAge(20).Build()
	user3 := models.NewUser().SetFirstName("Shinta").SetLastName("Nad").SetGender("F").Build()
	user4 := models.NewUser().SetFirstName("Sakha").SetLastName("Ammar").SetPassword("RAHASIA").Build()

	users := []models.User{
		*user1,
		*user2,
		*user3,
		*user4,
	}

	for _, user := range users {
		fmt.Println(user)
	}
}
