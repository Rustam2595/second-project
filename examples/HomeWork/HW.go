package Hw

import (
	"fmt"
)

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
		fmt.Printf("'%s' contains %d words.\n", sentence, CountWords(sentence))
	}
}
