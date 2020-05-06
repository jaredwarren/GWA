package ext

import (
	"bytes"
	"fmt"
	"html/template"
)

// Application ...
type Application struct {
	Name string
	// Schemas map[string]Schema
	// Using   string // seledted schema
	MainView *View
}

// View ...
type View struct {
	// Name   string
	// Tables []Table
}

// TabPanel ...
type TabPanel struct {
	// Name   string
	// Tables []Table
}

// List ...
type List struct {
	Title   string
	Store   *Store
	Columns []*Column
}

// Column ...
type Column struct {
	Text      string
	DataIndex string
	Width     int
}

// Layout ...
type Layout struct {
	Type  string // absolute, accordion, border, card, tab, hbox, vbox
	Pack  string // start, end, center, space-between, space-arround, justify
	Align string // start, end, center, stretch
	Items []Renderer
}

// // Item ...
// type Item struct {
// 	Title     string
// 	IconClass string
// 	Layout    string
// 	Items     []Renderer
// 	// Tables []Table
// }

// Store ...
type Store struct {
}

// Renderer ...
type Renderer interface {
	Render() template.HTML
}

func render(t string, data interface{}) template.HTML {
	funcMap := template.FuncMap{
		"Render": func(item Renderer) template.HTML {
			return item.Render()
		},
	}

	buf := new(bytes.Buffer)
	tpl, err := template.New("base").Funcs(funcMap).ParseFiles(fmt.Sprintf("templates/%s.html", t))
	if err != nil {
		fmt.Printf("[E] %s parse error:%s\n", t, err)
	}
	templates := template.Must(tpl, err)
	err = templates.ExecuteTemplate(buf, "base", data)
	if err != nil {
		fmt.Printf("[E] %s exec error:%s\n", t, err)
	}
	return template.HTML(buf.String())
}

// Dockable ...
type Dockable interface {
	GetDocked() string
}
