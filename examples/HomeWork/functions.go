package Hw

import "strings"

func CountWords(s string) int {
	// Разделяем строку на слова
	words := strings.Fields(s)
	// Возвращаем количество слов
	return len(words)
}
