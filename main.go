package main

import (
	"fmt"
	"net/http"

	"github.com/jaredwarren/goext/gbt"
)

type Button struct {
	Handler func()
}

var (

	// this is here to show that objects can be in a saparate file
	mainController = &gbt.Controller{
		Handlers: gbt.Handlers{
			// "btnClick": func(id string) {
			// 	fmt.Print("Button Clicked:")
			// 	fmt.Printf("   %+v\n", id)

			// 	// Button update test
			// 	btn := app.Find(id)
			// 	if btn != nil {
			// 		btn.(*gbt.Button).Text = "Clicked!!!"
			// 		app.Update(btn)
			// 	}

			// 	// Update Tree Test
			// 	t := app.Find("tree-0")
			// 	if t != nil {
			// 		t.(*gbt.Tree).Root.Text = "UPDATED"
			// 		app.Update(t)
			// 	}
			// },
			// "logout": func(id string) {
			// 	fmt.Println("logout:", id)
			// },
			// "onTableSelect": func(id string) {
			// 	fmt.Println("onTableSelect:", id)
			// },
			"btnClick": `
			debugger;
		for (let i = 0; i < args.length; i++) {
 			console.log(args[i]);
		}
		`,
		},
		FormHandlers: gbt.FormHandlers{
			"formSubmit": func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("submit....")
			},
		},
	}
	app *gbt.Application
)

func main() {
	app = load()

	done := app.Launch()
	if done != nil {
		fmt.Println("Something Happened, Bye!", done)
	} else {
		fmt.Println("Good Bye!")
	}
}

func load() *gbt.Application {
	return &gbt.Application{
		XType: "app",
		Name:  "my app",
		Controllers: []*gbt.Controller{
			mainController,
		},
		Head: &gbt.Head{
			Title: "this is the title",
			// TODO: some of these can be made "default"
			Items: gbt.Items{
				// <meta>
				&gbt.Meta{
					Charset: "utf-8",
				},
				&gbt.Meta{
					Name:    "viewport",
					Content: "width=device-width, initial-scale=1, shrink-to-fit=no",
				},
				&gbt.Meta{
					HttpEquiv: "Content-Type",
					Content:   "text/html; charset=utf-8",
				},
				// Link/CSS
				gbt.CSSLink("https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200"),
				&gbt.Link{
					Href:        "https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css",
					Integrity:   "sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD",
					Crossorigin: "anonymous",
				},
				// Script
				&gbt.Script{
					Src:         "https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js",
					Integrity:   "sha384-w76AqPfDkMBDXo30jS1Sgez6pr3x5MlQ1ZAGC+nuZB+EYdgRZgiwxhTBTkF7CXvN",
					Crossorigin: "anonymous",
				},
				&gbt.Script{Src: "/static/js/test.js"},
			},
		},
		Nav: &gbt.Nav{
			Title:  "Nav Title",
			Shadow: true,
			Search: true,
			Theme:  gbt.ThemeDark,
			Brand: &gbt.NavBrand{
				Image: &gbt.Image{
					Src:    "https://getbootstrap.com/docs/5.3/assets/brand/bootstrap-logo.svg",
					Height: "20px",
				},
				Title: "this is my brand",
				Href:  "#",
			},
			Items: []gbt.INavItem{
				&gbt.NavItem{
					Title: "Item 1",
					Href:  "1",
				},
				&gbt.NavDropDown{
					Title: "Drop 1",
					Items: []gbt.INavItem{
						&gbt.DropDownItem{
							Title:    "drop 1",
							Disabled: true,
						},
						&gbt.DropDownItem{
							Title:  "drop 2",
							Active: true,
						},
						&gbt.DropDowndivider{},
						&gbt.DropDownItem{
							Title: "drop 3",
						},
					},
				},
			},
		},
		MainView: &gbt.Panel{
			// XType: "panel",
			// Shadow: true,
			// Layout: "hbox",
			// HTML:   "test",
			Items: gbt.Items{
				&gbt.Card{
					Styles: gbt.Styles{"width": "18em"},
					Header: &gbt.CardHeader{
						Body: gbt.Items{gbt.RawHTML("header")},
					},
					Body: gbt.Items{
						&gbt.Image{
							Src:     "/static/alien.svg",
							Classes: gbt.Classes{"card-img-top"},
						},
						&gbt.CardBody{
							Title: "Card Title",
							Text:  "Some quick example text to build on the card title and make up the bulk of the card's content.",
						},
						&gbt.CardBody{
							Body: gbt.Items{
								&gbt.CardLink{
									Href: "#",
									Text: "Card Link",
								},
								&gbt.CardLink{
									Href: "#",
									Text: "Another link",
								},
							},
						},
					},
					Footer: &gbt.CardFooter{
						Body: gbt.Items{gbt.RawHTML("Footer...")},
					},
				},
				&gbt.Panel{
					Docked:      "left",
					Collapsible: true,
					Collapsed:   true,
					Title:       "tree panel title",
					Items: gbt.Items{
						&gbt.Tree{
							XType:    "tree",
							ShowRoot: true,
							Width:    300,
							Title:    "Tree Stuff",
							// BranchIcon: "",
							// LeafIcon:   "",
							// ParentIcon: "",
							Root: &gbt.TreeNode{
								Text: "root",
								// IconClass: "fas fa-folder-open",
								Children: []*gbt.TreeNode{{
									Text:      "account_deletion_requests",
									IconClass: "fad fa-table",
									Handler:   "onTableSelect",
								}, {
									Text:      "c2",
									IconClass: "fad fa-acorn",
									Items: gbt.Items{
										&gbt.Button{
											Handler: "key",
										},
										&gbt.Button{
											Handler: "info",
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
									Children: []*gbt.TreeNode{{
										Text: "c2c1",
									}},
								}, {
									Text: "c3",
									Children: []*gbt.TreeNode{{
										Text: "c3c1",
									}},
								}},
							},
						},
					},
				},

				&gbt.Form{
					Items: gbt.Items{
						&gbt.Fieldset{
							Legend: "NEW Form Legend!!!",
							Items: gbt.Items{
								&gbt.FormEmail{
									Label:       "Email address",
									Name:        "email",
									HelpText:    "We'll never share your email with anyone else.",
									Placeholder: "some@email.com",
								},
								&gbt.Input{
									// XType: "input",
									// Label: "User Name:",
									// Name:  "username",
									Type: "textarea",
								},
								&gbt.Input{
									// XType: "input",
									// Label: "Send:",
									// Name:  "submit",
									Type: "submit",
								},
								&gbt.Button{
									Text:    "Settings",
									Outline: true,
									// IconPosition: "right",
									Icon: &gbt.Icon{
										Icon: "settings",
									},
									Badge: &gbt.Badge{
										Text:  "100",
										Style: gbt.ButtonDanger,
									},
								},
							},
						},
					},
				},

				&gbt.Table{
					Title: "Table data",
					// Header: gbt.TableHeader{{{
					// 	Innerhtml: "asdf",
					// 	Attributes: gbt.Attributes{
					// 		"colspan": "3",
					// 	},
					// }}},
					Header: gbt.TableHeader{{{
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
					Footer: gbt.TableFooter{{{
						Innerhtml: "asdf",
						Attributes: gbt.Attributes{
							"colspan": "3",
						},
					}}},
					Data: gbt.Rows{{
						"id":        1,
						"select":    &gbt.Button{},
						"something": "something...",
					}},
				},
			},
		},
	}
}
