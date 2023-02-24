package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/louisbranch/drake"
	"github.com/louisbranch/drake/db/postgres"
	"github.com/louisbranch/drake/db/sqlite"
	"github.com/louisbranch/drake/web/html"
	"github.com/louisbranch/drake/web/server"
)

func main() {

	dev := true
	if os.Getenv("APP_ENV") == "production" {
		dev = false
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	files := os.Getenv("FILES_PATH")
	if files == "" {
		files = "web"
	}

	var db drake.Database
	var err error

	dbuser := os.Getenv("POSTGRES_USER")
	if dbuser == "" {
		log.Println("using sqlite database")
		db, err = sqlite.New("drake.db")
	} else {
		log.Println("using postgres database")
		pswd := os.Getenv("POSTGRES_PASSWORD")
		host := os.Getenv("POSTGRES_HOSTNAME")
		dbname := os.Getenv("POSTGRES_DB")

		sslmode := "verify-full"
		if dev {
			sslmode = "disable"
		}

		connection := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
			dbuser, pswd, host, dbname, sslmode)
		db, err = postgres.New(connection)
	}
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
