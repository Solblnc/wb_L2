package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//Задача на распаковку
//Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
//"a4bc2d5e" => "aaaabccddddde"
//"abcd" => "abcd"
//"45" => "" (некорректная строка)
//"" => ""
//
//Дополнительно
//Реализовать поддержку escape-последовательностей.
//Например:
//qwe\4\5 => qwe45 (*)
//qwe\45 => qwe44444 (*)
//qwe\\5 => qwe\\\\\ (*)
//
//
//В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

func unpackString(s string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			return ""
		}

		result.WriteRune(runes[i])
		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			count, _ := strconv.Atoi(string(runes[i+1]))
			result.WriteString(strings.Repeat(string(runes[i]), count-1))
			i++
		}
	}

	return result.String()
}

func main() {
	s := "a4b3hjk"
	fmt.Println(unpackString(s))
}
