package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp4", ":1337")
	if err != nil {
		log.Fatalf("Listener error: %v", err)
	}
	defer ln.Close()
	log.Println("server started on :8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("read error:", err)
			return
		}
		data := buffer[:n]

		fmt.Println("received: ", string(data))

		conn.Write([]byte("message received\n"))
	}
}
