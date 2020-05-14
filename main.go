package main

import (
	"fmt"

	"github.com/jaredwarren/goext/ext"
)

var (
	// TODO ...
	TODO = []string{
		"\n☐ json to panel (use xtype)",
		"\n☐ make header 'docked'",              // I think this is done, but header template needs cleaned up
		"\n☐ fix things so they work in test\n", // almost done
		"\n☐ make 'app' class that's full-screen, merge with native app\n",
		"\n☐ figure a good way to load item from file\n",
		"\n☐ figure a way to re-attach to web service (keep running in backgorund (might have to be saparate app))\n",
		"\n☐ Figure a way to update single panel without page reload!!!\n",
		"\n☐ figure way to make controller work -> pass ui to controller, ui.bind?\n",
		"\n☐ \n",
	}
)

func main() {
	fmt.Println("TODO:", TODO)

	app := &ext.Application{
		Name: "my app",
		MainView: &ext.Panel{
			Title:  "Panel Title!",
			Shadow: true,
			Layout: "hbox",
			Controller: &ext.Controller{
				Handlers: ext.Handlers{
					"btnClick": func(args ...interface{}) {
						fmt.Println("Button Clicked!!!!")
					},
				},
			},
			HTML: "test",
			Items: []ext.Renderer{
				&ext.Panel{
					HTML:   "My panel text...1",
					Docked: "top",
				},
				&ext.Panel{
					HTML:   "My panel text...2",
					Docked: "right",
				},
				&ext.Panel{
					HTML:   "My panel text...3",
					Docked: "left",
				},
				&ext.Panel{
					HTML:   "My panel text...4",
					Docked: "bottom",
				},
				&ext.Button{
					Text:    "Click Here",
					Handler: "btnClick",
				},
			},
		},
	}
	done := app.Launch()
	if done != nil {
		fmt.Println("Something Happened, Bye!", done)
	} else {
		fmt.Println("Good Bye!")
	}
}
