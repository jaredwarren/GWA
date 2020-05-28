package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jaredwarren/goext/ext"
)

var (
	// TODO ...
	TODO = []string{
		"\n☐ json to panel (use xtype)",         // big problem with functions
		"\n☐ make header 'docked'",              // I think this is done, but header template needs cleaned up
		"\n☐ fix things so they work in test\n", // almost done
		// "\n☐ make 'app' class that's full-screen, merge with native app\n",
		// "\n☐ figure a good way to load item from file\n", // done?
		"\n☐ figure a way to re-attach to web service (keep running in backgorund (might have to be saparate app))\n",
		"\n☐ Figure a way to update single panel without page reload!!!\n",
		"\n☐ figure way to make controller work -> pass ui to controller, ui.bind?\n",
		"\n☐ store (get data from ui.bind->ui.eval? or ajax, something...)\n",
		"\n☐ fix handler problem with type (might have to make all the same!), wonder if I can override json marshaller?, if not then what?\n",
		"\n☐ create FORM and figure good way to submit to controller\n",
		"\n☐ create multiple sessions/instances of app\n",
		"\n☐ create multiple windows\n",
		"\n☐ save app state, have to do manually\n",
		"\n☐ Look for template e.g. panel.html\n",
		"\n☐ \n",
		"\n☐ replace all woff2 in pro.min.css https://kit-pro.fontawesome.com/releases/v5.13.0/webfonts/pro-fa-brands-400-5.12.0.woff2\n",
		"\n☐ \n",
		"\n☐ \n",
		"\n☐ Tree Handler needs to propogate to all leaf(or all?) nodes\n",
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
			"onPalm": func(id string) {
				fmt.Println("onPalm:", id)
			},
			"onJet": func(id string) {
				fmt.Println("onJet:", id)
			},
		},
		FormHandlers: ext.FormHandlers{
			"formSubmit": func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("submit....")
			},
		},
	}
	app *ext.Application
)

func main() {
	// app = loadFromJSON()

	app = load()

	done := app.Launch()
	if done != nil {
		fmt.Println("Something Happened, Bye!", done)
	} else {
		fmt.Println("Good Bye!")
	}
}

func loadFromJSON() *ext.Application {
	dat, err := ioutil.ReadFile("./app.json")
	app := &ext.Application{}
	err = json.Unmarshal(dat, app)
	if err != nil {
		fmt.Println(err)
	}

	app.Controllers = []*ext.Controller{
		mainController,
	}

	return app
}

func load() *ext.Application {
	return &ext.Application{
		XType: "app",
		Name:  "my app",
		Controllers: []*ext.Controller{
			mainController,
		},
		MainView: &ext.Panel{
			XType: "panel",
			// Title: "Panel Title!",
			// Header: &ext.Header{
			// 	XType: "header",
			// 	Title: "header title",
			// 	Items: ext.Items{
			// 		&ext.Button{
			// 			XType: "button",
			// 			// Text:  "Click Here",
			// 			IconClass: "fad fa-window-close",
			// 			Handler:   "close",
			// 		},
			// 	},
			// },
			Nav: &ext.Nav{
				Title: "Nav Title",
			},
			Shadow: true,
			Layout: "hbox",
			HTML:   "test",
			Items: []ext.Renderer{
				&ext.Panel{
					XType:  "panel",
					HTML:   "My panel text...1",
					Docked: "top",
				},
				&ext.Panel{
					XType:  "panel",
					HTML:   "My panel text...3",
					Docked: "left",
				},
				&ext.Panel{
					XType:  "panel",
					HTML:   "My panel text...4",
					Docked: "bottom",
				},
				&ext.Button{
					XType:     "button",
					Text:      "Click Here",
					Handler:   "btnClick",
					IconClass: "fad fa-window-close",
				},
				&ext.Button{
					XType: "button",
					Text:  "2 Here",
					HandlerFn: func(id string) {
						fmt.Print("Button 2 Clicked:")
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
				&ext.Form{
					XType: "form",
					// Text:    "Click Here",
					// Handler: "btnClick",
					// Method: "post",
					// Action: "submit",
					Handler: "formSubmit",
					// Handler: func(w http.ResponseWriter, r *http.Request) {
					// 	fmt.Println("submit....")
					// },

					Items: []ext.Renderer{
						&ext.Fieldset{
							XType:  "fieldset",
							Legend: "Form Legend",
							Items: []ext.Renderer{
								&ext.Input{
									XType: "input",
									Label: "User Name:",
									Name:  "username",
									Type:  "text",
								},
								&ext.Input{
									XType: "input",
									Label: "Send:",
									Name:  "submit",
									Type:  "submit",
								},
							},
						},
					},
				},
				&ext.Tree{
					XType:      "tree",
					Docked:     "right",
					ShowRoot:   true,
					BranchIcon: "",
					LeafIcon:   "",
					Root: &ext.TreeNode{
						Text: "root",
						// IconClass: "fas fa-folder-open",
						Children: []*ext.TreeNode{{
							Text:      "c1",
							IconClass: "fas fa-fighter-jet",
							Handler:   "onJet",
						}, {
							Text:      "c2",
							IconClass: "fad fa-acorn",
						}, {
							Text:      "c3",
							IconClass: "fad fa-arrow-alt-from-right",
						}, {
							Text:      "c4",
							IconClass: "fad fa-tree-palm",
							Handler:   "onPalm",
						}, {
							Text: "c2",
							Children: []*ext.TreeNode{{
								Text:     "c2c1",
								Children: []*ext.TreeNode{},
							}},
						}, {
							Text: "c3",
							Children: []*ext.TreeNode{{
								Text:     "c3c1",
								Children: []*ext.TreeNode{},
							}},
						}},
					},
				},
			},
		},
	}
}
