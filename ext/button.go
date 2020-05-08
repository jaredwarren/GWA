package ext

import (
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
	nb, _ := b.Build()
	return render(w, "button", nb)
}

// Build copys info to a new panel
func (b *Button) Build() (*Button, error) {
	n := &Button{}
	if b.ID != "" {
		n.ID = b.ID
	} else {
		n.ID = nextInnerhtmlID()
	}
	return n, nil
}

// Debug ...
func (b *Button) Debug() {}
