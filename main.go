package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/LeftTwixWand/Slowly/server"
)

const defaultPort = "8080"
const defaultHost = "localhost"

func main() {
	port, ok := os.LookupEnv("PORT") // If process already using some port
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("HOST") // If process already using some host
	if !ok {
		host = defaultHost
	}

	log.Println("Server address: http://" + host + ":" + port)

	if err := runServer(net.JoinHostPort(host, port)); err != nil { // Try to start a server
		log.Println(err) // Log runtime errors
		os.Exit(1)       // Close the process
	}
}

// Method, which configuring a server and start a listening
func runServer(address string) (err error) {
	mux := http.NewServeMux()             // Create a multiplexer object for routing
	configServer := server.NewServer(mux) // Create an object for server configuration
	configServer.Initialize()             // Apply configuration

	workingServer := &http.Server{Addr: address, Handler: configServer} // Create a server object
	return workingServer.ListenAndServe()                               // Start listening
}
