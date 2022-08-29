package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
	Утилита telnet

	Реализовать простейший telnet-клиент.

	Примеры вызовов:
	go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

	Требования:
	Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
	После подключения STDIN программы должен записываться в сокет,
	а данные полученные и сокета должны выводиться в STDOUT
	Опционально в программу можно передать таймаут на подключение
	 к серверу (через аргумент --timeout, по умолчанию 10s)
	При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
	 Если сокет закрывается со стороны сервера, программа должна также завершаться.
	При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

func main() {
	timeout := flag.Int("timeout", 60, "timeout connection")
	flag.Parse()
	if len(flag.Args()) < 2 {
		log.Fatalln("Enter port and host")
	}
	host := flag.Arg(0)
	port := flag.Arg(1)
	t := time.Duration(*timeout) * time.Second

	//пытаемся подключиться к серверу
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), t)
	// если не подключились к серверу ждем timeout
	if err != nil {
		time.Sleep(t)
		log.Fatalln(err)
	}
	defer conn.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGQUIT)

	// читаем из stdin и пишем в сокет
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				conn.Close()
				log.Fatalln("Connection is broken")
			}
			_, err = conn.Write([]byte(text))
			if err != nil {
				conn.Close()
				log.Fatalln(err)
			}
		}
	}()

	// читаем из сокета и печатаем
	go func() {
		reader := bufio.NewReader(conn)
		for {
			text := make([]byte, 1024)
			n, err := reader.Read(text)
			if err != nil {
				conn.Close()
				log.Fatalln("Connection is broken")
			}
			fmt.Println(string(text[:n]))
		}
	}()

	// ждем сигнала Ctrl+D
	select {
	case <-sigCh:
		conn.Close()
	}
}
