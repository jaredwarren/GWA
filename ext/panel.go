package ext

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
)

var (
	panelID = 0
)

// NewPanel ...
func NewPanel() *Panel {
	return &Panel{
		ID:     nextPanelID(),
		Width:  300,
		Height: 200,
		Border: template.CSS("1px solid lightgrey"),
	}

}

// Panel ...
type Panel struct {
	ID         string // how to auto generate
	Title      string
	IconClass  string
	Layout     string
	HTML       template.HTML
	Width      int // float?
	Height     int // float?
	Items      Items
	Header     *Header
	Body       *Body
	Border     template.CSS
	Docked     string // top, bottom, left, right, ''
	Flex       int
	Shadow     bool
	Classes    []string
	Styles     map[string]string
	Controller *Controller
	Parent     Renderer
	// RenderTo  string // type???
}

// Render ...
func (p *Panel) Render(w io.Writer) error {
	if p.ID == "" {
		p.ID = nextPanelID()
	}

	// default classes
	classess := map[string]bool{
		"x-panel": true,
	}
	// copy classes
	for _, c := range p.Classes {
		if _, ok := classess[c]; !ok {
			classess[c] = true
		}
	}
	if p.Shadow {
		classess["x-shadow"] = true
	}

	// copy styles
	styles := map[string]string{}
	if len(p.Styles) > 0 {
		styles = p.Styles
	}

	// append new styles based on p's properties
	// what if I want width to be 0px?
	if p.Width != 0 && p.Docked != "top" && p.Docked != "bottom" {
		styles["width"] = fmt.Sprintf("%dpx", p.Width)
		classess["x-widthed"] = true
	}
	// what if I want height to be 0px?
	if p.Height != 0 && p.Docked != "left" && p.Docked != "right" {
		styles["height"] = fmt.Sprintf("%dpx", p.Height)
		classess["x-heighted"] = true
	}
	if p.Border != "" {
		styles["border"] = string(p.Border)
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
	var header *Header
	if p.Title != "" {
		if p.Header == nil {
			header = NewHeader(p.Title)
		} else if p.Header.Title == "" {
			header = p.Header
			header.Title = p.Title
			header.Docked = "top"
		} // else header is all set, ignore Title attribute
	}
	// append header as docked item[0]
	if header != nil {
		items = append(items, header)
	}

	// CONTROLLER
	// TODO: controllers don't work kere yet
	if p.Controller != nil {
		items = append(items, p.Controller)
	}

	// append rest of items
	if len(p.Items) > 0 {
		items = append(items, p.Items...)
	}

	// HTML
	if p.HTML != "" {
		items = append(items, &Innerhtml{
			HTML: p.HTML,
		})
	}

	// TODO: if panel has "layout" set that up here
	// // This layout should only apply to non-docked items!

	for _, i := range items {
		c, ok := i.(Child)
		if ok {
			c.SetParent(p)
		}
	}

	div := &DivContainer{
		ID:      p.ID,
		Classes: npClasses,
		Styles:  styles,
		Items:   LayoutItems(items),
	}
	return renderDiv(w, div)
}

func nextPanelID() string {
	id := fmt.Sprintf("panel-%d", panelID)
	panelID++
	return id
}

// GetDocked ...
func (p *Panel) GetDocked() string {
	return p.Docked
}

// SetParent ...
func (p *Panel) SetParent(parent Renderer) {
	p.Parent = parent
}

// GetChildren ...
func (p *Panel) GetChildren() Items {
	return p.Items
}

// GetID ...
func (p *Panel) GetID() string {
	return p.ID
}

// MarshalJSON ...
func (p *Panel) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		XType     string            `json:"xtype"`
		ID        string            `json:"id,omitempty"`
		Title     string            `json:"title,omitempty"`
		IconClass string            `json:"iconClass,omitempty"`
		Layout    string            `json:"layout,omitempty"`
		HTML      template.HTML     `json:"html,omitempty"`
		Width     int               `json:"width,omitempty"`
		Height    int               `json:"height,omitempty"`
		Items     Items             `json:"items,omitempty"`
		Header    *Header           `json:"header,omitempty"`
		Body      *Body             `json:"body,omitempty"`
		Border    template.CSS      `json:"border,omitempty"`
		Docked    string            `json:"docked,omitempty"`
		Flex      int               `json:"flex,omitempty"`
		Shadow    bool              `json:"shadow,omitempty"`
		Classes   []string          `json:"classes,omitempty"`
		Styles    map[string]string `json:"styles,omitempty"`
	}{
		XType:     "panel",
		ID:        p.ID,
		Title:     p.Title,
		IconClass: p.IconClass,
		Layout:    p.Layout,
		HTML:      p.HTML,
		Width:     p.Width,
		Height:    p.Height,
		Border:    p.Border,
		Docked:    p.Docked,
		Flex:      p.Flex,
		Shadow:    p.Shadow,
		Classes:   p.Classes,
		Styles:    p.Styles,
		Header:    p.Header,
		Body:      p.Body,
		Items:     p.Items,
	})
}
