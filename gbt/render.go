package gbt

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

// Items ...
type Items []Renderer

// Render ...
func (i Items) Render() Stringer {
	return renderToHTML(`{{range $item := .}}
	{{if $item}}
	{{$item.Render}}
	{{end}}
	{{end}}`, i)
}

/**
* Attributes
 */

// Attributes ...
type Attributes map[string]Stringer

func (a Attributes) Render() Stringer {
	al := ""
	for k, v := range a {
		if v == "" {
			al = al + fmt.Sprintf(` %s`, k)
		} else {
			al = al + fmt.Sprintf(` %s="%s"`, k, v)
		}
	}
	return template.HTMLAttr(strings.TrimSpace(al))
}

// Styles ...
type Styles map[string]string

// ToAttr ...
func (s Styles) ToAttr() string {
	sp := []string{}
	for k, v := range s {
		sp = append(sp, fmt.Sprintf("%s:%s;", k, v))
	}
	return strings.Join(sp, " ")
}

/**
* Classes
 */

// Classes ...
type Classes []string

// ToAttr ...
func (c Classes) ToAttr() string {
	classess := map[string]bool{}
	// copy classes to map to prevent duplicates
	for _, c := range c {
		if _, ok := classess[c]; !ok {
			classess[c] = true
		}
	}
	// convert class back to array
	npClasses := []string{}
	for k := range classess {
		npClasses = append(npClasses, k)
	}

	return strings.Join(npClasses, " ")
}

// Add ... not sure if this is a good idea or not
func (c *Classes) Add(class string) {
	*c = append(*c, class)
}

// Len ...
func (c *Classes) Len() int {
	return len(*c)
}

func (c Classes) Render() Stringer {
	return template.HTMLAttr(strings.Join(c, " "))
}

/**
* Events
* https://developer.mozilla.org/en-US/docs/Web/Events
 */

// Events ...
type Events map[string]*Event

// Event ...
type Event struct {
	Name      string            `json:"name,omitempty"`
	Handler   template.HTMLAttr `json:"handler,omitempty"`
	HandlerFn func()            `json:"handlerfn,omitempty"`
}

// ToAttr prepend 'on' to event name e.g. "click" -> "onclick"
func (e Events) ToAttr() map[string]template.HTMLAttr {
	o := map[string]template.HTMLAttr{}
	for n, v := range e {
		nn := fmt.Sprintf("on%s", n)
		o[nn] = v.Handler
	}
	return o
}

/**
* Data
* https://developer.mozilla.org/en-US/docs/Learn/HTML/Howto/Use_data_attributes
 */

// Data ...
type Data map[string]template.HTMLAttr

// ToAttr prepend 'data-' to name
func (d Data) ToAttr() map[string]template.HTMLAttr {
	o := map[string]template.HTMLAttr{}
	for n, v := range d {
		nn := fmt.Sprintf("data-%s", n)
		o[nn] = v
	}
	return o
}

/**
*
 */

// Allow string, template.HTML, etc
type Stringer any

// Renderer an item that can be rendered
type Renderer interface {
	Render() Stringer
}

func renderToHTML(tmp string, data any) template.HTML {
	t := template.New("")
	t, err := t.Parse(tmp)
	if err != nil {
		fmt.Println("[E] parse error:", err)
		return template.HTML(err.Error())
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		fmt.Println("[E] execute error:", err)
		fmt.Println(tmp)
		fmt.Printf("%+v\n", data)
		return template.HTML(err.Error())
	}

	return template.HTML(buf.String())
}

// LayoutItems ...
func LayoutItems(oi Items) Items {
	// if there's only one item there's nothing to layout
	if len(oi) < 2 {
		return oi
	}

	bodyItems := Items{}
	layout := &Layout{
		Items: Items{},
	}
	var di Dockable
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
				di = i.(Dockable)
				switch docked {
				case "top":
					layout.Type = "hbox"
					di.SetStyle("width", "100%")
				case "bottom":
					layout.Type = "hbox"
					layout.Reverse = true
					di.SetStyle("width", "100%")
				case "left":
					layout.Type = "vbox"
					di.SetStyle("height", "100%")
				case "right":
					layout.Type = "vbox"
					layout.Reverse = true
					di.SetStyle("height", "100%")
				default:
					// what to do
				}
				layout.Align = "start" // should always be start?
			} else {
				bodyItems = append(bodyItems, i)
			}
		} else {
			bodyItems = append(bodyItems, i)
		}
	}

	// Nothing to layout
	if di != nil {
		i := di.(Renderer)
		layout.Items = Items{i}
		// Add body items to layout
		if len(bodyItems) > 0 {
			layout.Items = append(layout.Items, NewBody(bodyItems))
		} // else what to do????? add a blank one?
	} else {
		return oi
	}

	// add rest
	return Items{layout}
}

// NewBody ...
func NewBody(items Items) *DivContainer {
	return &DivContainer{
		Classes: []string{
			"x-panel-body",
			"x-body-wrap-el",
			"x-panel-body-wrap-el",
			"x-container-body-wrap-el",
			"x-component-body-wrap-el",
		},
		Items: LayoutItems(items),
	}
}

// DivContainer gerneric div container
type DivContainer struct {
	ID            string
	Items         Items
	ContainerType string
	Classes       Classes
	Styles        Styles
	Attributes    Attributes
}

// Render ...
func (d *DivContainer) Render() Stringer {
	el := &Element{
		Name: "div",
		Attributes: Attributes{
			"class": d.Classes.ToAttr(),
			"style": d.Styles.ToAttr(),
		},
		Items: d.Items,
	}
	return el.Render()
}

// GetID ...
func (d *DivContainer) GetID() string {
	return d.ID
}

/**
*
 */

// Spacer ...
type Spacer struct {
	Type string // vertical or horizontal, TODO: figure out how to guess, based on parent orintation
	// TODO: add other style/class stuff
}

// Render ...
func (s *Spacer) Render() Stringer {
	sp := &Element{
		Name: "span",
		Attributes: Attributes{
			"style": "border-left:1px solid darkgray; margin: 0 4px;",
		},
		InnerHTML: "&nbsp;",
	}
	return sp.Render()
}

// GetID ...
func (s *Spacer) GetID() string {
	return ""
}
