package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func RunServer(port int) {
	log.Println("Starting server on port", port)
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

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
func handleConnection(conn net.Conn) {
	input := bufio.NewScanner(conn)

	for input.Scan() {
		echo(conn, input.Text(), 1*time.Second)
	}
	conn.Close()
	log.Println("Closed connection:", conn.RemoteAddr())
}
