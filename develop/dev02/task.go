package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Unpacking(s string) (string, error){
	var str strings.Builder
	r := []rune(s)
	backslash := false
	for i, v := range r {
		if unicode.IsDigit(v) && i == 0 {
			return "", errors.New("Invalid string")
		}
		if unicode.IsDigit(v) && !unicode.IsDigit(r[i - 1]) && !backslash {
			for j := 0; j < int(v - '0') - 1; j++ {
				str.WriteRune(r[i - 1])
			}
		}
		if unicode.IsLetter(v) {
			str.WriteRune(v)
		}
		if v == '\\' {
			backslash = true
		}
	}
	return str.String(), nil
}


func main() {
	s := "qwe\\4"
	news, err := Unpacking(s)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(news)
}
