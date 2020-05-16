package ext

import (
	"fmt"
	"html/template"
	"io"
)

var (
	innerhtmlID = 0
)

// NewInnerhtml ...
func NewInnerhtml(title string) *Innerhtml {
	return &Innerhtml{
		ID: nextInnerhtmlID(),
	}
}

// Innerhtml ...
type Innerhtml struct {
	ID   string
	HTML template.HTML
}

// Render ...
func (h *Innerhtml) Render(w io.Writer) error {
	if h.ID == "" {
		h.ID = nextInnerhtmlID()
	}

	div := &DivContainer{
		ID:      h.ID,
		Classes: []string{"x-innerhtml"},
		Items: Items{&RawHTML{
			HTML: h.HTML,
		}}, // don't layout again just copy items
	}
	return renderDiv(w, div)
}

func nextInnerhtmlID() string {
	id := fmt.Sprintf("innerhtml-%d", innerhtmlID)
	innerhtmlID++
	return id
}

// RawHTML ...
type RawHTML struct {
	HTML template.HTML
}

// Render ...
func (h *RawHTML) Render(w io.Writer) error {
	return render(w, `{{$.HTML}}`, h)
}
