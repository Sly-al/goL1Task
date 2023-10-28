package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func Sleep(i int) {
	<-time.After(time.Duration(i) * time.Second)
}

func main() {
	var i int
	fmt.Print("Enter number of seconds: ")
	fmt.Scan(&i)
	Sleep(i)
	fmt.Println("End")
}
