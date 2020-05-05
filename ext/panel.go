package ext

import (
	"fmt"
	"html/template"
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
	Items     []Renderer
	RenderTo  string // type???
	Header    *Header
	Body      *Body
	Border    template.CSS
	Docked    string // top, bottom, left, right
	Flex      int
	Style     string
	Shadow    bool
	Classes   []string
	Styles    map[string]string
}

// Render ...
func (p *Panel) Render() template.HTML {
	fmt.Println("  render panel")
	if p.Header == nil {
		p.Header = NewHeader(p.Title)
	}

	if p.ID == "" {
		p.ID = nextPanelID()
	}

	// default classes
	p.Classes = []string{
		"x-panel",
		"x-container",
		"x-component",
		"x-noborder-trbl",
		"x-header-position-top",
		"x-root",
	}

	if p.Shadow {
		p.Classes = append(p.Classes, "x-shadow")
	}

	styles := map[string]string{}
	if p.Width != 0 { // what if I want width to be 0px?
		styles["width"] = fmt.Sprintf("%dpx", p.Width)
		p.Classes = append(p.Classes, "x-widthed")
	}
	if p.Height != 0 { // what if I want height to be 0px?
		styles["height"] = fmt.Sprintf("%dpx", p.Height)
		p.Classes = append(p.Classes, "x-heighted")
	}
	if p.Border != "" {
		styles["border"] = string(p.Border)
		p.Classes = append(p.Classes, "x-managed-border")
	}

	// TODO: might have to check all items, if docked?

	// HTML
	if p.HTML != "" {
		p.Items = append(p.Items, &Container{
			HTML: p.HTML,
		})
	}

	// ITEMS
	if len(p.Items) == 0 {
		// debug add dummy html
		html := p.HTML
		if html == "" {
			html = ":("
		}
		p.Items = []Renderer{
			&Container{
				HTML: html,
			},
		}
	}

	// BODY
	if p.Body == nil {
		p.Body = NewBody(p.Items)
	}

	// TODO: parse p.Style and add to styles
	p.Styles = styles

	// body
	// x-panel-body x-body-wrap-el x-panel-body-wrap-el x-container-body-wrap-el x-component-body-wrap-el
	// x-body-wrap-el x-panel-body-wrap-el x-container-body-wrap-el x-component-body-wrap-el

	return render("panel", p)
}

// Layout ...
type Layout struct {
	Type  string // absolute, accordion, border, card, tab, hbox, vbox
	Pack  string // start, end, center, space-between, space-arround, justify
	Align string // start, end, center, stretch
}

func nextPanelID() string {
	id := fmt.Sprintf("%d", panelID)
	panelID++
	return id
}
