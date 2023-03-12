package web

import (
	"io"
)

type Page struct {
	Title    string
	Header   string
	Website  string
	Layout   string
	Partials []string
	Content  interface{}
}

type Template interface {
	Render(w io.Writer, page Page) error
}
