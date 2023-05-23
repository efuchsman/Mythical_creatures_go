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

type Person struct {
	Medusa
	Name   string
	Stoned bool
}

func NewPerson(name string) Person {
	person := Person{
		Name:   name,
		Stoned: false,
	}
	return person
}

func (m *Medusa) Stare(p *Person) {
	m.Statues = append(m.Statues, p.Name)
	p.Stoned = true
}
