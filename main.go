package main

import (
	"flag"
	"fmt"

	LRPSignals "github.com/PiterWeb/LibreRemotePlaySignals/v1"
)

func main() {

	// -p flag for port number
	port := flag.Uint("port", 80, "Port number to listen on")

	flag.Parse()

	ips_chan := make(chan []string, 1)
	err_chan := make(chan error, 1)

	// Create a new instance of the server
	go func() {
		err := LRPSignals.InitServer(LRPSignals.ServerOptions{Port: uint16(*port)}, ips_chan)

		err_chan <- err
	}()

	// Wait for the server to stop
	err := <-err_chan

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Server stopped successfully")
	}
}
