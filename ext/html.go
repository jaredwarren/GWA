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
	n := &Innerhtml{}
	if h.ID != "" {
		n.ID = h.ID
	} else {
		n.ID = nextInnerhtmlID()
	}
	n.HTML = h.HTML
	return render(w, "innerhtml", n)
}

func nextInnerhtmlID() string {
	id := fmt.Sprintf("%d", innerhtmlID)
	innerhtmlID++
	return id
}
