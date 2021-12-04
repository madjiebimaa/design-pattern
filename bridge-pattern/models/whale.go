package models

import "go-design-pattern/bridge-pattern/models/bridges"

type Whale struct {
}

func NewWhale() bridges.AnimalWater {
	return &Whale{}
}

func (w *Whale) GetName() string {
	return "Whale"
}

func (w *Whale) IsAnimalWater() bool {
	return true
}

func (w *Whale) IsAnimalLand() bool {
	return false
}

func (w *Whale) NumberOfFins() int {
	return 2
}
