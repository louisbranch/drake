package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/louisbranch/drake"
	"github.com/louisbranch/drake/web"
)

var alphanum = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (srv *Server) newSession() drake.Session {
	b := make([]rune, 6)
	for i := range b {
		b[i] = alphanum[srv.Random.Intn(len(alphanum))]
	}
	name := fmt.Sprintf("%s-%s", string(b[:3]), string(b[3:]))

	session := drake.Session{
		Name:      name,
		CreatedAt: time.Now(),
	}

	return session
}

func (srv *Server) sessions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		name := r.URL.Path[len("/sessions/"):]
		if name == "" {
			sessions, err := srv.DB.FindSessions()
			if err != nil {
				srv.renderError(w, r, err)
				return
			}

			content := struct {
				Sessions []drake.Session
			}{
				Sessions: sessions,
			}

			page := web.Page{
				Title:    "Sessions",
				Content:  content,
				Partials: []string{"sessions"},
			}

			srv.render(w, page)
			return
		}

		session, err := srv.DB.FindSession(name)
		if err != nil {
			srv.renderError(w, r, err)
			return
		}

		content := struct {
			Session drake.Session
		}{
			Session: session,
		}

		page := web.Page{
			Title:    fmt.Sprintf("Session %s", name),
			Content:  content,
			Partials: []string{"session"},
		}
		srv.render(w, page)

	case "POST":
		session := srv.newSession()

		err := srv.DB.CreateSession(&session)
		if err != nil {
			srv.renderError(w, r, err)
			return
		}

		uri := fmt.Sprintf("/sessions/%s", session.Name)

		http.Redirect(w, r, uri, http.StatusFound)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
