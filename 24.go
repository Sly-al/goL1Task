package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

// структурка для точки x y - неэкспортируемые поля
type Point struct {
	x, y int
}

// конструктор для создания новой точки
func newPoint(xVal, yVal int) Point {
	return Point{x: xVal, y: yVal}
}

// нахождение расстояние между двумя точками
func distPoints(p1, p2 Point) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func main() {
	var x, y int
	fmt.Print("Введите координаты x и y точки ")
	fmt.Scan(&x, &y)
	point1 := newPoint(x, y)
	fmt.Print("Введите координаты x и y точки ")
	fmt.Scan(&x, &y)
	point2 := newPoint(x, y)
	fmt.Println("Расстояние ", distPoints(point1, point2))
}
