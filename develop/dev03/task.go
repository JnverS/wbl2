package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
	Утилита sort
	Отсортировать строки в файле по аналогии с консольной утилитой sort
	(man sort — смотрим описание и основные параметры):
	на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

	Реализовать поддержку утилитой следующих ключей:

	-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок,
		по умолчанию разделитель — пробел)
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки

	Дополнительно

	Реализовать поддержку утилитой следующих ключей:

	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учетом суффиксов
*/

type flags struct {
	k       int
	n, r, u bool
}

func main() {
	// парсим флаги
	k := flag.Int("k", -1, "sort by column")
	n := flag.Bool("n", false, "sort by number")
	r := flag.Bool("r", false, "reverse sort")
	u := flag.Bool("u", false, "only unique strings")
	flag.Parse()
	flags := flags{*k, *n, *r, *u}

	// читаем из файла
	file, err := ReadFile("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	//соортируем
	outFile := SortStrings(file, flags)

	//записываем в файл
	if err = WriteFile(outFile); err != nil {
		log.Fatalln(err)
	}
}

// ReadFile считываем файл в слайс строк
func ReadFile(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	result := strings.Split(string(buffer), "\n")
	return result, nil
}

// WriteFile записывает отсортированный результат в файл
func WriteFile(str []string) error {
	file, err := os.Create("outFile.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	for _, v := range str {
		_, err = file.WriteString(v + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

// SortStrings здесь происходит сортировка
func SortStrings(str []string, f flags) []string {
	newStr := str
	if f.u {
		newStr = makeUnique(str)
	}
	if f.n {
		sort.Slice(newStr, func(i, j int) bool {
			first, _ := strconv.Atoi(newStr[i])
			second, _ := strconv.Atoi(newStr[j])
			return first < second
		})
		if f.r {
			newStr = Reverse(newStr)
		}
		return newStr
	}
	if f.k > -1 {
		sort.Slice(newStr, func(i int, j int) bool {
			left := strings.Split(str[i], " ")
			right := strings.Split(str[j], " ")
			if len(left) <= f.k || len(right) <= f.k {
				log.Fatalln("k > count column")
				return false
			}
			return left[f.k] < right[f.k]
		})
		if f.r {
			newStr = Reverse(newStr)
		}
		return newStr
	}
	if f.r {
		newStr = Reverse(newStr)
		return newStr
	}
	sort.Slice(newStr, func(i, j int) bool {
		return newStr[i] < newStr[j]
	})
	return newStr
}

// makeUnique убирает дубли
func makeUnique(str []string) []string {
	buffer := make(map[string]struct{})
	result := make([]string, 0, len(str))
	for _, v := range str {
		if _, ok := buffer[v]; !ok {
			buffer[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Reverse сортирует в обратном порядке
func Reverse(str []string) []string {

	sort.Slice(str, func(i, j int) bool {
		return str[i] > str[j]
	})
	return str
}
