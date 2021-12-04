package models

type Category struct {
	Name string
}

func NewCategory(name string) *Category {
	return &Category{Name: name}
}
