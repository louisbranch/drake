package server

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/louisbranch/drake"
	"github.com/louisbranch/drake/web"
)

type Server struct {
	DB       drake.Database
	Template web.Template
	Random   *rand.Rand
}

func (srv *Server) NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/sessions/", srv.sessions)
	mux.HandleFunc("/results/", srv.results)

	mux.HandleFunc("/", srv.index)

	return mux
}

func (srv *Server) render(w http.ResponseWriter, page web.Page) {
	if page.Layout == "" {
		page.Layout = "layout"
	}

	err := srv.Template.Render(w, page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
	}
}

func (srv *Server) renderError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	page := web.Page{
		Title:    "500",
		Content:  err,
		Partials: []string{"500"},
	}
	srv.render(w, page)
}

func (srv *Server) renderNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	page := web.Page{
		Title:    "Not Found",
		Partials: []string{"404"},
	}
	srv.render(w, page)
}

func (srv *Server) index(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/"):]

	if name == "" {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		page := web.Page{
			Title:      "Drake Equation",
			ActiveMenu: "index",
			Partials:   []string{"index"},
		}
		srv.render(w, page)
		return
	}

	srv.surveys(w, r, name)
}
