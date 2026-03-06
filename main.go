package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatalf("Error with listen port 1337 %v", err)
	}
	defer ln.Close()

	log.Println("Port 1337 listening")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Error, connection with client %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		bufferReader, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Read error %v", err)
			return
		}

		data := buffer[:bufferReader]

		log.Println("Buffer is:", string(data))

		_, err = conn.Write([]byte("Package accept"))
		if err != nil {
			log.Printf("Answer error %v", err)
			return
		}
	}

}
