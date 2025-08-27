package main

import (
	"flag"
	"go-echo-server/client"
	"go-echo-server/server"
)

func main() {
	port := flag.Int("p", 9000, "port to listen on")
	mode := flag.String("m", "server", "mode to run in")
	flag.Parse()

	switch *mode {
	case "server":
		server.RunServer(*port)
	case "client":
		client.RunClient(*port)
	}
}
