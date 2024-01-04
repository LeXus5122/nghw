package main

import (
	"time"
)

type Camel struct {
	PackAnimals
}

func (c Camel) New(name string, birthDate time.Time, commands []string) AnimalInterface {
	c.BirthDate = birthDate
	c.Name = name
	c.Commands = commands
	return &c
}

func (c Camel) GetClass() string {
	return "Camel"
}
