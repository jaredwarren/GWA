package ext

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

var (
	navID = 0
)

// NewNav ...
func NewNav(title template.HTML) *Nav {
	return &Nav{
		ID:     nextNavID(),
		Title:  title,
		Border: template.CSS("1px solid lightgrey"),
		Docked: "top",
	}
}

// Nav ...
// TODO: I don't need all of this crap for container
type Nav struct {
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
func (n *Nav) Render(w io.Writer) error {
	if n.ID == "" {
		n.ID = nextNavID()
	}
	// default classes
	classess := map[string]bool{
		"navbar":           true,
		"navbar-expand-sm": true,
		"navbar-dark":      true,
		"bg-dar":           true,
	}
	// copy classes
	for _, c := range n.Classes {
		if _, ok := classess[c]; !ok {
			classess[c] = true
		}
	}

	if n.Shadow {
		classess["x-shadow"] = true
	}

	// copy styles
	styles := map[string]string{}
	if len(n.Styles) > 0 {
		styles = n.Styles
	}

	// append new styles based on p's properties
	if n.Width != 0 && n.Docked != "top" && n.Docked != "bottom" {
		styles["width"] = fmt.Sprintf("%dpx", n.Width)
		classess["x-widthed"] = true
	}
	if n.Height != 0 && n.Docked != "left" && n.Docked != "right" {
		styles["height"] = fmt.Sprintf("%dpx", n.Height)
		classess["x-heighted"] = true
	}
	if n.Border != "" {
		styles["border"] = string(n.Border)
		classess["x-managed-border"] = true
	}

	// convert class back to array
	npClasses := []string{}
	for k := range classess {
		npClasses = append(npClasses, k)
	}

	items := Items{}

	// Title
	if n.Title != "" {
		title := &Innerhtml{
			HTML: n.Title,
			Styles: map[string]string{
				"flex":       "1",
				"text-align": "center",
			},
		}
		items = append(items, title)
	}

	// append rest of items
	if len(n.Items) > 0 {
		items = append(items, n.Items...)
	}

	// update parent
	for _, i := range items {
		c, ok := i.(Child)
		if ok {
			c.SetParent(n)
		}
	}

	// Attributes
	attrs := map[string]template.HTMLAttr{
		"id":    template.HTMLAttr(n.ID),
		"class": template.HTMLAttr(strings.Join(npClasses, " ")),
	}
	if len(styles) > 0 {
		attrs["style"] = styleToAttr(styles)
	}

	navEl := &Element{
		Name:       "nav",
		Attributes: attrs,
		Items:      items,
	}
	return navEl.Render(w)
}

// GetID ...
func (n *Nav) GetID() string {
	return n.ID
}

// GetDocked ...
func (n *Nav) GetDocked() string {
	return n.Docked
}

// SetStyle ...
func (n *Nav) SetStyle(key, value string) {
	if n.Styles == nil {
		n.Styles = map[string]string{}
	}
	n.Styles[key] = value
}

func nextNavID() string {
	id := fmt.Sprintf("nav-%d", navID)
	navID++
	return id
}
