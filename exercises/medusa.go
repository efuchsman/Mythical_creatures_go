package exercises

type Medusa struct {
	Name    string
	Statues []Person
}

func NewMedusa(name string) Medusa {
	medusa := Medusa{
		Name:    name,
		Statues: []Person{},
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
	m.Statues = append(m.Statues, *p)
	p.Stoned = true

	if len(m.Statues) > 3 {
		removedPerson := &m.Statues[0]
		m.Statues = m.Statues[1:]
		removedPerson.Stoned = false
	}
}
