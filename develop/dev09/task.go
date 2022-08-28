package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

/*
	Утилита wget

	Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		log.Fatalln("wget: missing URL")
	}
	err := wget(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
}

// wget функция для скачивания сайтов
func wget(url string) error {
	// делаем get запрос по указанному url
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// создаем файл
	fileName := "index.html"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// копируем полученные данные в файл
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("%s - '%s' saved [%d]\n", time.Now().Format("2006-08-02 15:04:05"), fileName, size)
	return nil
}
