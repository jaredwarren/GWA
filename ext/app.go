package ext

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
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
	Render(w io.Writer) error
	Build() (Renderer, error)
	Debug()
}

func render(w io.Writer, t string, data interface{}) error {
	funcMap := template.FuncMap{
		"Render": func(item Renderer) template.HTML {
			buf := new(bytes.Buffer)
			item.Render(w)
			return template.HTML(buf.String())
		},
	}

	// buf := new(bytes.Buffer)
	tpl, err := template.New("base").Funcs(funcMap).ParseFiles(fmt.Sprintf("templates/%s.html", t))
	if err != nil {
		fmt.Printf("[E] %s parse error:%s\n", t, err)
	}
	templates := template.Must(tpl, err)
	return templates.ExecuteTemplate(w, "base", data)
}

// Dockable ...
type Dockable interface {
	GetDocked() string
}
