package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/louisbranch/drake/sqlite"
	"github.com/louisbranch/drake/web/html"
	"github.com/louisbranch/drake/web/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := sqlite.New("drake.db")
	if err != nil {
		log.Fatal(err)
	}

	srv := &server.Server{
		DB:       db,
		Template: html.New("web/templates"),
		Random:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	mux := srv.NewServeMux()

	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
