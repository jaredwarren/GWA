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
	nh, _ := h.Build()
	return render(w, "innerhtml", nh)
}

// Build copys info to a new panel
func (h *Innerhtml) Build() (Renderer, error) {
	nh := &Innerhtml{}
	if h.ID != "" {
		nh.ID = h.ID
	} else {
		nh.ID = nextInnerhtmlID()
	}
	nh.HTML = h.HTML
	return nh, nil
}

// Debug ...
func (h *Innerhtml) Debug() {}

func nextInnerhtmlID() string {
	id := fmt.Sprintf("%d", innerhtmlID)
	innerhtmlID++
	return id
}
