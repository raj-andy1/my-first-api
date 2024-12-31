package main

import (
	"log"
	name "my-first-api/internal/name"
	transport "my-first-api/internal/transport"
)

func main() {
	svc := name.NewService()
	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}

}
