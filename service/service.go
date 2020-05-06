package service

import (
	"fmt"
	"html/template"
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

	html := template.HTML("")
	// html = t1()
	html = t3()
	// html = t2()

	templates := template.Must(template.ParseFiles("templates/base.html"))
	templates.ExecuteTemplate(w, "base", &struct {
		Title string
		// Error   error
		Body template.HTML
	}{
		Title: "test",
		// Error:   nil,
		Body: html,
	})
}

func t3() template.HTML {
	/*
			Ext.create({
		    xtype: 'panel',

		    renderTo: Ext.getBody(),
		    width: 300,
		    height: 200,

		    shadow: true,
		    scrollable: true,
		    padding: 20,

		    title: 'Panel',

		    layout: 'hbox',

		    defaults: {
		        flex: 1
			},
			html: "test",

		    items: [{
		        html: 'My panel text...',
		        style: '',
		        docked: "left"
		    },{
		        html: 'My panel text...',
		        style: '',
		        docked: "left"
		    }]
		});

	*/
	panel := &ext.Panel{
		Title:  "Panel",
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
				Style:  "",
				Flex:   1, // because default doesn't work
			},
			&ext.Panel{
				HTML:   "My panel text...2",
				Docked: "left",
				Style:  "",
				Flex:   1, // because default doesn't work
			},
		},
	}

	return panel.Render()
}

func t1() template.HTML {

	panel := &ext.Panel{
		Title: "test2",
		Items: []ext.Renderer{
			&ext.Panel{
				HTML: template.HTML("<div>This is the sub body!</div>"),
				Items: []ext.Renderer{&ext.Button{
					Text:      "howdy",
					Handler:   "alert('yo')",
					IconClass: "fas fa-home",
				}},
			},
			&ext.Panel{
				Docked: "bottom",
				Items: []ext.Renderer{&ext.Button{
					Text:      "howdy",
					Handler:   "alert('yo')",
					IconClass: "fas fa-home",
				}},
			},
		},
	}

	return panel.Render()
}

func t2() template.HTML {

	panel := ext.NewPanel()
	panel.Title = "test"
	// panel.HTML = template.HTML("<div>This is the body!</div>")

	sp := ext.NewPanel()
	// panel.Title = "test"
	sp.HTML = template.HTML("<div>This is the sub body!</div>")
	sp.Items = []ext.Renderer{&ext.Button{
		Text:      "howdy",
		Handler:   "alert('yo')",
		IconClass: "fas fa-home",
	}}

	panel.Items = []ext.Renderer{sp}
	return panel.Render()
}

// func render(item ext.Renderer) {
// 	fmt.Printf("%s -> %+v\n", reflect.TypeOf(item), item)

// 	fmt.Printf("\n<html>%s</html>\n", item.Render())
// }
