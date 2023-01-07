package gbt

import (
	"fmt"
	"html/template"
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
	XType     string            `json:"xtype"`
	ID        string            `json:"id,omitempty"` // how to auto generate
	Title     template.HTML     `json:"title,omitempty"`
	IconClass string            `json:"iconClass,omitempty"`
	Layout    string            `json:"layout,omitempty"`
	HTML      template.HTML     `json:"html,omitempty"`
	Width     int               `json:"width,omitempty"`
	Height    int               `json:"height,omitempty"`
	Items     []Renderer        `json:"items,omitempty"`
	Border    template.CSS      `json:"border,omitempty"`
	Docked    string            `json:"docked,omitempty"` // top, bottom, left, right
	Style     string            `json:"style,omitempty"`
	Classes   []string          `json:"classes,omitempty"`
	Styles    map[string]string `json:"styles,omitempty"`
	Shadow    bool              `json:"shadow,omitempty"`
}

// Render ...
func (h *Header) Render() Stringer {
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
	if h.Shadow {
		classess["x-shadow"] = true
	}

	// copy styles
	styles := map[string]string{}
	if len(h.Styles) > 0 {
		styles = h.Styles
	}

	// append new styles based on p's properties
	if h.Width != 0 && h.Docked != "top" && h.Docked != "bottom" {
		styles["width"] = fmt.Sprintf("%dpx", h.Width)
		classess["x-widthed"] = true
	}
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
		Items:   Items{layout},
	}
	return div.Render()
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
