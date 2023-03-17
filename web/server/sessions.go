package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/louisbranch/drake"
	"github.com/louisbranch/drake/web/presenter"
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
	printer, page := srv.i18n(w, r)

	switch r.Method {
	case "GET":

		name := r.URL.Path[len("/sessions/"):]
		if name == "" {
			sessions, err := srv.DB.FindSessions()
			if err != nil {
				srv.renderError(w, r, err)
				return
			}

			page.Title = printer.Sprintf("Sessions")
			page.Partials = []string{"sessions"}
			page.Content = struct {
				Sessions     []presenter.Session
				Latest       string
				Name         string
				Participants string
				Created      string
				None         string
				Back         string
			}{
				Sessions:     presenter.SessionsList(sessions, printer),
				Latest:       printer.Sprintf("Latest Sessions"),
				Name:         printer.Sprintf("Name"),
				Participants: printer.Sprintf("Participants"),
				Created:      printer.Sprintf("Created"),
				None:         printer.Sprintf("No available sessions"),
				Back:         printer.Sprintf("Back"),
			}

			srv.render(w, page)
			return
		}

		session, err := srv.DB.FindSession(name)
		if err != nil {
			srv.renderError(w, r, err)
			return
		}

		page.Title = printer.Sprintf("Session %s", name)
		page.Partials = []string{"session"}
		page.Content = struct {
			Session drake.Session
			Share   string
			Join    string
			Results string
		}{
			Session: session,
			Share:   printer.Sprintf("Share Link:"),
			Join:    printer.Sprintf("Join"),
			Results: printer.Sprintf("See Results"),
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
