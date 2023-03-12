package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"

	"github.com/louisbranch/drake"
)

func (srv *Server) results(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/results/"):]

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

	result := drake.Result{
		Surveys: surveys,
	}

	s, _ := json.Marshal(result.Buckets())
	buckets := template.JS(string(s))

	s, _ = json.Marshal(result.PresurveyData())
	predata := template.JS(string(s))

	s, _ = json.Marshal(result.PostsurveyData())
	postdata := template.JS(string(s))

	printer, page := srv.i18n(w, r)
	page.Title = printer.Sprintf("Results for Session %s", name)
	page.Partials = []string{"result"}
	page.Content = struct {
		Session        drake.Session
		Survey         drake.Survey
		Buckets        template.JS
		PresurveyData  template.JS
		PostsurveyData template.JS
		Guesses        string
		Guess          string
		Results        string
		Estimation     string
		Civilizations  string
		Participants   string
	}{
		Session:        session,
		Survey:         survey,
		Buckets:        buckets,
		PresurveyData:  predata,
		PostsurveyData: postdata,
		Guesses:        printer.Sprintf("Initial Guesses"),
		Guess: printer.Sprintf("You guessed %d civilizations",
			fprtToInt(survey.PresurveyAssessment)),
		Results: printer.Sprintf("Final Results"),
		Estimation: printer.Sprintf("You estimated %d civilizations",
			fprtToInt(survey.N)),
		Civilizations: printer.Sprintf("Civilizations"),
		Participants:  printer.Sprintf("Participants"),
	}

	srv.render(w, page)
}

func fprtToInt(n *float64) int {
	if n == nil {
		return 0
	}

	return int(*n)
}
