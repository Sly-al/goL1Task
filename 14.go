package main

import (
	"fmt"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	varValues := []interface{}{
		42,
		"hello",
		true,
		make(chan int),
	}

	for _, value := range varValues {
		determineType(value)
	}
}

func determineType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("'%v' имеет тип int\n", v)
	case string:
		fmt.Printf("'%v' имеет тип string\n", v)
	case bool:
		fmt.Printf("'%v' имеет тип bool\n", v)
	case chan int:
		fmt.Println("Значение имеет тип chan int")
	default:
		fmt.Printf("Тип '%v' не поддерживается\n", v)
	}
}
