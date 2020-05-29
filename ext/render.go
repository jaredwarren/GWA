package ext

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strings"
)

// Items ...
type Items []Renderer

// Render ...
func (i Items) Render(w io.Writer) error {
	for _, item := range i {
		err := item.Render(w)
		if err != nil {
			return err
		}
	}
	return nil
}

// Attributes ...
type Attributes map[string]template.HTMLAttr

// Styles ...
type Styles map[string]string

// ToAttr ...
func (s Styles) ToAttr() template.HTMLAttr {
	sp := []string{}
	for k, v := range s {
		sp = append(sp, fmt.Sprintf("%s:%s;", k, v))
	}
	return template.HTMLAttr(strings.Join(sp, " "))
}

// Classes ...
type Classes []string

// ToAttr ...
func (c Classes) ToAttr() template.HTMLAttr {
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

	return template.HTMLAttr(strings.Join(npClasses, " "))
}

// Add ... not sure if this is a good idea or not
func (c *Classes) Add(class string) {
	*c = append(*c, class)
}

// Renderer an item that can be rendered
type Renderer interface {
	Render(w io.Writer) error
	GetID() string
}

// Render a template
func renderTemplate(w io.Writer, t string, data interface{}) error {
	funcMap := template.FuncMap{
		"Render": func(item Renderer) template.HTML {
			buf := new(bytes.Buffer)
			err := item.Render(buf)
			if err != nil {
				fmt.Printf("[E] renderTemplate:%s -> %+v\n", err, item)
			}
			return template.HTML(buf.String())
		},
	}
	tpl, err := template.New("base").Funcs(funcMap).ParseFiles(fmt.Sprintf("templates/%s.html", t))
	if err != nil {
		fmt.Printf("[E] %s parse error:%s\n", t, err)
	}
	templates := template.Must(tpl, err)
	return templates.ExecuteTemplate(w, "base", data)
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

func render(w io.Writer, t string, data interface{}) error {
	tpl, err := template.New("base").Funcs(template.FuncMap{
		"Render": func(item Renderer) template.HTML {
			if item == nil {
				return template.HTML("NULL ITEM")
			}
			buf := new(bytes.Buffer)
			err := item.Render(buf)
			if err != nil {
				fmt.Printf("[E] html Render:%s -> %+v\n", err, item)
			}
			return template.HTML(buf.String())
		},
	}).Parse(t)
	if err != nil {
		fmt.Printf("[E] parse error:%s\n", err)
	}

	templates := template.Must(tpl, err)
	err = templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Printf("[E] execute error:%s -\n", err)
	}
	return err
}

// DivContainer gerneric div container
type DivContainer struct {
	ID            string
	Items         Items
	ContainerType string
	Classes       []string
	Styles        map[string]string
	Attributes    map[string]template.HTMLAttr
}

// Render ...
func (d *DivContainer) Render(w io.Writer) error {
	// TODO: replace this with Render Element
	return render(w, `<div id="{{.ID}}" class="{{range $c:= $.Classes}}{{$c}} {{end}}" style="{{range $k, $s:= $.Styles}}{{$k}}:{{$s}}; {{end}}">
			{{range $item := $.Items}}
			{{if $item}}{{Render $item}}{{else}}NULL---{{end}}
			{{end}}</div>`, d)
}

// Element gerneric div container
type Element struct {
	Name       string
	Attributes map[string]template.HTMLAttr
	Items      Items
	Innerhtml  template.HTML
}

// Render ...
func (e *Element) Render(w io.Writer) error {
	name := strings.ToLower(e.Name)
	attrs := ""

	// Some Attributes will produce garbage if not added this way i.e. "type" & "onclick"
	for k, val := range e.Attributes {
		attrs = fmt.Sprintf("%s %s=\"%s\"", attrs, k, val)
	}

	if isSelfClosing(name) {
		return render(w, fmt.Sprintf(`<%s %s>`, name, attrs), e)
	}

	if e.Innerhtml != "" {
		_, err := fmt.Fprintf(w, `<%s %s>%s</%s>`, name, attrs, e.Innerhtml, name)
		return err
	}
	return render(w, fmt.Sprintf(`<%s %s>{{range $item := $.Items}}{{if $item}}{{Render $item}}{{else}}NULL---{{end}}{{end}}</%s>`, name, attrs, name), e)
}

// List of self closing tags
var closing = map[string]bool{
	"area":   true,
	"base":   true,
	"br":     true,
	"col":    true,
	"embed":  true,
	"hr":     true,
	"img":    true,
	"input":  true,
	"link":   true,
	"meta":   true,
	"param":  true,
	"source": true,
	"track":  true,
	"wbr":    true,
}

// check if node name is self closing
func isSelfClosing(name string) bool {
	// name = strings.ToLower(name)
	_, ok := closing[name]
	return ok
}

// GetID ...
func (e *Element) GetID() string {
	return ""
}

// // getClasses ...
// func getClasses(c []string) []string {
// 	// default classes
// 	classess := map[string]bool{}
// 	// copy classes
// 	for _, c := range c {
// 		if _, ok := classess[c]; !ok {
// 			classess[c] = true
// 		}
// 	}
// 	// convert class back to array
// 	npClasses := []string{}
// 	for k := range classess {
// 		npClasses = append(npClasses, k)
// 	}
// 	return npClasses
// }
