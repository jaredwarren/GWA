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
	Items     []Renderer
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
	fmt.Println("  render panel")

	np, err := p.Build()
	if err != nil {
		fmt.Println("[E] ", err)
	}

	return render(w, "panel", np)
}

// Build copys info to a new panel
func (p *Panel) Build() (Renderer, error) {
	np := &Panel{
		Docked: p.Docked,
	}
	if p.ID != "" {
		np.ID = p.ID
	} else {
		np.ID = nextPanelID()
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
	np.Styles = styles

	np.Classes = []string{}
	for k := range classess {
		np.Classes = append(np.Classes, k)
	}

	// ITEMS
	items := []Renderer{}

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

	// Build everything
	//*/

	is := make([]Renderer, len(items))
	for i, item := range items {
		ii, _ := item.Build()
		// fmt.Printf("... %+v --> %s\n", ii, reflect.TypeOf(ii))
		is[i] = ii
	}
	// np.Items = is
	np.Items = LayoutItems(is)
	/*/
	for j, i := range items {
		fmt.Printf(" (%d)======> %+v\n", j, i)
	}
	np.Items = items
	np.Items = LayoutItems(items)
	//*/

	return np, nil
}

// LayoutItems ...
func LayoutItems(oi []Renderer) []Renderer {
	// if there's only one item there's nothing to layout
	if len(oi) < 2 {
		return oi
	}

	bodyItems := []Renderer{}
	layout := &Layout{
		Items: []Renderer{},
	}
	var di Renderer
	for _, i := range oi {
		// i, _ = i.Build()
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
			bi := LayoutItems(bodyItems)
			if len(bi) > 0 {
				layout.Items = append(layout.Items, NewBody(bi))
			} // else nothing?
		} // else what to do????? add a blank one?
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
