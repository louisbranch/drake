package server

import (
	"net/http"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	_ "github.com/louisbranch/drake/translations"
	"github.com/louisbranch/drake/web"
)

func (s *Server) i18n(w http.ResponseWriter, r *http.Request) (*message.Printer, web.Page) {

	var lang language.Tag

	query := r.URL.Query().Get("lang")
	cookie, err := r.Cookie("lang")
	if err != nil || query != "" {
		http.SetCookie(w, &http.Cookie{
			Name:   "lang",
			Value:  query,
			MaxAge: 24 * 60 * 60 * 365, // 1 year
		})
	}

	if query == "" && cookie != nil {
		query = cookie.Value
	}

	switch query {
	case "pt-BR":
		lang = language.MustParse("pt-BR")
	default:
		lang = language.MustParse("en")
	}

	printer := message.NewPrinter(lang)

	page := web.Page{
		Header:  printer.Sprintf("Drake Equation"),
		Website: printer.Sprintf("Astronomy Education"),
	}

	return printer, page
}
