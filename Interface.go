package main

import "time"

type AnimalInterface interface {
	New(string, time.Time, []string) AnimalInterface
	GetName() string
	GetClass() string
	GetType() string
	ShowCommands()
	AddCommand(string)
}
