package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	//"strings"
)

type Human struct {
	name    string
	surname string
	gender  string
	age     int
}

func main() {
	var count int
	var name, surname, gender string
	var age int

	fmt.Println("Введите количество человек: ")
	fmt.Scanf("%d\n", &count)
	var massiv = make([]Human, count)
	var massiv1 []Human

	// открытие файла для чтения
	file, _ := os.ReadFile("json1.json")

	err := json.Unmarshal(file, &massiv1)
	if err != nil {
		log.Fatal(err)
	}
	for i := range massiv1 {
		fmt.Println(massiv1[i])
	}

	for i := 0; i < count; i++ {
		fmt.Print("Введите имя человека: ")
		fmt.Scanf("%s\n", &name)
		fmt.Print()

		fmt.Print("Введите фамилию человека: ")
		fmt.Scanf("%s\n", &surname)
		fmt.Print()

		fmt.Print("Введите пол человека: ")
		fmt.Scanf("%s\n", &gender)
		fmt.Print()

		fmt.Print("Введите возраст человека: ")
		fmt.Scanf("%d\n", &age)
		fmt.Println()

		human := Human{name, surname, gender, age}
		massiv[i] = human
	}

	for i := 0; i < count; i++ {
		fmt.Println(massiv[i])
	}
}
