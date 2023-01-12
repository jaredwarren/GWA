package gbt

import (
	"fmt"
	"html/template"
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
func (h *Innerhtml) Render() Stringer {
	if h.ID == "" {
		h.ID = nextInnerhtmlID()
	}

	h.Classes = append(h.Classes, "x-innerhtml")

	div := &DivContainer{
		ID:      h.ID,
		Classes: h.Classes,
		Styles:  h.Styles,
		Items:   Items{RawHTML(h.HTML)},
	}
	return div.Render()
}

func nextInnerhtmlID() string {
	id := fmt.Sprintf("innerhtml-%d", innerhtmlID)
	innerhtmlID++
	return id
}

// RawHTML ...
type RawHTML template.HTML

// Render ...
func (h RawHTML) Render() Stringer {
	return template.HTML(h)
}

// GetID ...
func (h *RawHTML) GetID() string {
	return ""
}
