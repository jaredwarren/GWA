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
			item.Render(buf)
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
				fmt.Printf("[E] html Render:%s\n", err)
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
	Name        string
	SelfClosing bool
	Attributes  map[string]template.HTMLAttr
	Items       Items
}

// Render ...
func (e *Element) Render(w io.Writer) error {
	name := string(e.Name)
	attrs := ""

	// Some Attributes will produce garbage if not added this way i.e. "type" & "onclick"
	for k, val := range e.Attributes {
		attrs = fmt.Sprintf("%s %s=\"%s\"", attrs, k, val)
	}

	if e.SelfClosing {
		return render(w, fmt.Sprintf(`<%s %s>`, name, attrs), e)
	}

	return render(w, fmt.Sprintf(`<%s %s>{{range $item := $.Items}}
			{{if $item}}{{Render $item}}{{else}}NULL---{{end}}
			{{end}}</%s>`, name, attrs, name), e)
}

// GetID ...
func (e *Element) GetID() string {
	return ""
}

func styleToAttr(styles map[string]string) template.HTMLAttr {
	sp := []string{}
	for k, v := range styles {
		sp = append(sp, fmt.Sprintf("%s:%s;", k, v))
	}
	return template.HTMLAttr(strings.Join(sp, " "))
}
