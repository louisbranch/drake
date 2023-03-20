package server

import (
	"encoding/json"
	"net/http"

	"github.com/louisbranch/drake/web/presenter"
)

func (srv *Server) statistics(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/drake/statistics/"):]

	session, err := srv.DB.FindSession(name)
	if err != nil {
		srv.renderError(w, r, err)
		return
	}

	surveys, err := srv.DB.FindSurveys(session.ID)
	if err != nil {
		srv.renderError(w, r, err)
		return
	}

	printer, page := srv.i18n(w, r)

	stats := presenter.NewStatistics(surveys, printer)

	if r.Header.Get("Content-type") == "application/json" {

		data := stats.Data()

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			srv.renderError(w, r, err)
			return
		}

		return
	}

	page.Title = printer.Sprintf("Stastics for Session %s", name)
	page.Partials = []string{"statistics"}
	page.Content = struct {
		Survey       presenter.Survey
		Statistics   string
		Options      string
		Participants string
		Labels       [][]string
	}{
		Survey:       stats.Survey,
		Statistics:   printer.Sprintf("Statistics"),
		Options:      printer.Sprintf("Options"),
		Participants: printer.Sprintf("Participants"),
		Labels:       stats.Labels,
	}

	srv.render(w, page)
}
