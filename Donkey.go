package main

import (
	"time"
)

type Donkey struct {
	PackAnimals
}

func (d Donkey) New(name string, birthDate time.Time, commands []string) AnimalInterface {
	d.BirthDate = birthDate
	d.Name = name
	d.Commands = commands
	return &d
}

func (d Donkey) GetClass() string {
	return "Donkey"
}
