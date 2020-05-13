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
		Width:  300,
		Height: 200,
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
	RenderTo  string // type???
	Header    *Header
	Border    template.CSS
	Docked    string // top, bottom, left, right
	Flex      int
	Style     string
	Classes   []string
}

// Render ...
func (h *Header) Render(w io.Writer) error {
	fmt.Print("  render header:")
	n := &Header{}
	if h.ID != "" {
		n.ID = h.ID
	} else {
		n.ID = nextInnerhtmlID()
	}
	fmt.Println(n.ID) // show id
	n.Title = h.Title
	n.Docked = h.Docked
	return renderTemplate(w, "header", n)
}

// GetDocked ...
func (h *Header) GetDocked() string {
	return h.Docked
}

func nextHeaderID() string {
	id := fmt.Sprintf("%d", headerID)
	headerID++
	return id
}
