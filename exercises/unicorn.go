package exercises

import "fmt"

type Unicorn struct {
	Name  string
	Color string
}

func NewUnicorn(name string, color string) Unicorn {
	unicorn := Unicorn{
		Name:  name,
		Color: color,
	}

	if color != "" {
		unicorn.Color = color
	} else {
		unicorn.Color = "white"
	}
	return unicorn
}

func (u Unicorn) IsWhite() bool {
	return u.Color == "white"
}

func (u Unicorn) Says(input string) string {
	return fmt.Sprintf("**;* %s *;**", input)
}
