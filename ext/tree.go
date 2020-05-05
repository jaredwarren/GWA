package ext

import (
	"bytes"
	"fmt"
	"html/template"
	"net/url"
)

// Tree ...
type Tree struct {
	ID    string
	Store *TreeStore
	// Text      template.HTML
	// Handler   template.JS
	// UI        string // TODO
	// IconClass string
}

// Render ...
func (t *Tree) Render() template.HTML {

	t.ID = fmt.Sprintf("%d", buttonID)
	buttonID++

	buf := new(bytes.Buffer)
	templates := template.Must(template.ParseFiles("templates/button.html"))
	templates.ExecuteTemplate(buf, "base", t)
	return template.HTML(buf.String())
}

// TreeStore ...
type TreeStore struct {
	ID         string
	AutoLoad   bool
	Root       *Node  // TODO: type
	Proxy      *Proxy // TODO: type
	Sorters    string // TODO type
	FolderSort bool   // TODO:
}

// Proxy ...
type Proxy struct {
	Type   string
	URL    *url.URL
	reader *Reader
}

// Reader ...
type Reader struct {
	Type         string
	RootProperty string
}

// Node ...
type Node struct {
	Text      string
	ID        string
	Expanded  bool
	leaf      bool
	IconClass string
	Children  []*Node
}
