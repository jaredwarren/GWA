package gbt

import (
	"fmt"
	"html/template"
	"strings"
)

func NewElement(name string, opts ...Option[any]) *Element {
	fb := &Element{
		Name: name,
	}
	for _, op := range opts {
		op(any(fb))
	}
	return fb
}

// Element gerneric div container
type Element struct {
	ID   string
	Name string
	Attributes
	Classes
	Items
	InnerHTML Stringer
}

func (e *Element) SetInnerHTML(html Stringer) {
	e.InnerHTML = html
	e.Items = nil
}

func (e *Element) SetItems(items Items) {
	e.Items = items
	e.InnerHTML = ""
}

func (e *Element) SetAttributes(a Attributes) {
	e.Attributes = a
}

func (e *Element) SetClasses(c Classes) {
	e.Classes = c
}

// Render ...
func (e *Element) Render() Stringer {
	if e.Name == "" {
		e.Name = "div"
	}
	name := strings.ToLower(e.Name)

	if e.Attributes == nil {
		e.Attributes = Attributes{}
	}

	if e.ID != "" {
		e.Attributes["id"] = e.ID
	}

	if _, ok := e.Attributes["class"]; !ok {
		c := e.Classes.Render()
		if c != "" {
			e.Attributes["class"] = c
		}
	}

	if IsSelfClosing(name) {
		return template.HTML(renderToHTML(fmt.Sprintf(`<%s {{.Attributes.Render}}>`, name), e))
	}

	if e.InnerHTML != "" && e.InnerHTML != nil {
		return template.HTML(renderToHTML(fmt.Sprintf(`<%s {{.Attributes.Render}}>{{.InnerHTML}}</%s>`, name, name), e))
	}

	return template.HTML(renderToHTML(fmt.Sprintf(`<%s {{.Attributes.Render}}>{{.Items.Render}}</%s>`, name, name), e))
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
func IsSelfClosing(name string) bool {
	_, ok := closing[name]
	return ok
}

// Div easier
func NewDiv(opts ...Option[any]) *Element {
	fb := &Element{
		Name: "div",
	}
	for _, op := range opts {
		op(any(fb))
	}
	return fb
}
