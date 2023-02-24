package server

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/louisbranch/drake"
	"github.com/louisbranch/drake/db/sqlite"
	"github.com/louisbranch/drake/web"
)

func (srv *Server) surveys(w http.ResponseWriter, r *http.Request, name string) {

	session, err := srv.DB.FindSession(name)
	if err != nil {
		srv.renderError(w, err)
		return
	}

	var token string
	at, err := r.Cookie("access_token")
	if err == nil {
		token = at.Value
	} else {
		b := make([]byte, 8)
		srv.Random.Read(b)
		token = fmt.Sprintf("%x", b)
		http.SetCookie(w, &http.Cookie{
			Name:   "access_token",
			Value:  token,
			MaxAge: 24 * 60 * 60, // 24 hours
		})
	}

	survey, err := srv.DB.FindSurvey(session.ID, token)
	if errors.Is(err, sqlite.ErrNotFound) {
		err = srv.DB.CreateSurvey(&drake.Survey{
			SessionID:   session.ID,
			AccessToken: token,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}
	if err != nil {
		srv.renderError(w, err)
		return
	}

	switch r.Method {
	case "GET":
		content := struct {
			Session drake.Session
			Survey  drake.Survey
		}{
			Session: session,
			Survey:  survey,
		}

		page := web.Page{
			Title:    fmt.Sprintf("Drake Equation - %s", name),
			Content:  content,
			Partials: []string{"survey"},
		}
		srv.render(w, page)
	case "POST":
		err := r.ParseForm()
		if err != nil {
			srv.renderError(w, err)
			return
		}

		form := r.PostForm

		if val := powers(form.Get("presurvey_assessment")); val != nil {
			survey.PresurveyAssessment = val
		}
		if val := powers(form.Get("r")); val != nil {
			survey.R = val
		}
		if val := fraction(form.Get("fp")); val != nil {
			survey.Fp = val
		}
		if val := powers(form.Get("ne")); val != nil {
			survey.Ne = val
		}
		if val := fraction(form.Get("fl")); val != nil {
			survey.Fl = val
		}
		if val := fraction(form.Get("fi")); val != nil {
			survey.Fi = val
		}
		if val := fraction(form.Get("fc")); val != nil {
			survey.Fc = val
		}
		if val := powers(form.Get("l")); val != nil {
			survey.L = val
			survey.Result()
		}
		survey.UpdatedAt = time.Now()

		err = srv.DB.UpdateSurvey(&survey)
		if err != nil {
			srv.renderError(w, err)
			return
		}

		http.Redirect(w, r, name, http.StatusFound)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func powers(value string) *float64 {
	if value == "" {
		return nil
	}
	n := strings.Count(value, "0")
	p := math.Pow(10, float64(n))
	return &p
}

func fraction(value string) *float64 {
	if value == "" {
		return nil
	}
	n, _ := strconv.ParseFloat(value, 64)
	return &n
}
