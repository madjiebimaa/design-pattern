package models

import "go-design-pattern/bridge-pattern/models/bridges"

type Cat struct {
}

func NewCat() bridges.AnimalLand {
	return &Cat{}
}

func (w *Cat) GetName() string {
	return "Cat"
}

func (w *Cat) IsAnimalWater() bool {
	return false
}

func (w *Cat) IsAnimalLand() bool {
	return true
}

func (w *Cat) NumberOfLegs() int {
	return 4
}
