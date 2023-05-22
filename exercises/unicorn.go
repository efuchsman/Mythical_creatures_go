package exercises

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
