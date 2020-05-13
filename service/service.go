package service

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jaredwarren/app"
	"github.com/jaredwarren/goext/ext"
)

// Controller implements the home resource.
type Controller struct {
	Mux  *mux.Router
	wapp *app.Service
}

// Register attach mux to service
func Register(wapp *app.Service) *Controller {
	c := &Controller{
		Mux:  wapp.Mux,
		wapp: wapp,
	}

	m := wapp.Mux
	m.HandleFunc("/", c.Home)

	return c
}

// Close handler.
func (c *Controller) Close(w http.ResponseWriter, r *http.Request) {
	c.wapp.Exit <- nil
}

// Home show current invoice
func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")

	// html = t1()
	buf := new(bytes.Buffer)
	// item.Render(buf)
	err := t3(buf)
	if err != nil {
		fmt.Println("[EE]", err)
	}
	// html = t2()

	templates := template.Must(template.ParseFiles("templates/base.html"))
	templates.ExecuteTemplate(w, "base", &struct {
		Title string
		// Error   error
		Body template.HTML
	}{
		Title: "test",
		// Error:   nil,
		Body: template.HTML(buf.String()),
	})
}

func t3(w io.Writer) error {
	app := &ext.Application{
		Name: "my app",
		MainView: &ext.Panel{
			Title:  "Panel Title!",
			Width:  300,
			Height: 200,
			Shadow: true,
			// Scrollable: true,
			// Padding: 20,
			Layout: "hbox",
			// Defaults: []interface{}{
			// 	Flex: 1,
			// },
			HTML: "test",
			Items: []ext.Renderer{
				&ext.Panel{
					HTML:   "My panel text...1",
					Docked: "top",
					// Style:  "",
					Flex: 1, // because default doesn't work
				},
				&ext.Panel{
					HTML:   "My panel text...2",
					Docked: "right",
					// Style:  "",
					Flex: 1, // because default doesn't work
				},
			},
		},
	}

	return app.Render(w)
}
