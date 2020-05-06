package ext

import (
	"fmt"
	"html/template"
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
	ID   string // how to auto generate
	HTML template.HTML
}

// Render ...
func (h *Innerhtml) Render() template.HTML {

	if h.ID == "" {
		h.ID = nextInnerhtmlID()
	}

	return render("innerhtml", h)
}

// Debug ...
func (h *Innerhtml) Debug() {}

func nextInnerhtmlID() string {
	id := fmt.Sprintf("%d", innerhtmlID)
	innerhtmlID++
	return id
}
