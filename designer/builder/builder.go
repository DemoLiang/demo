package main

import "fmt"

type Character struct {
	Name string
	Arms string
}

func (this *Character) SetName(name string) {
	this.Name = name
}

func (this *Character) SetArms(arms string) {
	this.Arms = arms
}

func (this Character) GetName() string {
	return this.Name
}
func (this Character) GetArms() string {
	return this.Arms
}

type Builder interface {
	SetName(name string) Builder
	SetArms(arms string) Builder
	Build() *Character
}

type CharacterBuilder struct {
	character *Character
}

func (this *CharacterBuilder) SetName(name string) Builder {
	if this.character == nil {
		this.character = &Character{}
	}
	this.character.SetName(name)
	return this
}

func (this *CharacterBuilder) SetArms(arms string) Builder {
	if this.character == nil {
		this.character = &Character{}
	}
	this.character.SetArms(arms)
	return this
}

func (this *CharacterBuilder) Build() *Character {
	return this.character
}

type Director struct {
	builder Builder
}

func (this Director) Create(name, arms string) *Character {
	return this.builder.SetName(name).SetArms(arms).Build()
}

func main() {
	var builder Builder = &CharacterBuilder{}
	var director *Director = &Director{builder: builder}
	var character *Character = director.Create("loader", "AK47")
	fmt.Printf(character.GetName() + "," + character.GetArms())
}
