package main

/*
	Задача на распаковку
	Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны,
	например:
	"a4bc2d5e" => "aaaabccddddde"
	"abcd" => "abcd"
	"45" => "" (некорректная строка)
	"" => ""

	Дополнительно
	Реализовать поддержку escape-последовательностей.
	Например:
	qwe\4\5 => qwe45 (*)
	qwe\45 => qwe44444 (*)
	qwe\\5 => qwe\\\\\ (*)

	В случае если была передана некорректная строка, функция должна возвращать ошибку.
	Написать unit-тесты.
*/

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// Unpacking using for unpacking string
func Unpacking(s string) (string, error){
	var str strings.Builder
	r := []rune(s)
	backslash := false
	for i, v := range r {
		// проверяем случай, когда в начале строки цифра
		if unicode.IsDigit(v) && i == 0 {
			return "", errors.New("Invalid string")
		}
		// если цифра - добалвяем предыдущий символ n раз
		if unicode.IsDigit(v) && !backslash {
			for n := 0; n < int(v - '0') - 1; n++ {
				str.WriteRune(r[i - 1])
			}
		} else {
			// если встретили escape-последовательность
			if v == '\\' && r[i - 1] != '\\'{
				backslash = true
			} else {
				backslash = false
			}
			// добавляем символ в строку
			if !backslash {
				str.WriteRune(v)
			}
		}
	}
	return str.String(), nil
}


func main() {
	s := "a4bc2d5e"
	news, err := Unpacking(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(news)
}

