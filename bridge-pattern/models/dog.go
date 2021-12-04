package models

import "go-design-pattern/bridge-pattern/models/bridges"

type Dog struct {
}

func NewDog() bridges.AnimalLand {
	return &Dog{}
}

func (w *Dog) GetName() string {
	return "Dog"
}

func (w *Dog) IsAnimalWater() bool {
	return false
}

func (w *Dog) IsAnimalLand() bool {
	return true
}

func (w *Dog) NumberOfLegs() int {
	return 4
}
