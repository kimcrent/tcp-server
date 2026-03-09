package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalf("Port listening error %v", err)
	}

	defer ln.Close()

	log.Println("Port 1337 listening")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Accept error %v", err)
			return
		}

		go handleConnaction(conn)
	}
}

func handleConnaction(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		bufferReader, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Read error %v", err)
			return
		}

		data := buffer[:bufferReader]

		log.Println("Massage is: ", string(data))

		_, err = conn.Write(data)
		if err != nil {
			log.Printf("Write error %v", err)
			return
		}
	}
}
