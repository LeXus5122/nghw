package main

import (
	"time"
)

type Dog struct {
	Pets
}

func (d Dog) New(name string, birthDate time.Time, commands []string) AnimalInterface {
	d.BirthDate = birthDate
	d.Name = name
	d.Commands = commands
	return &d
}

func (d Dog) GetClass() string {
	return "Dog"
}
