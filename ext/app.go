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

// should be interface

// Item ...
type Item struct {
	Title     string
	IconClass string
	Layout    string
	Items     []Renderer // maybe items should be an interface????
	// Tables []Table
}

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
