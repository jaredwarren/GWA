package ext

import (
	"fmt"
	"html/template"
	"io"
)

type ScriptType string // string for now...

// Script gerneric div container
type Script struct {
	Type        ScriptType
	Src         string
	Integrity   string
	Crossorigin string
	InnerJS     template.JS
}

// Render ...
func (s *Script) Render(w io.Writer) error {
	attrs := ""

	src := s.Src
	if src != "" {
		s.InnerJS = ""
		attrs = attrs + fmt.Sprintf(` src="%s"`, src)
	}

	if s.Type != "" {
		attrs = attrs + fmt.Sprintf(` type="%s"`, s.Type)
	}

	if s.Integrity != "" {
		attrs = attrs + fmt.Sprintf(` integrity="%s"`, s.Integrity)
	}

	if s.Crossorigin != "" {
		attrs = attrs + fmt.Sprintf(` crossorigin="%s"`, s.Crossorigin)
	}

	_, err := fmt.Fprintf(w, `<script%s>%s</script>`, attrs, s.InnerJS)
	return err
}
