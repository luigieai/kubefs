package server

import (
	"log"
	"net/http"
)

// var wg sync.WaitGroup

type Server struct{}

// key, cert, port, dir string for future use
func (s *Server) Serve() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := "this works"
		w.Write([]byte(str))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil
}
