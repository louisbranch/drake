package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/louisbranch/drake/sqlite"
	"github.com/louisbranch/drake/web/html"
	"github.com/louisbranch/drake/web/server"
)

func main() {
	db, err := sqlite.New("drake.db")

	if err != nil {
		log.Fatal(err)
	}

	srv := &server.Server{
		DB:                 db,
		Template:           html.New("web/templates"),
	}
	mux := srv.NewServeMux()

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
