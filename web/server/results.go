package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"

	"github.com/louisbranch/drake/web/presenter"
)

func (srv *Server) results(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/drake/results/"):]

	session, err := srv.DB.FindSession(name)
	if err != nil {
		srv.renderError(w, r, err)
		return
	}

	var token string
	at, err := r.Cookie("access_token")
	if err == nil {
		token = at.Value
	}

	survey, err := srv.DB.FindSurvey(session.ID, token)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		srv.renderError(w, r, err)
		return
	}

	surveys, err := srv.DB.FindSurveys(session.ID)
	if err != nil {
		srv.renderError(w, r, err)
		return
	}

	printer, page := srv.i18n(w, r)

	result := presenter.Result{
		Surveys: surveys,
		Printer: printer,
	}

	if r.Header.Get("Content-type") == "application/json" {

		data := struct {
			PreSurveyData  []int `json:"predata"`
			PostSurveyData []int `json:"postdata"`
		}{
			PreSurveyData:  result.PresurveyData(),
			PostSurveyData: result.PostsurveyData(),
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			srv.renderError(w, r, err)
			return
		}

		return
	}

	s, _ := json.Marshal(result.DataLabels())
	labels := template.JS(string(s))

	page.Title = printer.Sprintf("Results for Session %s", name)
	page.Partials = []string{"result"}
	page.Content = struct {
		Survey        presenter.Survey
		DataLabels    template.JS
		Predictions   string
		Results       string
		Civilizations string
		Participants  string
	}{
		Survey:        presenter.Survey{Survey: survey, Printer: printer},
		DataLabels:    labels,
		Predictions:   printer.Sprintf("Initial Predictions"),
		Results:       printer.Sprintf("Final Estimations"),
		Civilizations: printer.Sprintf("Civilizations"),
		Participants:  printer.Sprintf("Participants"),
	}

	srv.render(w, page)
}
