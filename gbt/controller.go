package gbt

import (
	"fmt"
	"html/template"
	"net/http"
)

// Controller ...
type Controller struct {
	Handlers     Handlers
	FormHandlers FormHandlers
	// ui           lorca.UI
}

// Handlers ...
type Handlers map[string]template.JS

// FormHandlers ...
type FormHandlers map[string]FormHandler

// Handler ...
type Handler func(id string)

// FormHandler ...
type FormHandler func(w http.ResponseWriter, r *http.Request)

// Call ...
func (h Handler) Call(args ...interface{}) {
	// h(args...)
}

// Render ...
func (c *Controller) Render() Stringer {
	var innerJS template.JS
	for k, h := range c.Handlers {
		innerJS = innerJS + template.JS(fmt.Sprintf(`const %s = function (...args) {%s}`, k, h))
	}
	script := &Script{
		InnerJS: innerJS,
	}

	return script.Render()
}

// GetID ...
func (c *Controller) GetID() string {
	return ""
}
