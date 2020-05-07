package ext

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
)

var (
	buttonID = 0
)

// Button ...
type Button struct {
	ID        string
	Text      template.HTML
	Handler   template.JS
	UI        string // TODO
	IconClass string
}

// Render ...
func (b *Button) Render(w io.Writer) error {

	// prepend newline so html/js formatting works
	if b.Handler != "" {
		// b.Handler = "*/\n" + b.Handler + "\n/*"
	}

	b.ID = fmt.Sprintf("%d", buttonID)
	buttonID++

	buf := new(bytes.Buffer)
	templates := template.Must(template.ParseFiles("templates/button.html"))
	return templates.ExecuteTemplate(buf, "base", b)
}

// Debug ...
func (b *Button) Debug() {}
