package main

import (
	"time"
)

type Hamster struct {
	Pets
}

func (h Hamster) New(name string, birthDate time.Time, commands []string) AnimalInterface {
	h.BirthDate = birthDate
	h.Name = name
	h.Commands = commands
	return &h
}

func (h Hamster) GetClass() string {
	return "Hamster"
}
