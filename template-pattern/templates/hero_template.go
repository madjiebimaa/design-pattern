package templates

import "fmt"

type Square struct {
	Char       byte
	Width      uint
	Height     uint
	TitleBegin string
	TitleEnd   string
}

func (s *Square) New(char byte, width uint, height uint, titleBegin string, titleEnd string) *Square {
	fmt.Println(titleBegin)

	for i := 0; i < int(width); i++ {
		for i := 0; i < int(height); i++ {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}

	fmt.Println(titleEnd)

	return &Square{
		Char:       char,
		Width:      width,
		Height:     height,
		TitleBegin: titleBegin,
		TitleEnd:   titleEnd,
	}
}
