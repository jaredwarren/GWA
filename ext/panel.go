package ext

import (
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
	ID        string // how to auto generate
	Title     string
	IconClass string
	Layout    string
	HTML      template.HTML
	Width     int // float?
	Height    int // float?
	Items     Items
	RenderTo  string // type???
	Header    *Header
	Body      *Body
	Border    template.CSS
	Docked    string // top, bottom, left, right, ''
	Flex      int
	Shadow    bool
	Classes   []string
	Styles    map[string]string
}

// Render ...
func (p *Panel) Render(w io.Writer) error {
	if p.ID == "" {
		p.ID = nextPanelID()
	}

	// default classes
	classess := map[string]bool{
		"x-panel":               true,
		"x-container":           true,
		"x-component":           true,
		"x-noborder-trbl":       true,
		"x-header-position-top": true,
		"x-root":                true,
	}
	// add classses from p, no duplicates
	for _, c := range p.Classes {
		if _, ok := classess[c]; !ok {
			classess[c] = true
		}
	}
	if p.Shadow {
		classess["x-shadow"] = true
	}
	// TODO: add other classess?

	// copy styles from og
	styles := map[string]string{}
	if len(p.Styles) > 0 {
		styles = p.Styles
	}

	// append new styles based on p's properties
	// TODO: if docked, ignore width or height
	if p.Width != 0 { // what if I want width to be 0px?
		styles["width"] = fmt.Sprintf("%dpx", p.Width)
		classess["x-widthed"] = true
	}
	if p.Height != 0 { // what if I want height to be 0px?
		styles["height"] = fmt.Sprintf("%dpx", p.Height)
		classess["x-heighted"] = true
	}
	if p.Border != "" {
		styles["border"] = string(p.Border)
		classess["x-managed-border"] = true
	}

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
	// TODO: append header as docked item[0]
	if header != nil {
		// np.Header = header
		items = append(items, header)
	}

	if len(p.Items) > 0 {
		items = append(items, p.Items...)
	}

	// ITEMS
	// HTML
	if p.HTML != "" {
		items = append(items, &Innerhtml{
			HTML: p.HTML,
		})
	}

	// TODO: if panel has "layout" set that up here
	// // This layout should only apply to non-docked items!

	// Ad layout to all items
	div := &DivContainer{
		ID:      fmt.Sprintf("panel-%s", p.ID),
		Classes: npClasses,
		Styles:  styles,
		Items:   LayoutItems(items),
	}
	return renderDiv(w, "body", div)
}

func nextPanelID() string {
	id := fmt.Sprintf("%d", panelID)
	panelID++
	return id
}

// GetDocked ...
func (p *Panel) GetDocked() string {
	return p.Docked
}
