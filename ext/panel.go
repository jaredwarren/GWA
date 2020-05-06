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
	// HTML
	if p.HTML != "" {
		p.Items = append(p.Items, &Innerhtml{
			HTML: p.HTML,
		})
	}

	items := layoutItems(p.Items)

	// TODO: parse p.Style and add to styles
	p.Styles = styles

	p.Items = items
	return render("panel", p)
}
func layoutItems(oi []Renderer) []Renderer {
	if len(oi) < 2 {
		return oi
	}

	bodyItems := []Renderer{}
	layout := &Layout{
		Items: []Renderer{},
	}
	var di Renderer
	for _, i := range oi {
		// already find docked item, append rest to body and move on
		if di != nil {
			bodyItems = append(bodyItems, i)
			continue
		}

		// Look for docked item
		// if not dockable add to items, else add to body
		d, ok := i.(Dockable)
		if ok {
			docked := d.GetDocked()
			if docked != "" {
				switch docked {
				case "top":
					layout.Type = "hbox"
					layout.Pack = "start"
					// i goes first
				case "bottom":
					layout.Type = "hbox"
					layout.Pack = "end"
					// i goes last
				case "left":
					layout.Type = "vbox"
					layout.Pack = "start"
					// i goes first
				case "right":
					layout.Type = "vbox"
					layout.Pack = "end"
					// i goes last
				default:
					// what to do
				}
				layout.Align = "start" // should always be start?
				di = i
			} else {
				bodyItems = append(bodyItems, i)
			}
		} else {
			bodyItems = append(bodyItems, i)
		}
	}

	// Nothing to layout
	if di != nil {
		layout.Items = []Renderer{di}

		// Add body items to layout
		if len(bodyItems) > 0 {
			// fmt.Printf("=%d=  ", len(bodyItems))
			// lbi := len(bodyItems)
			// recurse on body
			bi := layoutItems(bodyItems)
			if len(bi) > 0 {
				layout.Items = append(layout.Items, NewBody(bi))
			} // else nothing?
		} else {
			// else what to do????? add a blank one?
			panic("no body")
		}
	} else {
		return oi
		// layout.Items = []Renderer{}
	}

	// add rest
	return []Renderer{layout}
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

// Debug ...
func (p *Panel) Debug() {
	fmt.Print("panel:")
	if len(p.Items) > 0 {
		fmt.Print("\n")
		for _, i := range p.Items {
			fmt.Print("  ")
			i.Debug()
		}
	} else {
		fmt.Printf("  %+v\n", p.HTML)
	}
	fmt.Print("\n")
}
