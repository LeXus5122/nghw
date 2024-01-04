package main

import (
	"fmt"
	"time"
)

type Animals struct {
	Name      string
	BirthDate time.Time
	Commands  []string
}

func (a Animals) ShowCommands() {
	s := "Команды:\n"
	for i, command := range a.Commands {
		s += fmt.Sprintf("%d. %s\n", i+1, command)
	}
	fmt.Print(s)
}

func (a *Animals) AddCommand(command string) {
	a.Commands = append(a.Commands, command)

}

func (a Animals) New(name string, birthDate time.Time, commands []string) AnimalInterface {
	a.BirthDate = birthDate
	a.Name = name
	a.Commands = commands
	return &a
}

func (a Animals) GetName() string {
	return a.Name
}

func (a Animals) GetClass() string {
	return "Не определен"
}

func (a Animals) GetType() string {
	return "Не определен"
}
