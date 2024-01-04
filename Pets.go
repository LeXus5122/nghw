package main

type Pets struct {
	Animals
}

func (p Pets) GetType() string {
	return "Домашнее животное"
}
