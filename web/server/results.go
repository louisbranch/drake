package server

import "net/http"

func (srv *Server) results(w http.ResponseWriter, r *http.Request) {
	srv.renderNotFound(w)
}
