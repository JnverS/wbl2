package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
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
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fileName := "index.html"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	str := parseURL(url, string(body))

	ioutil.WriteFile(fileName, []byte(str), 0777)
	return nil
}

func parseURL(url string, body string) string {
	c := colly.NewCollector()
	sBody := strings.Split(body, "\"")

	// Find all img
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("src"))
		for i, v := range sBody {
			if strings.EqualFold(v, e.Attr("src")) {
				sBody[i] = "./site/" + e.Attr("src")[strings.LastIndex(e.Attr("src"), "/")+1:]
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fileName := "./site/" + r.URL.String()[strings.LastIndex(r.URL.String(), "/")+1:]
		output, _ := os.Create(fileName)
		defer output.Close()

		resp, err := http.Get(r.URL.String())
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		io.Copy(output, resp.Body)
	})
	c.Visit(url)
	return strings.Join(sBody, "\"")
}
