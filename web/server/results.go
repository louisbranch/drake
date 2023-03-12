package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"github.com/louisbranch/drake"
	"github.com/louisbranch/drake/web"
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

	content := struct {
		Session        drake.Session
		Survey         drake.Survey
		Buckets        template.JS
		PresurveyData  template.JS
		PostsurveyData template.JS
	}{
		Session:        session,
		Survey:         survey,
		Buckets:        buckets,
		PresurveyData:  predata,
		PostsurveyData: postdata,
	}

	page := web.Page{
		Title:    fmt.Sprintf("Results for Session %s", name),
		Content:  content,
		Partials: []string{"result"},
	}
	srv.render(w, page)
}
