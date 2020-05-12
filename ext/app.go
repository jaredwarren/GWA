package ext

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"reflect"
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

// Debug ...
func Debug(p Renderer) {
	d(p, 0)
}

func d(p Renderer, depth int) {
	pd(depth)
	typeof := reflect.TypeOf(p).String()
	fmt.Print("| ", typeof)

	switch typeof {
	case "*ext.Panel":
		fmt.Println("  html:", p.(*Panel).HTML)
		pd(depth)
		fmt.Println("  style:", p.(*Panel).Styles)
		for _, i := range p.(*Panel).Items {
			d(i, depth+1)
		}
	case "*ext.Innerhtml":
		fmt.Println("  html:", p.(*Innerhtml).HTML)
	case "*ext.Layout":
		fmt.Println(":", p.(*Layout).Type)
		for _, i := range p.(*Layout).Items {
			d(i, depth+1)
		}
	case "*ext.Body":
		fmt.Println("")
		for _, i := range p.(*Body).Items {
			d(i, depth+1)
		}
	case "*ext.Header":
		fmt.Println("::", p.(*Header).Title)
	default:
		fmt.Println("")
	}
}

func pd(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
}
