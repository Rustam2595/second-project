package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 3 урок
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero") //fmt.Errorf("divide by zero)
	}
	return a / b, nil
}
func strToInt(val string) (int, error) {
	res, err := strconv.Atoi(val)
	if err != nil {
		return -1, fmt.Errorf("convert to int err:%w", err)
	}
	return res, nil
}

func causePanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic", err)
		}
	}()
	//panic("This is panic")
}

// Методы и структуры /  Интерфейсы в Go
type Animal interface {
	Run()
	Sleep()
}

type Cat struct {
	Name  string
	Color string
}

func (c Cat) Run() {
	fmt.Printf("%s running\n", c.Name)
}

func (c Cat) Sleep() {
	fmt.Printf("%s %s is sleep\n", c.Color, c.Name)
}

func tester(anim Animal) {
	anim.Run()
	anim.Sleep()
}
func gg(val any) {
	fmt.Println(val)
}

func main() {
	a, err := divide(0, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(a)

	_, err1 := strToInt("12")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	causePanic()

	cat := Cat{
		Name:  "Lusa",
		Color: "black",
	}
	cat.Run()
	tester(cat)
	fmt.Println()
	gg("ff")
	gg(false)
}
