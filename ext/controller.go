package ext

import (
	"html/template"
	"io"
)

// Controller ...
type Controller struct {
	// Name   string
	// Tables []Table
	Handlers Handlers
}

// Handlers ...
type Handlers map[template.JS]Handler

// Handler ...
type Handler func(arg ...interface{})

// Call ...
func (h Handler) Call() {
	// h()
}

// Render ...
func (c *Controller) Render(w io.Writer) error {
	// TODO: find ui, and call ui.bind????
	return renderTemplate(w, "controller", c)
}
