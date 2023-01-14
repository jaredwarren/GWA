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

	err := app.ToHTML("app.html")
	if err != nil {
		panic(err.Error())
	}

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
			// Brand: &gbt.NavBrand{
			// 	Image: &gbt.Image{
			// 		Src:    "https://getbootstrap.com/docs/5.3/assets/brand/bootstrap-logo.svg",
			// 		Height: "20px",
			// 	},
			// 	Title: "this is my brand",
			// 	Href:  "#",
			// },
			Brand: gbt.NewBrand(
				gbt.NavImage("https://getbootstrap.com/docs/5.3/assets/brand/bootstrap-logo.svg"),
				gbt.NavTitle("This is the title"),
			),
			Items: gbt.Items{
				&gbt.NavItem{
					Title: "Item 1",
					Href:  "1",
				},
				&gbt.NavDropDown{
					Title: "Drop 1",
					Items: gbt.Items{
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
				&gbt.Element{
					Attributes: gbt.Attributes{},
					Classes:    gbt.Classes{"container-fluid"},
					Items: gbt.Items{&gbt.Element{
						Attributes: gbt.Attributes{},
						Classes:    gbt.Classes{"row", "flex-nowrap"},
						Items: gbt.Items{&gbt.Element{
							Attributes: gbt.Attributes{},
							Classes:    gbt.Classes{"col-auto", "col-md-3", "col-xl-2", "px-sm-2", "px-0", "bg-dark"},
							Items: gbt.Items{&gbt.Element{
								Attributes: gbt.Attributes{},
								Classes:    gbt.Classes{"d-flex", "flex-column", "align-items-center", "align-items-sm-start", "px-3", "pt-2", "text-white", "min-vh-100"},
								Items: gbt.Items{&gbt.Element{
									Attributes: gbt.Attributes{"href": "/"},
									Classes:    gbt.Classes{"d-flex", "align-items-center", "pb-3", "mb-md-0", "me-md-auto", "text-white", "text-decoration-none"},
									Items: gbt.Items{&gbt.Element{
										Attributes: gbt.Attributes{},
										Classes:    gbt.Classes{"fs-5", "d-none", "d-sm-inline"},
										Items:      gbt.Items{gbt.RawHTML("Menu")},
										Name:       "span",
									}},
									Name: "a",
								}, &gbt.Element{
									Attributes: gbt.Attributes{"id": "menu"},
									Classes:    gbt.Classes{"nav", "nav-pills", "flex-column", "mb-sm-auto", "mb-0", "align-items-center", "align-items-sm-start"},
									Items: gbt.Items{&gbt.Element{
										Attributes: gbt.Attributes{},
										Classes:    gbt.Classes{"nav-item"},
										Items: gbt.Items{&gbt.Element{
											Attributes: gbt.Attributes{"href": "#"},
											Classes:    gbt.Classes{"nav-link", "align-middle", "px-0"},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"fs-4", "bi-house"},
												Items:      gbt.Items{},
												Name:       "i",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"ms-1", "d-none", "d-sm-inline"},
												Items:      gbt.Items{gbt.RawHTML("Home")},
												Name:       "span",
											}},
											Name: "a",
										}},
										Name: "li",
									}, &gbt.Element{
										Attributes: gbt.Attributes{},
										Classes:    gbt.Classes{},
										Items: gbt.Items{&gbt.Element{
											Attributes: gbt.Attributes{
												"data-bs-toggle": "collapse",
												"href":           "#submenu1",
											},
											Classes: gbt.Classes{"nav-link", "px-0", "align-middle"},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"fs-4", "bi-speedometer2"},
												Items:      gbt.Items{},
												Name:       "i",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"ms-1", "d-none", "d-sm-inline"},
												Items:      gbt.Items{gbt.RawHTML("Dashboard")},
												Name:       "span",
											}},
											Name: "a",
										}, &gbt.Element{
											Attributes: gbt.Attributes{
												"data-bs-parent": "#menu",
												"id":             "submenu1",
											},
											Classes: gbt.Classes{"collapse", "show", "nav", "flex-column", "ms-1"},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"w-100"},
												Items: gbt.Items{&gbt.Element{
													Attributes: gbt.Attributes{"href": "#"},
													Classes:    gbt.Classes{"nav-link", "px-0"},
													Items: gbt.Items{&gbt.Element{
														Attributes: gbt.Attributes{},
														Classes:    gbt.Classes{"d-none", "d-sm-inline"},
														Items:      gbt.Items{gbt.RawHTML("Item")},
														Name:       "span",
													}, gbt.RawHTML("1")},
													Name: "a",
												}},
												Name: "li",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{},
												Items: gbt.Items{&gbt.Element{
													Attributes: gbt.Attributes{"href": "#"},
													Classes:    gbt.Classes{"nav-link", "px-0"},
													Items: gbt.Items{&gbt.Element{
														Attributes: gbt.Attributes{},
														Classes:    gbt.Classes{"d-none", "d-sm-inline"},
														Items:      gbt.Items{gbt.RawHTML("Item")},
														Name:       "span",
													}, gbt.RawHTML("2")},
													Name: "a",
												}},
												Name: "li",
											}},
											Name: "ul",
										}},
										Name: "li",
									}, &gbt.Element{
										Attributes: gbt.Attributes{},
										Classes:    gbt.Classes{},
										Items: gbt.Items{&gbt.Element{
											Attributes: gbt.Attributes{"href": "#"},
											Classes:    gbt.Classes{"nav-link", "px-0", "align-middle"},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"fs-4", "bi-table"},
												Items:      gbt.Items{},
												Name:       "i",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"ms-1", "d-none", "d-sm-inline"},
												Items:      gbt.Items{gbt.RawHTML("Orders")},
												Name:       "span",
											}},
											Name: "a",
										}},
										Name: "li",
									}, &gbt.Element{
										Attributes: gbt.Attributes{},
										Classes:    gbt.Classes{},
										Items: gbt.Items{&gbt.Element{
											Attributes: gbt.Attributes{
												"data-bs-toggle": "collapse",
												"href":           "#submenu2",
											},
											Classes: gbt.Classes{"nav-link", "px-0", "align-middle"},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"fs-4", "bi-bootstrap"},
												Items:      gbt.Items{},
												Name:       "i",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"ms-1", "d-none", "d-sm-inline"},
												Items:      gbt.Items{gbt.RawHTML("Bootstrap")},
												Name:       "span",
											}},
											Name: "a",
										}, &gbt.Element{
											Attributes: gbt.Attributes{
												"data-bs-parent": "#menu",
												"id":             "submenu2",
											},
											Classes: gbt.Classes{"collapse", "nav", "flex-column", "ms-1"},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"w-100"},
												Items: gbt.Items{&gbt.Element{
													Attributes: gbt.Attributes{"href": "#"},
													Classes:    gbt.Classes{"nav-link", "px-0"},
													Items: gbt.Items{&gbt.Element{
														Attributes: gbt.Attributes{},
														Classes:    gbt.Classes{"d-none", "d-sm-inline"},
														Items:      gbt.Items{gbt.RawHTML("Item")},
														Name:       "span",
													}, gbt.RawHTML("1")},
													Name: "a",
												}},
												Name: "li",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{},
												Items: gbt.Items{&gbt.Element{
													Attributes: gbt.Attributes{"href": "#"},
													Classes:    gbt.Classes{"nav-link", "px-0"},
													Items: gbt.Items{&gbt.Element{
														Attributes: gbt.Attributes{},
														Classes:    gbt.Classes{"d-none", "d-sm-inline"},
														Items:      gbt.Items{gbt.RawHTML("Item")},
														Name:       "span",
													}, gbt.RawHTML("2")},
													Name: "a",
												}},
												Name: "li",
											}},
											Name: "ul",
										}},
										Name: "li",
									}, &gbt.Element{
										Attributes: gbt.Attributes{},
										Classes:    gbt.Classes{},
										Items: gbt.Items{&gbt.Element{
											Attributes: gbt.Attributes{
												"data-bs-toggle": "collapse",
												"href":           "#submenu3",
											},
											Classes: gbt.Classes{"nav-link", "px-0", "align-middle"},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"fs-4", "bi-grid"},
												Items:      gbt.Items{},
												Name:       "i",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"ms-1", "d-none", "d-sm-inline"},
												Items:      gbt.Items{gbt.RawHTML("Products")},
												Name:       "span",
											}},
											Name: "a",
										}, &gbt.Element{
											Attributes: gbt.Attributes{
												"data-bs-parent": "#menu",
												"id":             "submenu3",
											},
											Classes: gbt.Classes{"collapse", "nav", "flex-column", "ms-1"},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"w-100"},
												Items: gbt.Items{&gbt.Element{
													Attributes: gbt.Attributes{"href": "#"},
													Classes:    gbt.Classes{"nav-link", "px-0"},
													Items: gbt.Items{&gbt.Element{
														Attributes: gbt.Attributes{},
														Classes:    gbt.Classes{"d-none", "d-sm-inline"},
														Items:      gbt.Items{gbt.RawHTML("Product")},
														Name:       "span",
													}, gbt.RawHTML("1")},
													Name: "a",
												}},
												Name: "li",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{},
												Items: gbt.Items{&gbt.Element{
													Attributes: gbt.Attributes{"href": "#"},
													Classes:    gbt.Classes{"nav-link", "px-0"},
													Items: gbt.Items{&gbt.Element{
														Attributes: gbt.Attributes{},
														Classes:    gbt.Classes{"d-none", "d-sm-inline"},
														Items:      gbt.Items{gbt.RawHTML("Product")},
														Name:       "span",
													}, gbt.RawHTML("2")},
													Name: "a",
												}},
												Name: "li",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{},
												Items: gbt.Items{&gbt.Element{
													Attributes: gbt.Attributes{"href": "#"},
													Classes:    gbt.Classes{"nav-link", "px-0"},
													Items: gbt.Items{&gbt.Element{
														Attributes: gbt.Attributes{},
														Classes:    gbt.Classes{"d-none", "d-sm-inline"},
														Items:      gbt.Items{gbt.RawHTML("Product")},
														Name:       "span",
													}, gbt.RawHTML("3")},
													Name: "a",
												}},
												Name: "li",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{},
												Items: gbt.Items{&gbt.Element{
													Attributes: gbt.Attributes{"href": "#"},
													Classes:    gbt.Classes{"nav-link", "px-0"},
													Items: gbt.Items{&gbt.Element{
														Attributes: gbt.Attributes{},
														Classes:    gbt.Classes{"d-none", "d-sm-inline"},
														Items:      gbt.Items{gbt.RawHTML("Product")},
														Name:       "span",
													}, gbt.RawHTML("4")},
													Name: "a",
												}},
												Name: "li",
											}},
											Name: "ul",
										}},
										Name: "li",
									}, &gbt.Element{
										Attributes: gbt.Attributes{},
										Classes:    gbt.Classes{},
										Items: gbt.Items{&gbt.Element{
											Attributes: gbt.Attributes{"href": "#"},
											Classes:    gbt.Classes{"nav-link", "px-0", "align-middle"},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"fs-4", "bi-people"},
												Items:      gbt.Items{},
												Name:       "i",
											}, &gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"ms-1", "d-none", "d-sm-inline"},
												Items:      gbt.Items{gbt.RawHTML("Customers")},
												Name:       "span",
											}},
											Name: "a",
										}},
										Name: "li",
									}},
									Name: "ul",
								}, &gbt.Element{
									Attributes: gbt.Attributes{},
									Classes:    gbt.Classes{},
									Name:       "hr",
								}, &gbt.Element{
									Attributes: gbt.Attributes{},
									Classes:    gbt.Classes{"dropdown", "pb-4"},
									Items: gbt.Items{&gbt.Element{
										Attributes: gbt.Attributes{
											"aria-expanded":  "false",
											"data-bs-toggle": "dropdown",
											"href":           "#",
											"id":             "dropdownUser1",
										},
										Classes: gbt.Classes{"d-flex", "align-items-center", "text-white", "text-decoration-none", "dropdown-toggle"},
										Items: gbt.Items{&gbt.Element{
											Attributes: gbt.Attributes{
												"alt":    "hugenerd",
												"height": "30",
												"src":    "https://github.com/mdo.png",
												"width":  "30",
											},
											Classes: gbt.Classes{"rounded-circle"},
											Name:    "img",
										}, &gbt.Element{
											Attributes: gbt.Attributes{},
											Classes:    gbt.Classes{"d-none", "d-sm-inline", "mx-1"},
											Items:      gbt.Items{gbt.RawHTML("loser")},
											Name:       "span",
										}},
										Name: "a",
									}, &gbt.Element{
										Attributes: gbt.Attributes{},
										Classes:    gbt.Classes{"dropdown-menu", "dropdown-menu-dark", "text-small", "shadow"},
										Items: gbt.Items{&gbt.Element{
											Attributes: gbt.Attributes{},
											Classes:    gbt.Classes{},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{"href": "#"},
												Classes:    gbt.Classes{"dropdown-item"},
												Items:      gbt.Items{gbt.RawHTML("New project...")},
												Name:       "a",
											}},
											Name: "li",
										}, &gbt.Element{
											Attributes: gbt.Attributes{},
											Classes:    gbt.Classes{},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{"href": "#"},
												Classes:    gbt.Classes{"dropdown-item"},
												Items:      gbt.Items{gbt.RawHTML("Settings")},
												Name:       "a",
											}},
											Name: "li",
										}, &gbt.Element{
											Attributes: gbt.Attributes{},
											Classes:    gbt.Classes{},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{"href": "#"},
												Classes:    gbt.Classes{"dropdown-item"},
												Items:      gbt.Items{gbt.RawHTML("Profile")},
												Name:       "a",
											}},
											Name: "li",
										}, &gbt.Element{
											Attributes: gbt.Attributes{},
											Classes:    gbt.Classes{},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{},
												Classes:    gbt.Classes{"dropdown-divider"},
												Name:       "hr",
											}},
											Name: "li",
										}, &gbt.Element{
											Attributes: gbt.Attributes{},
											Classes:    gbt.Classes{},
											Items: gbt.Items{&gbt.Element{
												Attributes: gbt.Attributes{"href": "#"},
												Classes:    gbt.Classes{"dropdown-item"},
												Items:      gbt.Items{gbt.RawHTML("Sign out")},
												Name:       "a",
											}},
											Name: "li",
										}},
										Name: "ul",
									}},
									Name: "div",
								}},
								Name: "div",
							}},
							Name: "div",
						}, &gbt.Element{
							Attributes: gbt.Attributes{},
							Classes:    gbt.Classes{"col", "py-3"},
							Items:      gbt.Items{gbt.RawHTML("Content area...")},
							Name:       "div",
						}},
						Name: "div",
					}},
					Name: "div",
				},
				// 				gbt.RawHTML(`<div class="container-fluid">
				//     <div class="row flex-nowrap">
				//         <div class="col-auto col-md-3 col-xl-2 px-sm-2 px-0 bg-dark">
				//             <div class="d-flex flex-column align-items-center align-items-sm-start px-3 pt-2 text-white min-vh-100">
				//                 <a href="/" class="d-flex align-items-center pb-3 mb-md-0 me-md-auto text-white text-decoration-none">
				//                     <span class="fs-5 d-none d-sm-inline">Menu</span>
				//                 </a>
				//                 <ul class="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start" id="menu">
				//                     <li class="nav-item">
				//                         <a href="#" class="nav-link align-middle px-0">
				//                             <i class="fs-4 bi-house"></i> <span class="ms-1 d-none d-sm-inline">Home</span>
				//                         </a>
				//                     </li>
				//                     <li>
				//                         <a href="#submenu1" data-bs-toggle="collapse" class="nav-link px-0 align-middle">
				//                             <i class="fs-4 bi-speedometer2"></i> <span class="ms-1 d-none d-sm-inline">Dashboard</span> </a>
				//                         <ul class="collapse show nav flex-column ms-1" id="submenu1" data-bs-parent="#menu">
				//                             <li class="w-100">
				//                                 <a href="#" class="nav-link px-0"> <span class="d-none d-sm-inline">Item</span> 1 </a>
				//                             </li>
				//                             <li>
				//                                 <a href="#" class="nav-link px-0"> <span class="d-none d-sm-inline">Item</span> 2 </a>
				//                             </li>
				//                         </ul>
				//                     </li>
				//                     <li>
				//                         <a href="#" class="nav-link px-0 align-middle">
				//                             <i class="fs-4 bi-table"></i> <span class="ms-1 d-none d-sm-inline">Orders</span></a>
				//                     </li>
				//                     <li>
				//                         <a href="#submenu2" data-bs-toggle="collapse" class="nav-link px-0 align-middle ">
				//                             <i class="fs-4 bi-bootstrap"></i> <span class="ms-1 d-none d-sm-inline">Bootstrap</span></a>
				//                         <ul class="collapse nav flex-column ms-1" id="submenu2" data-bs-parent="#menu">
				//                             <li class="w-100">
				//                                 <a href="#" class="nav-link px-0"> <span class="d-none d-sm-inline">Item</span> 1</a>
				//                             </li>
				//                             <li>
				//                                 <a href="#" class="nav-link px-0"> <span class="d-none d-sm-inline">Item</span> 2</a>
				//                             </li>
				//                         </ul>
				//                     </li>
				//                     <li>
				//                         <a href="#submenu3" data-bs-toggle="collapse" class="nav-link px-0 align-middle">
				//                             <i class="fs-4 bi-grid"></i> <span class="ms-1 d-none d-sm-inline">Products</span> </a>
				//                             <ul class="collapse nav flex-column ms-1" id="submenu3" data-bs-parent="#menu">
				//                             <li class="w-100">
				//                                 <a href="#" class="nav-link px-0"> <span class="d-none d-sm-inline">Product</span> 1</a>
				//                             </li>
				//                             <li>
				//                                 <a href="#" class="nav-link px-0"> <span class="d-none d-sm-inline">Product</span> 2</a>
				//                             </li>
				//                             <li>
				//                                 <a href="#" class="nav-link px-0"> <span class="d-none d-sm-inline">Product</span> 3</a>
				//                             </li>
				//                             <li>
				//                                 <a href="#" class="nav-link px-0"> <span class="d-none d-sm-inline">Product</span> 4</a>
				//                             </li>
				//                         </ul>
				//                     </li>
				//                     <li>
				//                         <a href="#" class="nav-link px-0 align-middle">
				//                             <i class="fs-4 bi-people"></i> <span class="ms-1 d-none d-sm-inline">Customers</span> </a>
				//                     </li>
				//                 </ul>
				//                 <hr>
				//                 <div class="dropdown pb-4">
				//                     <a href="#" class="d-flex align-items-center text-white text-decoration-none dropdown-toggle" id="dropdownUser1" data-bs-toggle="dropdown" aria-expanded="false">
				//                         <img src="https://github.com/mdo.png" alt="hugenerd" width="30" height="30" class="rounded-circle">
				//                         <span class="d-none d-sm-inline mx-1">loser</span>
				//                     </a>
				//                     <ul class="dropdown-menu dropdown-menu-dark text-small shadow">
				//                         <li><a class="dropdown-item" href="#">New project...</a></li>
				//                         <li><a class="dropdown-item" href="#">Settings</a></li>
				//                         <li><a class="dropdown-item" href="#">Profile</a></li>
				//                         <li>
				//                             <hr class="dropdown-divider">
				//                         </li>
				//                         <li><a class="dropdown-item" href="#">Sign out</a></li>
				//                     </ul>
				//                 </div>
				//             </div>
				//         </div>
				//         <div class="col py-3">
				//             Content area...
				//         </div>
				//     </div>
				// </div>`),
				// 				gbt.RawHTML(`<div class="container text-center">
				//   <div class="row">
				//     <div class="col-sm-3">
				//       Level 1: .col-sm-3
				//     </div>
				//     <div class="col-sm-9">
				//       <div class="row">
				//         <div class="col-8 col-sm-6">
				//           Level 2: .col-8 .col-sm-6
				//         </div>
				//         <div class="col-4 col-sm-6">
				//           Level 2: .col-4 .col-sm-6
				//         </div>
				//       </div>
				//     </div>
				//   </div>
				// </div>`),
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
