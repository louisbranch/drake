package server

import "net/http"

func (srv *Server) about(w http.ResponseWriter, r *http.Request) {
	printer, page := srv.i18n(w, r)
	page.Title = printer.Sprintf("About")
	page.Partials = []string{"about"}
	page.Content = struct {
		About         string
		References    string
		Context       string
		Terms         string
		Contributions string
		Source        string
	}{
		About:         printer.Sprintf("About"),
		References:    printer.Sprintf("References"),
		Context:       printer.Sprintf("This project was created as part of the course, Principles and Practices in Science Education, at the University of Toronto with the intention of being a free resource for educators to introduce the Drake Equation to a wider audience."),
		Terms:         printer.Sprintf("The questions for the Drake Equation on the survey are from LoPresto and Hubble-Zdanowski (2012). The code is open-source using the MIT License."),
		Contributions: printer.Sprintf("If you would like to contribute to the project, for example, adding more translations, get in touch:"),
		Source:        printer.Sprintf("Source Code"),
	}

	srv.render(w, page)
}

func (srv *Server) astro(w http.ResponseWriter, r *http.Request) {
	_, page := srv.i18n(w, r)
	page.Layout = "astro"

	srv.render(w, page)
}

func (srv *Server) equation(w http.ResponseWriter, r *http.Request) {
	printer, page := srv.i18n(w, r)
	page.Title = printer.Sprintf("Drake Equation")
	page.Partials = []string{"equation"}
	page.Content = struct {
		Equation      string
		References    string
		Context       string
		Terms         string
		Contributions string
		Source        string
	}{
		Equation:      printer.Sprintf("The Equation"),
		References:    printer.Sprintf("References"),
		Context:       printer.Sprintf("This project was created as part of the course, Principles and Practices in Science Education, at the University of Toronto with the intention of being a free resource for educators to introduce the Drake Equation to a wider audience."),
		Terms:         printer.Sprintf("The questions for the Drake Equation on the survey are from LoPresto and Hubble-Zdanowski (2012). The code is open-source using the MIT License."),
		Contributions: printer.Sprintf("If you would like to contribute to the project, for example, adding more translations, get in touch:"),
		Source:        printer.Sprintf("Source Code"),
	}

	srv.render(w, page)
}
