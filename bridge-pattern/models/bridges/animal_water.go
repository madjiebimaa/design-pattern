package bridges

// Derive of Animal
type AnimalWater interface {
	Animal
	NumberOfFins() int
}
