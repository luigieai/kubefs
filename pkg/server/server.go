package server

import (
	"log"
	"net/http"
)

// var wg sync.WaitGroup

type Server struct {
	Port string
	Dir  string
}

// key, cert, port, dir string for future use
func (s *Server) Serve(dir, port string) error {

	s.Port = port
	s.Dir = dir

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := "this works"
		w.Write([]byte(str))
	})

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return nil
}
