package ext

import (
	"fmt"
	"html/template"
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
func (h *Header) Render() template.HTML {
	// nothing to render
	if h.Title == "" && len(h.Items) == 0 {
		return ""
	}

	if h.ID == "" {
		h.ID = nextHeaderID()
	}

	// default classes
	h.Classes = []string{
		"x-panelheader",
		"x-container",
		"x-component",
		"x-docked-top",
		"x-horizontal",
		"x-noborder-tr",
	}

	return render("header", h)
}

func nextHeaderID() string {
	id := fmt.Sprintf("%d", headerID)
	headerID++
	return id
}
