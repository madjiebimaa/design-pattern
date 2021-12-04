package models

import "go-design-pattern/bridge-pattern/models/bridges"

type Fish struct {
}

func NewFish() bridges.AnimalWater {
	return &Fish{}
}

func (w *Fish) GetName() string {
	return "Fish"
}

func (w *Fish) IsAnimalWater() bool {
	return true
}

func (w *Fish) IsAnimalLand() bool {
	return false
}

func (w *Fish) NumberOfFins() int {
	return 4
}
