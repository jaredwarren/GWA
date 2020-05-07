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
	return render(w, "innerhtml", h)
}

// Debug ...
func (h *Innerhtml) Debug() {}

func nextInnerhtmlID() string {
	id := fmt.Sprintf("%d", innerhtmlID)
	innerhtmlID++
	return id
}
