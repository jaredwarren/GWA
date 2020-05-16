package ext

import (
	"fmt"
	"html/template"
	"io"
)

var (
	buttonID = 0
)

// Button ...
type Button struct {
	ID   string
	Text template.HTML
	// TODO: use ui.Eval()
	// Handler   template.JS
	Handler   template.JS
	UI        string // TODO
	IconClass string
	Parent    Renderer
}

// Render ...
func (b *Button) Render(w io.Writer) error {
	if b.ID == "" {
		b.ID = nextButtonID()
	}
	return renderTemplate(w, "button", b)
}

// SetParent ...
func (b *Button) SetParent(p Renderer) {
	b.Parent = p
}

func nextButtonID() string {
	id := fmt.Sprintf("button-%d", buttonID)
	buttonID++
	return id
}
