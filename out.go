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
}