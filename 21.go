package main

import (
	"fmt"
)

// реализация паттерна адаптер

// Допустим, нам пришлось рабоатть с очень старой библиотекой, которая при этом своеобразно написана

// структура ответа
type Response struct {
	status int
	prices map[string]int
}

func (res *Response) getResStatus() {
	fmt.Println(res.status)
}

// структура запроса
type Request struct {
	status string
	orders []string
}

func (req *Request) getReqStatus() {
	fmt.Println(req.status)
}

// на этом код старой библиотеки закончился

/* хотим объединить методы для отображение статуса
для этого воспользуемся паттерном адаптер */

type getMessage interface {
	GetStatus()
}

type ReqAdapter struct {
	*Request
}

type ResAdapter struct {
	*Response
}

func (adapter *ReqAdapter) GetStatus() {
	adapter.getReqStatus()
}

func (adapter *ResAdapter) GetStatus() {
	adapter.getResStatus()
}

func main() {
	request := &Request{status: "ok"}
	response := &Response{status: 200}

	reqAdapter := getMessage(&ReqAdapter{request})
	resAdapter := getMessage(&ResAdapter{response})

	reqAdapter.GetStatus()
	resAdapter.GetStatus()
}
