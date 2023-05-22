package exercises

type Medusa struct {
	Name    string
	Statues []string
}

func NewMedusa(name string) Medusa {
	medusa := Medusa{
		Name:    name,
		Statues: []string{},
	}
	return medusa
}
