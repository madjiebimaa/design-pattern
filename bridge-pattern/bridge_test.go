package bridge_test

import (
	"fmt"
	"go-design-pattern/bridge-pattern/models"
	"go-design-pattern/bridge-pattern/models/bridges"
	"strconv"
	"testing"
)

func TestBridge(t *testing.T) {
	animals := []bridges.Animal{
		models.NewCat(),
		models.NewDog(),
		models.NewFish(),
		models.NewWhale(),
	}

	for i := 0; i < len(animals); i++ {
		animalLand, ok := animals[i].(bridges.AnimalLand)
		if ok {
			fmt.Println(animalLand.GetName() + " Has " + strconv.Itoa((animalLand.NumberOfLegs())) + " Legs")
		}

		animalWater, ok := animals[i].(bridges.AnimalWater)
		if ok {
			fmt.Println(animalWater.GetName() + " Has " + strconv.Itoa((animalWater.NumberOfFins())) + " Fins")
		}
	}
}
