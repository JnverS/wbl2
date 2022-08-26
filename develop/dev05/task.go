package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
type flags struct {
	A, B, C int
	c, i, v, F, n bool
	str, fileName string
}

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
	if flags.i {

	}
	// ищем совпадение строк
	res, key := FindAllStrings(flags.str, file, *flags)
	if flags.c {
		fmt.Println(len(res))
		return
	}
	// печатаем получившиеся строки в консоль
	PrintResult(res, key, *flags)
}

func PrintResult(res map[int]string, key []int, flags flags) {
	for _, v := range key {
		if flags.n {
			fmt.Print(v, ":", res[v], "\n")
		} else {
			fmt.Println(res[v])
		}
	}
}

func FindAllStrings(str string, file []string, flags flags) (map[int]string, []int) {
	result := make(map[int]string)
	key := make([]int, 0)
	NA, NB, NC := 0, 0, 0
	for i, v := range file {
		w := v
		if flags.F {
			if v == str {
				result[i + 1] = v
				key = append(key, i + 1)
			}
			break
		}
		if flags.i {
			str = strings.ToLower(str)
			w = strings.ToLower(w)
		}
		if flags.v {
			if !strings.Contains(w, str) {
				result[i + 1] = v
				key = append(key, i + 1)
			}
		} else {
			if strings.Contains(w, str){
				result[i + 1] = v
				key = append(key, i + 1)
				if flags.A != -1{
					NA = flags.A
				}
				if flags.B != -1{
					j := i
					for NB = flags.B; NB > 0;NB--{
						result[j] = file[j]
						key = append(key, j)
						j--
					}
				}
				if flags.C != -1{
					j := i
					for NC = flags.C; NC > 0;NC--{
						result[j] = file[j - 1]
						key = append(key, j)
						j--
					}
					j = i + 1
					for NC = flags.C; NC > 0;NC--{
						result[j + 1] = file[j]
						key = append(key, j + 1)
						j++
					}
				}
			} else {
				if NA != 0 {
					result[i + 1] = v
					key = append(key, i + 1)
					NA--
				}
			}
		}
	}
	sort.Ints(key)
	return result, key
}

func ReadFile(name string) ([]string, error){
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

func ParseFlags() (*flags, error) {
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
	flags := flags{*A, *B, *C, *c, *i, *v, *F, *n, str, fileName}
	return &flags, nil
}

func main() {
	GrepFile()
}