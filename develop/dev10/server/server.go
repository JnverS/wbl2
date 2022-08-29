package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// начинаем слушать порт
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		// принимаем входящие подключения
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}
}

// handleConnection считываем из сокета и записываем ответ
func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		log.Println("reseived:", text)
		n, err := conn.Write([]byte("answer: " + text))
		if err != nil || n == 0 {
			log.Println("Write error:", err)
			return
		}
		if text == "exit" {
			break
		}
	}
	err := scanner.Err()
	if err != nil {
		log.Println(err)
	}
	log.Println("Closing connection")
}
