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
	ID      string
	HTML    template.HTML
	Classes []string
	Styles  map[string]string
}

// Render ...
func (h *Innerhtml) Render(w io.Writer) error {
	if h.ID == "" {
		h.ID = nextInnerhtmlID()
	}

	// copy styles
	styles := map[string]string{}
	if len(h.Styles) > 0 {
		styles = h.Styles
	}

	// default classes
	classess := map[string]bool{
		"x-innerhtml": true,
	}
	// copy classes
	for _, c := range h.Classes {
		if _, ok := classess[c]; !ok {
			classess[c] = true
		}
	}
	npClasses := []string{}
	for k := range classess {
		npClasses = append(npClasses, k)
	}

	div := &DivContainer{
		ID:      h.ID,
		Classes: npClasses,
		Styles:  styles,
		Items: Items{&RawHTML{
			HTML: h.HTML,
		}}, // don't layout again just copy items
	}
	return renderDiv(w, div)
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
	return render(w, `{{$.HTML}}`, h)
}

// GetID ...
func (h *RawHTML) GetID() string {
	return ""
}
