package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает
две числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	var (
		num1, num2 string
		op         rune
	)
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1) // считываем как строку
	fmt.Print("Second number: ")
	fmt.Scan(&num2)
	fmt.Print("Input operator: ")
	reader := bufio.NewReader(os.Stdin)
	op, _, err := reader.ReadRune()
	if err != nil {
		log.Fatalf(err.Error())
	}
	// создаем переменные big.Float
	a := new(big.Float)
	b := new(big.Float)
	c := big.NewFloat(0) // для хранение результата
	a.SetString(num1)
	b.SetString(num2)
	// в зависимости от символа операции выполняем нужную операцию
	switch op {
	case '+':
		fmt.Printf("%f / %f = %f\n", a, b, c.Add(a, b))
	case '-':
		fmt.Printf("%f / %f = %f\n", a, b, c.Sub(a, b))
	case '/':
		fmt.Printf("%f / %f = %f\n", a, b, c.Quo(a, b))
	case '*':
		fmt.Printf("%f / %f = %f\n", a, b, c.Mul(a, b))
	default:
		fmt.Print("Unknown operator\n")
	}
}
