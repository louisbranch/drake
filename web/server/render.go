package server

import (
	"net/http"
)

func (srv *Server) renderError(w http.ResponseWriter, r *http.Request, err error) {
	printer, page := srv.i18n(w, r)
	page.Title = printer.Sprintf("Internal Server Error")
	page.Content = err
	page.Partials = []string{"500"}

	w.WriteHeader(http.StatusInternalServerError)
	srv.render(w, page)
}

func (srv *Server) renderNotFound(w http.ResponseWriter, r *http.Request) {
	printer, page := srv.i18n(w, r)
	page.Title = printer.Sprintf("Page Not Found")
	page.Content = struct {
		Title string
		Home  string
	}{
		Title: printer.Sprintf("Page Not Found"),
		Home:  printer.Sprintf("Home"),
	}
	page.Partials = []string{"404"}

	w.WriteHeader(http.StatusNotFound)
	srv.render(w, page)
}
