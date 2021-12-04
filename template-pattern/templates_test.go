package templates_test

import (
	"fmt"
	"go-design-pattern/template-pattern/templates"
	"testing"
)

func TestTemplates(t *testing.T) {
	var square templates.Square
	square1 := square.New('O', 10, 10, "START", "END")
	square2 := square.New('X', 1, 1, "START", "END")
	square3 := square.New('-', 5, 5, "START", "END")

	fmt.Println(square1, square2, square3)
}
