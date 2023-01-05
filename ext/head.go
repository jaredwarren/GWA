package ext

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

type Head struct {
	Title string
	Items Items
}

type HeadElement Renderer

func (h *Head) RenderString() string {
	return "TODO: Application"
}

func (h *Head) Render(w io.Writer) error {
	if h.Title != "" {
		_, err := fmt.Fprintf(w, `<title>%s</title>`, h.Title)
		if err != nil {
			return err
		}
	}
	// TODO: wrap in head once i clean up base.html
	for _, hi := range h.Items {
		err := hi.Render(w)
		if err != nil {
			return err
		}
	}
	return nil
}

// Meta <meta> tag
type Meta struct {
	Charset   string
	Name      string
	Content   string
	HttpEquiv string
}

func (m *Meta) RenderString() string {
	return "TODO: Application"
}
func (m *Meta) Render(w io.Writer) error {
	attrs := ""
	if m.Charset != "" {
		attrs = attrs + fmt.Sprintf(` charset="%s"`, m.Charset)
	}

	if m.Name != "" {
		attrs = attrs + fmt.Sprintf(` name="%s"`, m.Name)
	}

	if m.Content != "" {
		attrs = attrs + fmt.Sprintf(` content="%s"`, m.Content)
	}

	if m.HttpEquiv != "" {
		attrs = attrs + fmt.Sprintf(` http-equiv="%s"`, m.HttpEquiv)
	}
	if attrs == "" {
		return nil
	}
	_, err := fmt.Fprintf(w, `<meta%s>`, attrs)
	return err
}

// Link <link> tag
type Link struct {
	Rel         string
	Type        string
	Href        string
	Integrity   string
	Crossorigin string
}

func (l *Link) Render(w io.Writer) error {
	attrs := ""

	if l.Href != "" {
		if strings.HasSuffix(l.Href, ".css") {
			l.Rel = "stylesheet"
			l.Type = "text/css"
		}
		attrs = attrs + fmt.Sprintf(` href="%s"`, l.Href)
	}

	if l.Rel != "" {
		attrs = attrs + fmt.Sprintf(` rel="%s"`, l.Rel)
	}

	if l.Type != "" {
		attrs = attrs + fmt.Sprintf(` type="%s"`, l.Type)
	}

	if l.Integrity != "" {
		attrs = attrs + fmt.Sprintf(` integrity="%s"`, l.Integrity)
	}

	if l.Crossorigin != "" {
		attrs = attrs + fmt.Sprintf(` crossorigin="%s"`, l.Crossorigin)
	}

	if attrs == "" {
		return nil
	}
	_, err := fmt.Fprintf(w, `<link%s>`, attrs)
	return err
}

// Style <style> tag
type Style struct {
	Type string
	Body template.CSS
}

func (s *Style) Render(w io.Writer) error {
	attrs := ""
	if s.Type != "" {
		attrs = attrs + fmt.Sprintf(` type="%s"`, s.Type)
	}
	_, err := fmt.Fprintf(w, `<style%s>%s</style>`, attrs, s.Body)
	return err
}

// CSSLink plain local css link
type CSSLink string

// Render ...
// <link rel="stylesheet" type="text/css" href="/static/css/all.min.css">
func (s CSSLink) Render(w io.Writer) error {
	_, err := fmt.Fprintf(w, `<link rel="stylesheet" type="text/css" href="%s">`, s)
	return err
}
