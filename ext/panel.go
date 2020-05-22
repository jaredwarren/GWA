package ext

import (
	"fmt"
	"html/template"
	"io"
	"unicode"
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
	ID          string            `json:"id,omitempty"`
	XType       string            `json:"xtype,omitempty"`
	Title       template.HTML     `json:"title,omitempty"`
	IconClass   string            `json:"iconClass,omitempty"`
	Layout      string            `json:"layout,omitempty"`
	HTML        template.HTML     `json:"html,omitempty"`
	Width       int               `json:"width,omitempty"`
	Height      int               `json:"height,omitempty"`
	Items       Items             `json:"items,omitempty"`
	Header      *Header           `json:"header,omitempty"`
	Body        *Body             `json:"body,omitempty"`
	Border      template.CSS      `json:"border,omitempty"`
	Docked      string            `json:"docked,omitempty"` // top, bottom, left, right, ''
	Flex        int               `json:"flex,omitempty"`
	Shadow      bool              `json:"shadow,omitempty"`
	Closable    bool              `json:"closable,omitempty"`
	Collapsable bool              `json:"collapsable,omitempty"`
	Classes     []string          `json:"classes,omitempty"`
	Styles      map[string]string `json:"styles,omitempty"`
	Controller  *Controller       `json:"-"`
	Parent      Renderer          `json:"-"`
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
	} else if p.Header != nil {
		header = p.Header
	} // else assume no header
	// append header as docked item[0]
	if header != nil {
		if header.Docked == "" {
			header.Docked = "top"
		}

		// TODO: if

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

// // MarshalJSON ...
// func (p *Panel) MarshalJSON() ([]byte, error) {
// 	result := map[string]interface{}{}
// 	e := reflect.ValueOf(p).Elem()
// 	for i := 0; i < e.NumField(); i++ {
// 		// fmt.Println(varName, "->", , " ---- ")
// 		varName := ""

// 		omit := true

// 		jsonTag := e.Type().Field(i).Tag.Get("json")
// 		if jsonTag != "" {
// 			tags := strings.Split(jsonTag, ",")
// 			varName = tags[0]
// 			if len(tags) > 1 && tags[1] == "omitempty" {

// 			}
// 		} else {
// 			varName := lowerInitial(e.Type().Field(i).Name)
// 		}

// 		if e.Field(i).CanInterface() && varName != "" && omit {
// 			result[varName] = e.Field(i).Interface()
// 		}
// 	}
// 	return json.Marshal(result)
// }

func lowerInitial(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
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

	if body, ok := ii["body"]; ok {
		p.Body = addChild(body).(*Body)
	}

	if border, ok := ii["border"]; ok {
		p.Border = template.CSS(border.(string))
	}

	if docked, ok := ii["docked"]; ok {
		p.Docked = docked.(string)
	}
	if flex, ok := ii["flex"]; ok {
		p.Flex = flex.(int)
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
