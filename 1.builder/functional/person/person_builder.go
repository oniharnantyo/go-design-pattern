package person

type PersonMod func(*Person)
type PersonBuilder struct {
	Actions []PersonMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.Actions = append(b.Actions, func(person *Person) {
		person.Name = name
	})

	return b
}

func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.Actions = append(b.Actions, func(person *Person) {
		person.Position = position
	})

	return b
}

func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.Actions {
		a(&p)
	}

	return &p
}
