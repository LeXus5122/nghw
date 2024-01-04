package main

import (
	"time"
)

type Cat struct {
	Pets
}

func (c Cat) New(name string, birthDate time.Time, commands []string) AnimalInterface {
	c.BirthDate = birthDate
	c.Name = name
	c.Commands = commands
	return &c
}

func (c Cat) GetClass() string {
	return "Cat"
}
