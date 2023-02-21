package server

import (
	"fmt"
	"net/http"

	"github.com/louisbranch/drake"
	"github.com/louisbranch/drake/web"
)

type Server struct {
	DB       drake.Database
	Template web.Template
}

func (srv *Server) NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("web/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", srv.index)

	mux.HandleFunc("/sessions/", srv.sessions)

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
		Partials: []string{"400"},
	}
	srv.render(w, page)
}

func (srv *Server) index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	page := web.Page{
		Title:      "Index",
		ActiveMenu: "index",
		//Content:    content,
		Partials: []string{"index"},
	}
	srv.render(w, page)
}
