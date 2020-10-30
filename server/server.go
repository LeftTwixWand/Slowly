package server

import (
	"encoding/json"
	"log"
	"net/http" // net/http starting from Go version 1.6 automatically supports HTTP/2
	"strconv"
	"time"

	"../dto/resulterror"
	"../dto/resultok"
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

	writer.Header().Add("Content-Type", "application/json") // Add a header with the Content-Type

	if number, err := strconv.Atoi(keys[0]); err == nil {
		if number > 5000 {
			writer.WriteHeader(http.StatusBadRequest) // Add status 400 Bad Request to a returned request
			json.NewEncoder(writer).Encode(resulterror.ResultError{Error: "timeout too long"})
		} else {
			time.Sleep(time.Duration(number) * time.Millisecond) // We can use time.Sleep function to set a timeout, because request is async

			// Default status: 200 OK
			json.NewEncoder(writer).Encode(resultok.ResultOk{Status: "ok"}) // Return result
		}

	} else {
		log.Println("String parsing error!")
	}
}
