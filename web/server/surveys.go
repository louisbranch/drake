package server

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/louisbranch/drake"
	"github.com/louisbranch/drake/web/presenter"
)

func (srv *Server) surveys(w http.ResponseWriter, r *http.Request, name string) {

	session, err := srv.DB.FindSession(name)
	if errors.Is(err, sql.ErrNoRows) {
		srv.renderNotFound(w, r)
		return
	} else if err != nil {
		srv.renderError(w, r, err)
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
			Path:   "/drake",
			MaxAge: 24 * 60 * 60, // 24 hours
		})
	}

	survey, err := srv.DB.FindSurvey(session.ID, token)
	if errors.Is(err, sql.ErrNoRows) {
		err = srv.DB.CreateSurvey(&drake.Survey{
			SessionID:   session.ID,
			AccessToken: token,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}
	if err != nil {
		srv.renderError(w, r, err)
		return
	}

	if survey.NextQuestion() == "" {
		uri, _ := url.JoinPath("results", name)
		http.Redirect(w, r, uri, http.StatusFound)
		return
	}

	switch r.Method {
	case "GET":
		printer, page := srv.i18n(w, r)
		page.Title = printer.Sprintf("Drake Equation - %s", name)
		page.Partials = []string{"survey"}

		survey := presenter.Survey{
			Survey:  survey,
			Printer: printer,
		}

		page.Content = struct {
			Session drake.Session
			Survey  presenter.Survey
			Next    string
			Choice  string
		}{
			Session: session,
			Survey:  survey,
			Next:    printer.Sprintf("Next"),
			Choice:  printer.Sprintf("Select the choice that best agrees with what you think."),
		}

		srv.render(w, page)
	case "POST":
		err := r.ParseForm()
		if err != nil {
			srv.renderError(w, r, err)
			return
		}

		form := r.PostForm

		if val := ptrFloat(form.Get("presurvey_assessment")); val != nil {
			survey.PresurveyAssessment = val
		}
		if val := ptrFloat(form.Get("r")); val != nil {
			survey.R = val
		}
		if val := ptrFloat(form.Get("fp")); val != nil {
			survey.Fp = val
		}
		if val := ptrFloat(form.Get("ne")); val != nil {
			survey.Ne = val
		}
		if val := ptrFloat(form.Get("fl")); val != nil {
			survey.Fl = val
		}
		if val := ptrFloat(form.Get("fi")); val != nil {
			survey.Fi = val
		}
		if val := ptrFloat(form.Get("fc")); val != nil {
			survey.Fc = val
		}
		if val := ptrFloat(form.Get("l")); val != nil {
			survey.L = val
			survey.Result()
		}
		if val := prtBool(form.Get("learn_gain")); val != nil {
			survey.PostsurveyLearnGain = val
		}
		if val := prtTxt(form.Get("reason")); val != nil {
			survey.PostsurveyReason = val
		}
		survey.UpdatedAt = time.Now()

		err = srv.DB.UpdateSurvey(&survey)
		if err != nil {
			srv.renderError(w, r, err)
			return
		}

		http.Redirect(w, r, name, http.StatusFound)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ptrFloat(value string) *float64 {
	if value == "" {
		return nil
	}
	n, err := strconv.ParseFloat(value, 64)
	if err == nil {
		return &n
	}
	return nil
}

func prtBool(value string) *bool {
	switch value {
	case "true", "1":
		b := true
		return &b
	case "false", "0":
		b := false
		return &b
	default:
		return nil
	}
}

func prtTxt(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
