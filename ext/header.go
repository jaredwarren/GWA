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
func NewHeader(title template.HTML) *Header {
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
	XType     string
	ID        string // how to auto generate
	Title     template.HTML
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
	Styles    map[string]string
}

// // Render ...
// func (h *Header) Render(w io.Writer) error {
// 	if h.ID == "" {
// 		h.ID = nextHeaderID()
// 	}
// 	// TODO: add stuff from header.html as items
// 	return renderTemplate(w, "header", h)
// }

// Render ...
func (h *Header) Render(w io.Writer) error {
	if h.ID == "" {
		h.ID = nextHeaderID()
	}
	// default classes
	classess := map[string]bool{
		"x-panel": true,
	}
	// copy classes
	for _, c := range h.Classes {
		if _, ok := classess[c]; !ok {
			classess[c] = true
		}
	}
	// if h.Shadow {
	// 	classess["x-shadow"] = true
	// }

	// copy styles
	styles := map[string]string{}
	if len(h.Styles) > 0 {
		styles = h.Styles
	}

	// append new styles based on p's properties
	// what if I want width to be 0px?
	if h.Width != 0 && h.Docked != "top" && h.Docked != "bottom" {
		styles["width"] = fmt.Sprintf("%dpx", h.Width)
		classess["x-widthed"] = true
	}
	// what if I want height to be 0px?
	if h.Height != 0 && h.Docked != "left" && h.Docked != "right" {
		styles["height"] = fmt.Sprintf("%dpx", h.Height)
		classess["x-heighted"] = true
	}
	if h.Border != "" {
		styles["border"] = string(h.Border)
		classess["x-managed-border"] = true
	}

	// convert class back to array
	npClasses := []string{}
	for k := range classess {
		npClasses = append(npClasses, k)
	}

	// ITEMS
	items := Items{}

	// HEADER
	if h.Title != "" {
		title := &Innerhtml{
			HTML: h.Title,
			Styles: map[string]string{
				"flex":       "1",
				"text-align": "center",
			},
		}
		items = append(items, title)
	}
	// append title

	// append rest of items
	if len(h.Items) > 0 {
		items = append(items, h.Items...)
	}

	// TODO: if panel has "layout" set that up here
	// // This layout should only apply to non-docked items!

	for _, i := range items {
		c, ok := i.(Child)
		if ok {
			c.SetParent(h)
		}
	}

	layout := &Layout{
		Items: items,
		Type:  "vbox",
		Pack:  "center",
		Align: "center",
	}

	div := &DivContainer{
		ID:      h.ID,
		Classes: npClasses,
		Styles:  styles,
		// Items:   LayoutItems(items),
		Items: Items{layout},
	}
	return renderDiv(w, div)
}

// GetID ...
func (h *Header) GetID() string {
	return h.ID
}

// GetDocked ...
func (h *Header) GetDocked() string {
	return h.Docked
}

// SetStyle ...
func (h *Header) SetStyle(key, value string) {
	if h.Styles == nil {
		h.Styles = map[string]string{}
	}
	h.Styles[key] = value
}

func nextHeaderID() string {
	id := fmt.Sprintf("header-%d", headerID)
	headerID++
	return id
}
