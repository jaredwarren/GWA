package ext

import (
	"fmt"
	"io"
	"net/http"

	"github.com/zserge/lorca"
)

// Controller ...
type Controller struct {
	Handlers     Handlers
	FormHandlers FormHandlers
	ui           lorca.UI
}

// Handlers ...
type Handlers map[string]Handler

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
func (c *Controller) Render(w io.Writer) error {
	for name, f := range c.Handlers {
		// for some reason this has to be async, I think it's because ui isn't running yet
		go c.ui.Bind(name, f)
	}

	for name, f := range c.FormHandlers {
		web.Mux.HandleFunc(fmt.Sprintf("/submit/%s", name), f).Methods("POST")
	}

	// currently nothing to render
	return nil
}

// GetID ...
func (c *Controller) GetID() string {
	return ""
}
