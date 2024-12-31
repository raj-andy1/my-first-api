package http

import (
	"encoding/json"
	"log"
	name2 "my-first-api/internal/name"
	"net/http"
	"slices"
)

type name struct {
	Name string `json:"name"`
}

type message struct {
	Message string `json:"message"`
}

type Server struct {
	mux *http.ServeMux
}

func NewServer(namesvc *name2.Service) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /greet", func(w http.ResponseWriter, r *http.Request) {
		n, err := json.Marshal(namesvc.GetNames())
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = w.Write(n)
		if err != nil {
			log.Fatal(err)
		}
	})

	mux.HandleFunc("POST /greet", func(w http.ResponseWriter, r *http.Request) {
		var n name
		err := json.NewDecoder(r.Body).Decode(&n)
		if err != nil {
			log.Printf("Error: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if slices.Contains(namesvc.GetNames(), n.Name) {
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte("Name exists")); err != nil {
				log.Fatal(err)
			}

		} else {
			namesvc.Add(n.Name)
			w.WriteHeader(http.StatusCreated)
		}
		return
	})

	mux.HandleFunc("PUT /greet", func(w http.ResponseWriter, r *http.Request) {
		var n name
		var nn []byte
		err := json.NewDecoder(r.Body).Decode(&n)
		names := namesvc.GetNames()
		if err != nil {
			log.Printf("Error: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if slices.Contains(names, n.Name) {
			i := slices.Index(names, n.Name)
			names = append(names[:i], names[i+1:]...)
			nn, err = json.Marshal(names)
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = w.Write(nn)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(http.StatusCreated)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	})
	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
