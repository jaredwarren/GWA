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
		// "\n☐ make 'app' class that's full-screen, merge with native app\n",
		// "\n☐ figure a good way to load item from file\n", // done?
		"\n☐ figure a way to re-attach to web service (keep running in backgorund (might have to be saparate app))\n",
		"\n☐ Figure a way to update single panel without page reload!!!\n",
		"\n☐ figure way to make controller work -> pass ui to controller, ui.bind?\n",
		"\n☐ store (get data from ui.bind->ui.eval? or ajax, something...)\n",
		"\n☐ fix handler problem with type (might have to make all the same!), wonder if I can override json marshaller?, if not then what?\n",
		"\n☐ \n",
	}

	// this is here to show that objects can be in a saparate file
	mainController = &ext.Controller{
		Handlers: ext.Handlers{
			"btnClick": func(id string) {
				fmt.Print("Button Clicked:")
				fmt.Printf("   %+v\n", id)

				// Button update test
				btn := app.Find(id)
				if btn != nil {
					btn.(*ext.Button).Text = "Clicked!!!"
					app.Update(btn)
				}

				// Update Tree Test
				t := app.Find("tree-0")
				if t != nil {
					t.(*ext.Tree).Root.Text = "UPDATED"
					app.Update(t)
				}
			},
		},
	}
	app *ext.Application
)

func main() {
	fmt.Println("TODO:", TODO)

	app = &ext.Application{
		Name: "my app",
		Controllers: []*ext.Controller{
			mainController,
		},
		MainView: &ext.Panel{
			Title:  "Panel Title!",
			Shadow: true,
			Layout: "hbox",
			HTML:   "test",
			Items: []ext.Renderer{
				&ext.Panel{
					HTML:   "My panel text...1",
					Docked: "top",
				},
				// &ext.Panel{
				// 	HTML:   "My panel text...2",
				// 	Docked: "right",
				// },
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
				&ext.Tree{
					Docked:   "right",
					ShowRoot: true,
					Root: &ext.TreeNode{
						Text: "root",
						Children: []*ext.TreeNode{
							&ext.TreeNode{
								Text: "c1",
							},
							&ext.TreeNode{
								Text: "c2",
								Children: []*ext.TreeNode{
									&ext.TreeNode{
										Text:     "c2c1",
										Children: []*ext.TreeNode{},
									},
								},
							},
						},
					},
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
