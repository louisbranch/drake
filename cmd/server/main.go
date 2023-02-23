package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
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

	files := os.Getenv("FILES_PATH")
	if files == "" {
		files = "web"
	}

	// TODO: config postgres db
	db, err := sqlite.New("drake.db")
	if err != nil {
		log.Fatal(err)
	}

	srv := &server.Server{
		DB:       db,
		Template: html.New(filepath.Join(files, "templates")),
		Assets:   http.FileServer(http.Dir(filepath.Join(files, "assets"))),
		Random:   rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	mux := srv.NewServeMux()

	log.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
