package ext

import (
	"fmt"
	"html/template"
	"io"
)

var (
	headerID = 0
)

// NewHeader ...
func NewHeader(title string) *Header {
	return &Header{
		ID:     nextHeaderID(),
		Title:  title,
		Border: template.CSS("1px solid lightgrey"),
		Docked: "top",
	}
}

// Header ...
// TODO: I don't need all of this crap for container
type Header struct {
	ID        string // how to auto generate
	Title     string
	IconClass string
	Layout    string
	HTML      template.HTML
	Width     int // float?
	Height    int // float?
	Items     []Renderer
	Header    *Header
	Border    template.CSS
	Docked    string // top, bottom, left, right
	Flex      int
	Style     string
	Classes   []string
}

// Render ...
func (h *Header) Render(w io.Writer) error {
	if h.ID == "" {
		h.ID = nextInnerhtmlID()
	}
	// TODO: add stuff from header.html as items
	return renderTemplate(w, "header", h)
}

// GetDocked ...
func (h *Header) GetDocked() string {
	return h.Docked
}

func nextHeaderID() string {
	id := fmt.Sprintf("header-%d", headerID)
	headerID++
	return id
}
