package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jaredwarren/goext/ext"
	v2 "github.com/jaredwarren/goext/v2"
)

type Button struct {
	Handler func()
}

var (

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
			"logout": func(id string) {
				fmt.Println("logout:", id)
			},
			"onTableSelect": func(id string) {
				fmt.Println("onTableSelect:", id)
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
	if true {
		v2.Load()
		fmt.Println("Bye...")
		return
	}
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
			Nav: &ext.Nav{
				Title:  "Nav Title",
				Shadow: true,
				Items: ext.Items{
					&ext.Button{
						XType:     "button",
						Text:      "",
						Handler:   "logout",
						IconClass: "fad fa-sign-out",
					},
				},
			},
			// Shadow: true,
			// Layout: "hbox",
			// HTML:   "test",
			Items: ext.Items{
				&ext.Panel{
					Docked:      "left",
					Collapsible: true,
					Collapsed:   true,
					Title:       "tree panel title",
					Items: ext.Items{
						&ext.Tree{
							XType:    "tree",
							ShowRoot: true,
							Width:    300,
							Title:    "Tree Stuff",
							// BranchIcon: "",
							// LeafIcon:   "",
							// ParentIcon: "",
							Root: &ext.TreeNode{
								Text: "root",
								// IconClass: "fas fa-folder-open",
								Children: []*ext.TreeNode{{
									Text:      "account_deletion_requests",
									IconClass: "fad fa-table",
									Handler:   "onTableSelect",
								}, {
									Text:      "c2",
									IconClass: "fad fa-acorn",
									Items: ext.Items{
										&ext.Button{
											UI:        "none",
											IconClass: "far fa-key",
											Handler:   "key",
										},
										&ext.Button{
											UI:        "none",
											IconClass: "far fa-info-circle",
											Handler:   "info",
										},
									},
								}, {
									Text: "|c3",
									// IconClass: "fad fa-arrow-alt-from-right",
								}, {
									Text:      "|c4",
									IconClass: "fad fa-tree-palm",
									Handler:   "onPalm",
								}, {
									Text: `<i class="fad fa-database"></i> Bladehq`,
									Children: []*ext.TreeNode{{
										Text: "c2c1",
									}},
								}, {
									Text: "c3",
									Children: []*ext.TreeNode{{
										Text: "c3c1",
									}},
								}},
							},
						},
					},
				},

				&ext.Form{
					XType: "form",
					// Docked: "top",
					// Text:    "Click Here",
					// Handler: "btnClick",
					// Method: "post",
					// Action: "submit",
					Resize:  "vertical",
					Handler: "formSubmit",
					// Handler: func(w http.ResponseWriter, r *http.Request) {
					// 	fmt.Println("submit....")
					// },

					Items: ext.Items{
						&ext.Panel{
							XType:  "panel",
							Docked: "bottom",
							Layout: "hbox",
							// Resize:  "vertical",
							Classes: ext.Classes{"toolbar"},
							Items: ext.Items{
								&ext.Button{
									XType:     "button",
									Text:      "Run",
									Handler:   "btnClick",
									IconClass: "far fa-play",
									Classes:   ext.Classes{"button-success", "pure-button"},
								},
								&ext.Spacer{},
								&ext.Input{
									XType: "input",
									Label: "limit:",
									Name:  "limit",
									Type:  "number",
									Data: ext.Data{
										"dname": "v",
									},
									Events: ext.Events{
										"keyup": &ext.Event{
											Handler: "limitChange",
										},
									},
								},
								&ext.Input{
									XType: "input",
									Label: "Show All:",
									Name:  "show_all",
									Type:  "checkbox",
								},
								&ext.Button{
									XType:     "button",
									Text:      "Click Here",
									Handler:   "btnClick",
									IconClass: "fad fa-window-close",
								},
							},
						},
						&ext.Fieldset{
							XType:  "fieldset",
							Legend: "Form Legend",
							Resize: "vertical",
							Items: ext.Items{
								&ext.Input{
									XType: "input",
									Label: "User Name:",
									Name:  "username",
									Type:  "textarea",
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

				&ext.Table{
					Title: "Table data",
					// Header: ext.TableHeader{{{
					// 	Innerhtml: "asdf",
					// 	Attributes: ext.Attributes{
					// 		"colspan": "3",
					// 	},
					// }}},
					Header: ext.TableHeader{{{
						Innerhtml: "ID",
						DataIndex: "id",
					}, {
						Innerhtml: "Select",
						DataIndex: "select",
					}, {
						Innerhtml: ` <a class="button-xsmall pure-button button-success"
          href="/{{$.SelectedDB}}/run?query={{$.Query}}&sortdir=ASC&sortname={{$c}};"><i
            class="fas fa-sort-amount-up-alt"></i></a>
        <a class="button-xsmall pure-button button-success"
          href="/{{$.SelectedDB}}/run?query={{$.Query}}&sortdir=DESC&sortname={{$c}};"> <i
            class="fas fa-sort-amount-down"></i></a>
        <button class="button-xsmall pure-button button-secondary" onclick="search('{{$c}}')"><i
            class="fas fa-search"></i></button>`,
						DataIndex: "something",
					}}},
					Footer: ext.TableFooter{{{
						Innerhtml: "asdf",
						Attributes: ext.Attributes{
							"colspan": "3",
						},
					}}},
					Data: ext.Rows{{
						"id": 1,
						"select": &ext.Button{
							IconClass: "fad fa-times-circle",
							UI:        "none",
						},
						"something": "something...",
					}},
				},

				// &ext.Panel{
				// 	XType:  "panel",
				// 	HTML:   "TABLE",
				// 	Docked: "left",
				// },
				// &ext.Panel{
				// 	XType:  "panel",
				// 	HTML:   "My panel text...4",
				// 	Docked: "bottom",
				// },
				// &ext.Button{
				// 	XType:     "button",
				// 	Text:      "Click Here",
				// 	Handler:   "btnClick",
				// 	IconClass: "fad fa-window-close",
				// },
				// &ext.Button{
				// 	XType: "button",
				// 	Text:  "2 Here",
				// 	HandlerFn: func(id string) {
				// 		fmt.Print("Button 2 Clicked:")
				// 		fmt.Printf("   %+v\n", id)

				// 		// Button update test
				// 		btn := app.Find(id)
				// 		if btn != nil {
				// 			btn.(*ext.Button).Text = "Clicked!!!"
				// 			app.Update(btn)
				// 		}

				// 		// Update Tree Test
				// 		t := app.Find("tree-0")
				// 		if t != nil {
				// 			t.(*ext.Tree).Root.Text = "UPDATED"
				// 			app.Update(t)
				// 		}
				// 	},
				// },
			},
		},
	}
}
