package gbt

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
	ID          string        `json:"id,omitempty"`
	XType       string        `json:"xtype,omitempty"`
	Title       template.HTML `json:"title,omitempty"`
	IconClass   string        `json:"iconClass,omitempty"`
	Layout      string        `json:"layout,omitempty"`
	HTML        template.HTML `json:"html,omitempty"`
	Width       int           `json:"width,omitempty"`
	Height      int           `json:"height,omitempty"`
	Resizable   bool          `json:"resizable,omitempty"`
	Resize      string        `json:"resize,omitempty"`
	Collapsible bool          `json:"collapsible,omitempty"`
	Collapsed   bool          `json:"collapsed,omitempty"`
	Items       Items         `json:"items,omitempty"`
	Header      *Header       `json:"header,omitempty"`
	// Nav         *Nav          `json:"nav,omitempty"`
	Body        Renderer     `json:"body,omitempty"`
	Border      template.CSS `json:"border,omitempty"`
	Docked      string       `json:"docked,omitempty"` // top, bottom, left, right, ''
	Flex        string       `json:"flex,omitempty"`
	Shadow      bool         `json:"shadow,omitempty"`
	Closable    bool         `json:"closable,omitempty"`
	Collapsable bool         `json:"collapsable,omitempty"`
	Classes     Classes      `json:"classes,omitempty"`
	Styles      Styles       `json:"styles,omitempty"`
	Controller  *Controller  `json:"-"`
	Parent      Renderer     `json:"-"`
	// RenderTo  string // type???
}

// Render ...
func (p *Panel) Render() Stringer {
	ddiv := &DivContainer{
		ID:      p.ID,
		Classes: p.Classes,
		// Styles:  styles,
		Items: p.Items,
	}
	return ddiv.Render()
	// default classes
	p.Classes.Add("x-panel")
	if p.Shadow {
		p.Classes.Add("x-shadow")
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
		p.Classes.Add("x-widthed")
	}
	// what if I want height to be 0px?
	if p.Height != 0 && p.Docked != "left" && p.Docked != "right" {
		styles["height"] = fmt.Sprintf("%dpx", p.Height)
		p.Classes.Add("x-heighted")
	}
	if p.Border != "" {
		styles["border"] = string(p.Border)
		p.Classes.Add("x-managed-border")
	}

	if p.Layout == "hbox" {
		styles["flex-direction"] = "row"
		// p.Classes.Add("x-vbox")
	} else {
		// p.Classes.Add("x-hbox")
		styles["flex-direction"] = "column"
	}

	if p.Flex != "" {
		styles["flex"] = p.Flex
	}

	if p.Resizable {
		styles["resize"] = "both"
	} else if p.Resize != "" {
		styles["resize"] = p.Resize
	}

	// ITEMS
	items := Items{}

	// NAV
	// if p.Nav != nil {
	// 	if p.Title != "" {
	// 		p.Nav.Title = p.Title
	// 	}
	// 	if p.Nav.Docked == "" {
	// 		p.Nav.Docked = "top"
	// 	}
	// 	// append nav as docked item[0]
	// 	items = append(items, p.Nav)
	// }

	// HEADER
	var header *Header
	if p.Title != "" {
		// TODO: if nav has title don't add to header???
		if p.Header == nil {
			header = NewHeader(p.Title)
		} else if p.Header.Title == "" {
			header = p.Header
			header.Title = p.Title
			header.Docked = "top"
		} // else header is all set, ignore Title attribute
	} else if p.Header != nil {
		header = p.Header
	} // else assume no header

	// Collapsed
	if p.Collapsible || p.Collapsed {
		if header == nil {
			header = NewHeader("&nbsp;")
		}
		header.Items = append(header.Items, &Button{
			Handler: "collapsePanel",
			// IconClass: "far fa-angle-left",
		})

		collapsedPanel := &Panel{
			Docked: "right",
			Items: Items{&Button{
				Handler: "expandPanel",
				// IconClass: "far fa-angle-right",
			}},
		}

		if p.Collapsed {
			// collapsedPanel.Styles = Styles{
			// 	"display": "flex",
			// }
			// styles["display"] = "none"

			p.Classes.Add("collapsed")

		} else {
			p.Classes.Add("expanded")
			collapsedPanel.Styles = Styles{
				"display": "none",
			}
		}

		// TODO: add title vertically?
		items = append(items, collapsedPanel)
	}

	// append header as docked item[0] or item[1] if nav
	if header != nil {
		if header.Docked == "" {
			header.Docked = "top"
		}
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

	for _, i := range items {
		c, ok := i.(Child)
		if ok {
			c.SetParent(p)
		}
	}

	div := &DivContainer{
		ID:      p.ID,
		Classes: p.Classes,
		Styles:  styles,
		Items:   LayoutItems(items),
	}
	return div.Render()
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

// SetStyle ...
func (p *Panel) SetStyle(key, value string) {
	if p.Styles == nil {
		p.Styles = map[string]string{}
	}
	p.Styles[key] = value
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

func buildPanel(i interface{}) *Panel {
	ii := i.(map[string]interface{})

	p := &Panel{}

	if title, ok := ii["title"]; ok {
		p.Title = template.HTML(title.(string))
	}

	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if IconClass, ok := ii["iconClass"]; ok {
		p.IconClass = IconClass.(string)
	}

	if Layout, ok := ii["layout"]; ok {
		p.Layout = Layout.(string)
	}

	if HTML, ok := ii["html"]; ok {
		p.HTML = template.HTML(HTML.(string))
	}

	if Width, ok := ii["width"]; ok {
		p.Width = Width.(int)
	}

	if Height, ok := ii["height"]; ok {
		p.Height = Height.(int)
	}

	if header, ok := ii["header"]; ok {
		p.Header = addChild(header).(*Header)
	}

	// if body, ok := ii["body"]; ok {
	// 	p.Body = addChild(body).(*DivContainer)
	// }

	if border, ok := ii["border"]; ok {
		p.Border = template.CSS(border.(string))
	}

	if docked, ok := ii["docked"]; ok {
		p.Docked = docked.(string)
	}
	if flex, ok := ii["flex"]; ok {
		p.Flex = flex.(string)
	}

	if shadow, ok := ii["shadow"]; ok {
		p.Shadow = shadow.(bool)
	}

	if c, ok := ii["classes"]; ok {
		jclass := c.([]interface{})
		classes := make([]string, len(jclass))
		for i, cl := range jclass {
			classes[i] = cl.(string)
		}
		p.Classes = classes
	}

	if s, ok := ii["styles"]; ok {
		jclass := s.(map[string]interface{})
		styles := map[string]string{}
		for i, cl := range jclass {
			styles[i] = cl.(string)
		}
		p.Styles = styles
	}

	items := []Renderer{}
	if ii, ok := ii["items"]; ok {
		is := ii.([]interface{})
		for _, i := range is {
			item := addChild(i)
			items = append(items, item)
		}
	}

	p.Items = items

	return p
}
