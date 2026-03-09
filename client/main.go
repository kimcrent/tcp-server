package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "82.40.38.98:1337")
	if err != nil {
		log.Fatalf("Conntion with server error %v", err)
	}
	defer conn.Close()

	fmt.Println("Server connected")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter ur massage")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Input massage error %v", err)
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Fatalf("Send massage error %v", err)
		}

		bufferReader := make([]byte, 1024)

		n, err := conn.Read(bufferReader)
		if err != nil {
			log.Fatalf("Read massge error %v", err)
		}
		fmt.Println("Server answer:", string(bufferReader[:n]))
	}
}
