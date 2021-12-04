package models

// Derive of Cate
type CompositeCategory struct {
	Category
	Categories []Category
}

func NewCompositeCategory(name string, categories []Category) *CompositeCategory {
	return &CompositeCategory{
		Category:   *NewCategory(name),
		Categories: categories,
	}
}
