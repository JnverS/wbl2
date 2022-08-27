package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

/*
	Утилита grep

	Реализовать утилиту фильтрации по аналогии с консольной утилитой
	(man grep — смотрим описание и основные параметры).
	Реализовать поддержку утилитой следующих ключей:
++	-A - "after" печатать +N строк после совпадения
++	-B - "before" печатать +N строк до совпадения
++	-C - "context" (A+B) печатать ±N строк вокруг совпадения
++	-c - "count" (количество строк)
++	-i - "ignore-case" (игнорировать регистр)
++	-v - "invert" (вместо совпадения, исключать)
++	-F - "fixed", точное совпадение со строкой, не паттерн
++	-n - "line num", напечатать номер строки
*/

// Flags структура для хранения флагов
type Flags struct {
	A, B, C       int
	c, i, v, F, n bool
	str, fileName string
}

// GrepFile функция grep
func GrepFile() {
	// парсим входящие данные
	flags, err := ParseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	// читаем файл
	file, err := ReadFile(flags.fileName)
	if err != nil {
		log.Fatalln(err)
	}
	if flags.C > 0 {
		flags.A = flags.C
		flags.B = flags.C
	}
	// ищем совпадение строк
	res, key := flags.FindAllStrings(file)

	// если флаг с просто выводим количество найденных строк
	if flags.c {
		fmt.Println(len(res))
		return
	}
	// печатаем получившиеся строки в консоль
	flags.PrintResult(res, key)
}

// PrintResult печатаем результат
func (f *Flags) PrintResult(res map[int]string, key []int) {
	for _, v := range key {
		if f.n {
			fmt.Print(v, ":", res[v], "\n")
		} else {
			fmt.Println(res[v])
		}
	}
}

//FindAllStrings ищем строчки
func (f *Flags) FindAllStrings(file []string) (map[int]string, []int) {
	result := make(map[int]string)
	key := make([]int, 0)
	write := false
	for i, v := range file {
		w := v
		if f.i {
			w = strings.ToLower(w)
			f.str = strings.ToLower(f.str)
		}
		if f.F {
			if strings.Contains(w, f.str) {
				result[i+1] = v
				key = append(key, i+1)
				write = true
			}
		} else {
			check, err := regexp.MatchString(f.str, w)
			if err != nil {
				log.Fatal(err)
			}
			if f.v {
				if !check {
					result[i+1] = v
					key = append(key, i+1)
					write = true
				}
			} else {
				if check {
					result[i+1] = v
					key = append(key, i+1)
					write = true
				}
			}
			// если нужно добавить строки до и после
			if (f.B > 0 || f.A > 0) && write && !f.c {
				j := i + 1
				for n := f.A; n > 0; n-- {
					if j < len(file) {
						result[j+1] = file[j]
						key = append(key, j+1)
						j++
					}
				}
				j = i - 1
				for n := f.B; n > 0; n-- {
					result[j+1] = file[j]
					key = append(key, j+1)
					j--
				}
				write = false
			}
		}
	}
	sort.Ints(key)
	return result, key
}

// ReadFile считываем файл
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

// ParseFlags парсим флаги, которые нам пришли из cmd
func ParseFlags() (*Flags, error) {
	A := flag.Int("A", -1, "печатать +N строк после совпадения")
	B := flag.Int("B", -1, "печатать +N строк до совпадения")
	C := flag.Int("C", -1, "(A+B) печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "количество строк")
	i := flag.Bool("i", false, "игнорировать регистр")
	v := flag.Bool("v", false, "вместо совпадения, исключать")
	F := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()
	if len(flag.Args()) != 2 {
		return nil, errors.New("Enter string for found and filename")
	}
	str := flag.Arg(0)
	fileName := flag.Arg(1)
	flags := Flags{*A, *B, *C, *c, *i, *v, *F, *n, str, fileName}
	return &flags, nil
}

func main() {
	GrepFile()
}
