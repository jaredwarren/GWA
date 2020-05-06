package main

import (
	"fmt"

	"github.com/jaredwarren/app"
	"github.com/jaredwarren/goext/service"
)

func main() {
	serve()
}

func serve() {

	// panel := &ext.Panel{
	// 	Title:  "Panel",
	// 	Width:  300,
	// 	Height: 200,
	// 	Shadow: true,
	// 	// Scrollable: true,
	// 	// Padding: 20,
	// 	Layout: "hbox",
	// 	// Defaults: []interface{}{
	// 	// 	Flex: 1,
	// 	// },
	// 	HTML: "test",
	// 	Items: []ext.Renderer{
	// 		&ext.Panel{
	// 			HTML:   "My panel text...1",
	// 			Docked: "left",
	// 			Style:  "",
	// 			Flex:   1, // because default doesn't work
	// 		},
	// 		&ext.Panel{
	// 			HTML:   "My panel text...2",
	// 			Docked: "left",
	// 			Style:  "",
	// 			Flex:   1, // because default doesn't work
	// 		},
	// 	},
	// }
	// panel.Render()

	// panel.Debug()

	// return
	conf := &app.WebConfig{
		Host: "127.0.0.1",
		Port: 8083,
	}
	a := app.NewWeb(conf)

	service.Register(a)

	d := <-a.Exit
	fmt.Printf("Done:%+v\n", d)
}
