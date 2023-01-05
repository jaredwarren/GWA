package ext

import (
	"bytes"
	"fmt"
	"html/template"
)

/*
* This stuff needs to be moved to another file eventually
 */
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

// Store ...
type Store struct {
}

type Image struct {
	Src     string
	Alt     string
	Width   string
	Height  string
	Classes []string
}

func (i *Image) Render() Stringer {
	if i.Src == "" {
		return ""
	}

	t := template.New("nav")
	t, _ = t.Parse(`<img 
	src="{{.Src}}" 
	{{if ne .Alt ""}}alt="{{.Alt}}"{{end}} 
	{{if ne .Width ""}}width="{{.Width}}"{{end}} 
	{{if ne .Height ""}}height="{{.Height}}"{{end}} 
	class="{{range $c := .Classes}}{{$c}} {{end}}">`)

	buf := new(bytes.Buffer)
	err := t.Execute(buf, i)
	if err != nil {
		fmt.Println("[E] render head:", err)
	}

	return template.HTML(buf.String())
}
