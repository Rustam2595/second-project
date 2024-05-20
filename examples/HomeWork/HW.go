package main

import (
	"fmt"
	"strings"
)

// countWords подсчитывает количество слов в строке
func countWords(s string) int {
	// Разделяем строку на слова
	words := strings.Fields(s)
	// Возвращаем количество слов
	return len(words)
}

func main() {
	// Примеры строк для тестирования
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
}
