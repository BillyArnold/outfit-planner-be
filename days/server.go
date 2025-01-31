package days_server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	port string
}

type ServerOpts struct {
	port string
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		port: opts.port,
	}
}

func (s *Server) Start() {
	fmt.Printf("starting days server on port %s \n", s.port)
	r := mux.NewRouter()

	sr := r.PathPrefix("/api/v1").Subrouter()
	sr.HandleFunc("/", TestHandler)

	log.Fatal(http.ListenAndServe(s.port, r))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{'d'})
}
