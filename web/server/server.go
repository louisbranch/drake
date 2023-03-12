package server

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/louisbranch/drake"
	"github.com/louisbranch/drake/web"
)

type Server struct {
	DB       drake.Database
	Template web.Template
	Assets   http.Handler
	Random   *rand.Rand
}

func (srv *Server) NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", srv.Assets))

	mux.HandleFunc("/sessions/", srv.sessions)
	mux.HandleFunc("/results/", srv.results)

	mux.HandleFunc("/", srv.index)

	return mux
}

func (srv *Server) render(w http.ResponseWriter, page web.Page) {
	if page.Layout == "" {
		page.Layout = "layout"
	}

	err := srv.Template.Render(w, page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
	}
}

func (srv *Server) index(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/"):]

	if name != "" {
		srv.surveys(w, r, name)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	printer, page := srv.i18n(w, r)
	page.Title = printer.Sprintf("Drake Equation")
	page.Partials = []string{"index"}
	page.Content = struct {
		Title     string
		Subtitle  string
		Intro     string
		Goals     string
		GoalItems []string
		Create    string
		Join      string
	}{
		Title:    printer.Sprintf("Drake Equation"),
		Subtitle: printer.Sprintf("Are we alone in the Universe?"),
		Intro:    printer.Sprintf("Estimate the number of detectable alien civilizations in the Milky Way using the Drake Equation."),
		Goals:    printer.Sprintf("Learning Goals:"),
		GoalItems: []string{
			printer.Sprintf("to think about the size and composition of the galaxy and how it affects the possibility of intelligent life"),
			printer.Sprintf("to understand and estimate the terms of the Drake Equation"),
			printer.Sprintf("to compare your initial guess with the final value"),
		},
		Create: printer.Sprintf("Create Session"),
		Join:   printer.Sprintf("Join Session"),
	}

	srv.render(w, page)
}
