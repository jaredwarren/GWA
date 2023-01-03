package ext

import (
	"fmt"
	"html/template"
	"io"
	"net/url"
)

type ScriptType string // string for now...

// Script gerneric div container
type Script struct {
	Type    ScriptType
	Src     url.URL
	InnerJS template.JS
}

// Render ...
func (s *Script) Render(w io.Writer) error {
	name := "script"

	attrs := ""

	src := s.Src.String()
	if src != "" {
		s.InnerJS = ""
		attrs = fmt.Sprintf("%s %s=\"%s\"", attrs, "src", src)
	}

	if s.Type != "" {
		attrs = fmt.Sprintf("%s %s=\"%s\"", attrs, "type", s.Type)
	}

	_, err := fmt.Fprintf(w, `<%s %s>%s</%s>`, name, attrs, s.InnerJS, name)
	return err
	// if s.InnerJS != "" {
	// }

	// return render(w, fmt.Sprintf(`<%s %s>{{range $item := $.Items}}{{if $item}}{{Render $item}}{{else}}NULL---{{end}}{{end}}</%s>`, name, attrs, name), e)
}

// GetID ...
func (s *Script) GetID() string {
	return ""
}
