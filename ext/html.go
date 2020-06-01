package ext

import (
	"fmt"
	"html/template"
	"io"
)

var (
	innerhtmlID = 0
)

// Innerhtml ...
type Innerhtml struct {
	ID      string
	HTML    template.HTML
	Classes Classes
	Styles  Styles
}

// Render ...
func (h *Innerhtml) Render(w io.Writer) error {
	if h.ID == "" {
		h.ID = nextInnerhtmlID()
	}

	h.Classes = append(h.Classes, "x-innerhtml")

	div := &DivContainer{
		ID:      h.ID,
		Classes: h.Classes,
		Styles:  h.Styles,
		Items: Items{&RawHTML{
			HTML: h.HTML,
		}},
	}
	return div.Render(w)
}

// GetID ...
func (h *Innerhtml) GetID() string {
	return h.ID
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
	_, err := w.Write([]byte(h.HTML))
	return err
}

// GetID ...
func (h *RawHTML) GetID() string {
	return ""
}
