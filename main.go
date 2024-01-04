package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	animals := []AnimalInterface{}

	for {
		showMenu()
		var choice int
		_, err := fmt.Scanf("%d\n", &choice)
		if err != nil {
			continue
		}

		switch choice {
		case 1:
			addAnimal(&animals)
		case 2:
			showType(animals)
		case 3:
			showCommands(animals)
		case 4:
			addCommand(animals)
		case 5:
			return
		default:
			continue

		}
	}
}

func addCommand(animals []AnimalInterface) {
	animal := choiceAnimal(animals)

	fmt.Print("Введите новую команду: ")
	reader := bufio.NewReader(os.Stdin)

	command, err := reader.ReadString('\n')
	if err != nil {
		addCommand(animals)
		return
	}
	command, _ = strings.CutSuffix(command, "\n")
	animal.AddCommand(command)
	fmt.Print("Команда добавлена успешно. Для продолжения нажмите Enter")
	fmt.Scanf("\n")

}

func addAnimal(animals *[]AnimalInterface) {
	var choice int
	for {
		fmt.Println("Типы животных:\n" +
			"1) Cat\n" +
			"2) Dog\n" +
			"3) Hamster\n" +
			"4) Horse\n" +
			"5) Donkey\n" +
			"6) Camel\n" +
			"7) Назад\n" +
			"Введите номер: ")
		_, err := fmt.Scanf("%d\n", &choice)
		if err != nil {
			continue
		}
		if choice == 7 {
			return
		}
		if choice > 6 || choice < 1 {
			continue
		}
		break
	}

	fmt.Println("Введите имя: ")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')
	if err != nil {
		addAnimal(animals)
		return
	}

	var birthDay time.Time

	for {
		fmt.Print("Введите дату рождения (dd-mm-yyyy): ")
		var strBirthDay string
		_, err = fmt.Scanf("%s\n", &strBirthDay)
		if err != nil {
			continue
		}

		birthDay, err = time.Parse("02-01-2006", strBirthDay)
		if err != nil {
			fmt.Println("Попробуйте еще раз")
			continue
		}
		break
	}

	fmt.Print("Введите команды через запятую: ")

	strCommands, err := reader.ReadString('\n')

	commands := strings.Split(strCommands, ",")

	for i := 0; i < len(commands); i++ {
		commands[i], _ = strings.CutPrefix(commands[i], " ")
		commands[i], _ = strings.CutSuffix(commands[i], "\n")
	}

	switch choice {
	case 1:
		*animals = append(*animals, Cat{}.New(name, birthDay, commands))
	case 2:
		*animals = append(*animals, Dog{}.New(name, birthDay, commands))
	case 3:
		*animals = append(*animals, Hamster{}.New(name, birthDay, commands))
	case 4:
		*animals = append(*animals, Horse{}.New(name, birthDay, commands))
	case 5:
		*animals = append(*animals, Donkey{}.New(name, birthDay, commands))
	case 6:
		*animals = append(*animals, Camel{}.New(name, birthDay, commands))

	}

	fmt.Print("Новое животное добавлено успешно. Для продолжения нажмите Enter")
	fmt.Scanf("\n")
}

func showMenu() {
	fmt.Print("Добро пожаловать!\n" +
		"1) Создать новое животное\n" +
		"2) Определить тип животного\n" +
		"3) Посмотреть список команд животного\n" +
		"4) Добавить новую команду животному\n" +
		"5) Выход\n" +
		"Введите пункт меню: ")
}

func choiceAnimal(animals []AnimalInterface) AnimalInterface {
	s := "Список животных:\n"
	for i, animal := range animals {
		s += fmt.Sprintf("%d) %s %s", i+1, animal.GetClass(), animal.GetName())
	}
	s += "Введите номер животного: "
	fmt.Print(s)

	var choice int
	_, err := fmt.Scanf("%d\n", &choice)
	if err != nil {
		return choiceAnimal(animals)
	}

	return animals[choice-1]
}

func showType(animals []AnimalInterface) {
	animal := choiceAnimal(animals)
	fmt.Println(animal.GetType())

	fmt.Print("Для продолжения нажмите Enter")
	_, _ = fmt.Scanf("%s\n")
}

func showCommands(animals []AnimalInterface) {
	animal := choiceAnimal(animals)
	animal.ShowCommands()

	fmt.Print("Для продолжения нажмите Enter")
	_, _ = fmt.Scanf("%s\n")
}
