package main

import (
	"fmt"
	"math"
	"strings"
)

// countWords подсчитывает количество слов в строке
func countWords(s string) int {
	// Разделяем строку на слова
	words := strings.Fields(s)
	// Возвращаем количество слов
	return len(words)
}

// isPrime проверяет, является ли число простым
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// findMinMax находит наименьший и наибольший элемент в массиве
func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0 // Обрабатываем случай пустого массива
	}
	min := math.MaxInt64
	max := math.MinInt64
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

////////////////////////////////////////////////////////////////

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}
func (b *BankAccount) Withdraw(amount float64) error {
	if amount <= b.Balance {
		b.Balance -= amount
		return nil
	}
	return fmt.Errorf("недостаточно денег для снятия")
}

// //////////////////////////////////////////////////////////////////
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

type Dog struct {
	Name string
	Age  int
}

func (c Dog) Run() {
	fmt.Printf("%s running\n", c.Name)
}

func (c Dog) Sleep() {
	fmt.Printf("%s is sleep\n", c.Name)
}

func tester(anim Animal) {
	anim.Run()
	anim.Sleep()
}

func gg(val any) {
	fmt.Println(val)
}

/////////////////////////////////////////////////////////

func main() {
	// Примеры массивов для тестирования
	arrays := [][]int{
		{3, 5, 7, 2, 8, -1, 4, 10, 12},
		{1, 1, 1, 1, 1, 1},
		{-5, -1, -15, -3},
		{7},
		{},
	}

	for _, arr := range arrays {
		min, max := findMinMax(arr)
		fmt.Printf("Array: %v, Min: %d, Max: %d\n", arr, min, max)
	}

	//Тестируем функцию на нескольких значениях
	numbers := []int{1, 2, 3, 4, 5, 16, 17, 18, 19, 20, 23, 24}
	for _, num := range numbers {
		if isPrime(num) {
			fmt.Printf("%d is a prime number.\n", num)
		} else {
			fmt.Printf("%d is not a prime number.\n", num)
		}
	}
	//////////////////////////////////////////////////////////////////////
	sentences := []string{
		"Hello, world!",
		"   This is a test string.  ",
		"Go is a statically typed, compiled programming language.",
		"      ",
		"OneWord",
		"Two words",
	}
	for _, sentence := range sentences {
		fmt.Printf("'%s' contains %d words.\n", sentence, countWords(sentence))
	}

	////////////////////////////////////////////////
	b := BankAccount{Owner: "Alex", Balance: 0}
	if err := b.Withdraw(5); err != nil {
		fmt.Println(err)
	}
	b.Deposit(10)
	if err := b.Withdraw(5); err != nil {
		fmt.Println(err)
	}
	///////////////////////////////////////////////////////
	cat := Cat{
		Name:  "Lusa",
		Color: "black",
	}
	dog := Dog{
		Name: "Sharik",
		Age:  4,
	}
	cat.Run()
	dog.Sleep()
	fmt.Println()
	tester(dog)
	tester(cat)
	fmt.Println()
	gg(1)
	gg("ff")
	gg(false)
	//////////////////////////////////////////////

}
