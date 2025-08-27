package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 9000, "port to listen on")
	flag.Parse()
	runServer(*port)
}

func runServer(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		log.Println("Accepted connection:", conn.RemoteAddr())
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	io.Copy(conn, conn)
	conn.Close()
	log.Println("Closed connection:", conn.RemoteAddr())
}
