package ext

import (
	"bytes"
	"fmt"
	"html/template"
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
func (b *Button) Render() template.HTML {

	// prepend newline so html/js formatting works
	if b.Handler != "" {
		// b.Handler = "*/\n" + b.Handler + "\n/*"
	}

	b.ID = fmt.Sprintf("%d", buttonID)
	buttonID++

	buf := new(bytes.Buffer)
	templates := template.Must(template.ParseFiles("templates/button.html"))
	templates.ExecuteTemplate(buf, "base", b)
	return template.HTML(buf.String())
}

// Debug ...
func (b *Button) Debug() {}
