package gbt

import (
	"html/template"
)

type ScriptType string // string for now...

// Script gerneric div container
type Script struct {
	Type        ScriptType
	Src         string
	Integrity   string
	Crossorigin string
	InnerJS     template.JS
	Attributes
}

// Render ...
func (s *Script) Render() Stringer {
	if s.Attributes == nil {
		s.Attributes = Attributes{}
	}
	src := s.Src
	if src != "" {
		s.InnerJS = ""
		s.Attributes["src"] = src
	}

	if s.Type != "" {
		s.Attributes["type"] = s.Type
	}

	if s.Integrity != "" {
		s.Attributes["integrity"] = s.Integrity
	}

	if s.Crossorigin != "" {
		s.Attributes["crossorigin"] = s.Crossorigin
	}

	return renderToHTML(`<script {{.Attributes.Render}}>{{.InnerJS}}</script>`, s)
}
