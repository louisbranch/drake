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
		uri, _ := url.JoinPath("", "results", name)
		http.Redirect(w, r, uri, http.StatusFound)
		return
	}

	switch r.Method {
	case "GET":
		printer, page := srv.i18n(w, r)
		page.Title = printer.Sprintf("Drake Equation - %s", name)
		page.Partials = []string{"survey"}

		type Option struct {
			Text  string
			Value float64
		}

		lvalues := []Option{}
		for _, n := range []int{100, 1000, 10000, 100000} {
			lvalues = append(lvalues, Option{
				Text:  printer.Sprintf("%d years", n),
				Value: float64(n),
			})

		}
		lvalues = append(lvalues, Option{
			Text:  printer.Sprintf("%d years or more", 1000000),
			Value: 1000000,
		})

		page.Content = struct {
			Session         drake.Session
			Survey          drake.Survey
			Next            string
			Choice          string
			N               string
			NValues         []Option
			R               string
			RValues         []Option
			Fp              string
			FpValues        []Option
			Ne              string
			NeValues        []Option
			Fl              string
			FlValues        []Option
			Fi              string
			FiValues        []Option
			Fc              string
			FcValues        []Option
			L               string
			LValues         []Option
			Estimation      string
			Difference      string
			Agreement       string
			AgreementValues []Option
		}{
			Session: session,
			Survey:  survey,
			Next:    printer.Sprintf("Next"),
			Choice:  printer.Sprintf("Select the choice that best agrees with what you think."),
			N:       printer.Sprintf("How many technological advanced civilizations exist in the Milky Way?"),
			NValues: []Option{
				{
					Text:  printer.Sprintf("Only us, we are all alone"),
					Value: 1,
				},
				{
					Text:  printer.Sprintf("A few dozens"),
					Value: 10,
				},
				{
					Text:  printer.Sprintf("A few hundreds"),
					Value: 100,
				},
				{
					Text:  printer.Sprintf("A few thousands"),
					Value: 1000,
				},
				{
					Text:  printer.Sprintf("A few millions"),
					Value: 1000000,
				},
			},
			R: printer.Sprintf("About how many stars are there in our galaxy?"),
			RValues: []Option{
				{
					Text:  printer.Sprintf("A few hundred"),
					Value: 100,
				},
				{
					Text:  printer.Sprintf("A few hundred thousand"),
					Value: 100000,
				},
				{
					Text:  printer.Sprintf("A few hundred million"),
					Value: 100000000,
				},
				{
					Text:  printer.Sprintf("A few hundred billion"),
					Value: 100000000000,
				},
				{
					Text:  printer.Sprintf("A few hundred trillion"),
					Value: 100000000000000,
				},
			},
			Fp: printer.Sprintf("Approximately what percentage of the stars in our galaxy have planets in orbit?"),
			FpValues: []Option{
				{
					Text:  printer.Sprintf("Very few stars have planets in orbit"),
					Value: 0.01,
				},
				{
					Text:  printer.Sprintf("20%%"),
					Value: 0.2,
				},
				{
					Text:  printer.Sprintf("40%%"),
					Value: 0.4,
				},
				{
					Text:  printer.Sprintf("60%%"),
					Value: 0.6,
				},
				{
					Text:  printer.Sprintf("80%%-90%%"),
					Value: 0.85,
				},
			},
			Ne: printer.Sprintf("On average, how many habitable (Earth-like) planets are there per planetary system?"),
			NeValues: []Option{
				{
					Text:  printer.Sprintf("Very few (not even one in every solar system)"),
					Value: 0.1,
				},
				{
					Text:  printer.Sprintf("1"),
					Value: 1,
				},
				{
					Text:  printer.Sprintf("2 or 3"),
					Value: 2.5,
				},
				{
					Text:  printer.Sprintf("4 or 5"),
					Value: 4.5,
				},
				{
					Text:  printer.Sprintf("More than 5"),
					Value: 5.5,
				},
			},
			Fl: printer.Sprintf("On what fraction of habitable planets will life develop?"),
			FlValues: []Option{
				{
					Text:  printer.Sprintf("A very small percentage, life is very rare"),
					Value: 0.001,
				},
				{
					Text:  printer.Sprintf("20%%"),
					Value: 0.2,
				},
				{
					Text:  printer.Sprintf("50%%"),
					Value: 0.5,
				},
				{
					Text:  printer.Sprintf("80%%"),
					Value: 0.8,
				},
				{
					Text:  printer.Sprintf("100%%%%, if conditions are favorable, life is inevitable"),
					Value: 1,
				},
			},
			Fi: printer.Sprintf("On what fraction of planets on which life develops will life evolve to intelligence?"),
			FiValues: []Option{
				{
					Text:  printer.Sprintf("A very small percentage"),
					Value: 0.001,
				},
				{
					Text:  printer.Sprintf("1%%-10%%"),
					Value: 0.05,
				},
				{
					Text:  printer.Sprintf("50%%"),
					Value: 0.5,
				},
				{
					Text:  printer.Sprintf("80%%"),
					Value: 0.8,
				},
				{
					Text:  printer.Sprintf("100%%"),
					Value: 1,
				},
			},
			Fc: printer.Sprintf("What fraction of planets on which life evolves to intelligence will the intelligence develop a technological civilization capable of radio communication?"),
			FcValues: []Option{
				{
					Text:  printer.Sprintf("Less than 1%%%%"),
					Value: 0.001,
				},
				{
					Text:  printer.Sprintf("About 10%%%%"),
					Value: 0.1,
				},
				{
					Text:  printer.Sprintf("20%%-30%%"),
					Value: 0.25,
				},
				{
					Text:  printer.Sprintf("50%%-60%%"),
					Value: 0.55,
				},
				{
					Text:  printer.Sprintf("100%%%%, a technological civilization will always eventually develop if life is present"),
					Value: 1,
				},
			},
			L:          printer.Sprintf("What is the average lifetime of a technological civilization capable of communication?"),
			LValues:    lvalues,
			Estimation: printer.Sprintf("You have estimated that there are %d civilizations in the Milky Way.", fprtToInt(survey.N)),
			Difference: printer.Sprintf("This is a difference of %d orders of magnitude from your initial prediction of %d.",
				survey.Difference(), fprtToInt(survey.PresurveyAssessment)),
			Agreement: printer.Sprintf("Do you agree with your most recent estimation?"),
			AgreementValues: []Option{
				{
					Text:  printer.Sprintf("Yes, I am more confident using the Drake Equation calculation"),
					Value: 1,
				},
				{
					Text:  printer.Sprintf("Yes, the Drake Equation includes factors I didn't think during my initial prediction"),
					Value: 1,
				},
				{
					Text:  printer.Sprintf("No, I am more confident with my initial prediction"),
					Value: 0,
				},
				{
					Text:  printer.Sprintf("No, any guess is good as mine"),
					Value: 0,
				},
			},
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
