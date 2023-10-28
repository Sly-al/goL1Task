package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Name    string
	Surname string
}

func (h *Human) greeting() {
	fmt.Printf("Nicve to meet you %s %s \n", h.Name, h.Surname)
}

type Action struct {
	// безымянное встраивание структуры Human
	Human
	IsVacation bool
}

func (a *Action) status() {
	// вызываем метод greeting() структуры Human
	a.greeting()
	switch a.IsVacation {
	case true:
		fmt.Println("Congrats! You are having vacation!")
	case false:
		fmt.Println("Sorry, but you have to work")
	}
}

func main() {
	a := Action{
		Human: Human{
			Name:    "Json",
			Surname: "Statham",
		},
		IsVacation: false,
	}
	a.status()
	// метод human доступен
	a.greeting()
	// можно указывать явно
	a.Human.greeting()
}
