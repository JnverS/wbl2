package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
	Утилита cut

	Реализовать	утилиту	аналог	консольной	команды	cut	(man cut).
	Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

	Реализовать поддержку утилитой следующих ключей:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем
*/

// flags структура для хранения флагов
type flags struct {
	f int
	d string
	s bool
}

func main() {

	// парсим флаги
	f := flag.Int("f", 0, "выбрать поля (колонки)")
	d := flag.String("d", "	", "использовать другой разделитель")
	s := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()
	flags := flags{*f, *d, *s}
	if flags.f <= 0 {
		log.Fatalln("you must specify a list of fields")
	}

	// запускаем считывание с stdin
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ok := scanner.Scan()
		if !ok || scanner.Err() != nil {
			log.Fatalln("Scan error")
		}
		if scanner.Text() != "" {
			flags.printString(scanner.Text())
		}
	}
}

// printString печатаем нужную колонку строки
func (f *flags) printString(str string) {
	col := strings.Split(str, f.d)
	if f.s {
		if f.f > 0 && len(col) > 1{
			if len(col) >= f.f {
				fmt.Println(col[f.f-1])
			} else {
				fmt.Println("")
			}
		}
	} else {
		if f.f > 0 && len(col) >= f.f {
			fmt.Println(col[f.f-1])
		} else {
			fmt.Println(str)
		}
	}
}
