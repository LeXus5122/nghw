package main

type PackAnimals struct {
	Animals
}

func (p PackAnimals) GetType() string {
	return "Вьючное животное"
}
