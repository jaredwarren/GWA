package gbt

import (
	"html/template"
	"strings"
)

type Head struct {
	Title string
	Items
}

type HeadElement Renderer

func (h *Head) Render() Stringer {
	return renderToHTML(`
	<title>{{.Title}}</title>
	{{.Items.Render}}`, h)
}

// Meta <meta> tag
type Meta struct {
	Charset   string
	Name      string
	Content   string
	HttpEquiv string
	Attributes
}

func (m *Meta) Render() Stringer {
	if m.Attributes == nil {
		m.Attributes = Attributes{}
	}
	if m.Charset != "" {
		m.Attributes["charset"] = m.Charset
	}

	if m.Name != "" {
		m.Attributes["name"] = m.Name
	}

	if m.Content != "" {
		m.Attributes["content"] = m.Content
	}

	if m.HttpEquiv != "" {
		m.Attributes["http-equiv"] = m.HttpEquiv
	}
	if len(m.Attributes) == 0 {
		return ""
	}
	return renderToHTML(`<meta {{.Attributes.Render}}>`, m)
}

// Link <link> tag
type Link struct {
	Rel         string
	Type        string
	Href        string
	Integrity   string
	Crossorigin string
	Attributes
}

func (l *Link) Render() Stringer {
	if l.Attributes == nil {
		l.Attributes = Attributes{}
	}
	if l.Href != "" {
		if strings.HasSuffix(l.Href, ".css") {
			l.Rel = "stylesheet"
			l.Type = "text/css"
		}
		l.Attributes["href"] = l.Href
	}

	if l.Rel != "" {
		l.Attributes["rel"] = l.Rel
	}

	if l.Type != "" {
		l.Attributes["type"] = l.Type
	}

	if l.Integrity != "" {
		l.Attributes["integrity"] = l.Integrity
	}

	if l.Crossorigin != "" {
		l.Attributes["crossorigin"] = l.Crossorigin
	}

	if len(l.Attributes) == 0 {
		return ""
	}

	e := &Element{
		Name:       "link",
		Attributes: l.Attributes,
	}
	return e.Render()
}

// Style <style> tag
type Style struct {
	Type string
	Body template.CSS
	Attributes
}

func (s *Style) Render() Stringer {
	if s.Attributes == nil {
		s.Attributes = Attributes{}
	}
	if s.Type != "" {
		s.Attributes["type"] = s.Type
	}

	e := &Element{
		Name:       "style",
		Attributes: s.Attributes,
		InnerHTML:  template.HTML(s.Body),
	}
	return e.Render()
}

// CSSLink plain local css link
type CSSLink string

// Render ...
func (s CSSLink) Render() Stringer {
	e := &Element{
		Name: "link",
		Attributes: Attributes{
			"rel":  "stylesheet",
			"type": "text/css",
			"href": s,
		},
	}
	return e.Render()
}
