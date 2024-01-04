package main

import (
	"time"
)

type Horse struct {
	PackAnimals
}

func (h Horse) New(name string, birthDate time.Time, commands []string) AnimalInterface {
	h.BirthDate = birthDate
	h.Name = name
	h.Commands = commands
	return &h
}

func (h Horse) GetClass() string {
	return "Horse"
}
