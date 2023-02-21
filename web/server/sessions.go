package server

import (
	"errors"
	"net/http"

	"github.com/louisbranch/drake/web"
)

func (srv *Server) sessions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		page := web.Page{
			Title:      "Sessions",
			ActiveMenu: "sessions",
			//Content:    content,
			Partials: []string{"sessions"},
		}
		srv.render(w, page)
	case "POST":
		srv.renderError(w, errors.New("not implemented"))
	case "":
		srv.renderNotFound(w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
