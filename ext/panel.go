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
	Docked    string // top, bottom, left, right, ''
	Flex      int
	Style     string
	Shadow    bool
	Classes   []string
	Styles    map[string]string
}

// Render ...
func (p *Panel) Render() template.HTML {
	fmt.Println("  render panel")

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

	// TODO: might have to check all items, if docked and then add docked class here?

	// HEADER
	if p.Title != "" {
		if p.Header == nil {
			p.Header = NewHeader(p.Title)
		} else if p.Header.Title == "" {
			p.Header.Title = p.Title
		}
	}

	// ITEMS
	bodyItems := []Renderer{}

	// HTML
	if p.HTML != "" {
		// TODO: append to body!!!!
		bodyItems = append(bodyItems, &Innerhtml{
			HTML: p.HTML,
		})
	}

	items := layoutItems(p.Items)

	if len(bodyItems) > 0 {
		items = append(items, NewBody(bodyItems))
	}

	// TODO: parse p.Style and add to styles
	p.Styles = styles

	p.Items = items
	return render("panel", p)
}
func layoutItems(oi []Renderer) []Renderer {
	items := []Renderer{}
	for _, i := range oi {
		// if not dockable add to items, else add to body
		di, ok := i.(Dockable)
		if ok && di.GetDocked() != "" {
			items = append(items, i)
		} else {
			bodyItems = append(bodyItems, i)
		}
	}
	return items
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
