package server

import (
	"log"
	"net/http" // net/http starting from Go version 1.6 automatically supports HTTP/2
)

// Server struct with mux field
type Server struct {
	mux *http.ServeMux // HTTP request multiplexer
}

// NewServer returns pointer on a new object
func NewServer(mux *http.ServeMux) *Server {
	return &Server{mux: mux}
}

// Initialize function add a configuration to a server object
func (server *Server) Initialize() {
	server.mux.HandleFunc("/api/slow", server.slow) // POST request
}

// ServeHTTP is a method, which implements http.Handler Property and allows to use http.Server.ListenAndServe method
func (server *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	server.mux.ServeHTTP(writer, request)
}

func (server *Server) slow(writer http.ResponseWriter, request *http.Request) {
	keys, ok := request.URL.Query()["timeout"] // Get parameters from request with name "timeout"

	if !ok || len(keys[0]) < 1 { // If count of parameters with name "timeout" less, than 1
		log.Println("Url Param 'tineout' is missing")
		return
	}

	// Default status: 200 OK
	writer.Header().Add("Content-Type", "application/json")
	writer.Write([]byte("status: ok")) // Return result
}
