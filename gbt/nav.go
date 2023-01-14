package gbt

import (
	"fmt"
	"html/template"
	"math/rand"
)

// NavItem belongs to Nav
type NavItem struct {
	Image    *Image
	Title    string
	Href     string
	Disabled bool
	Active   bool
}

func (n *NavItem) Render() Stringer {
	attributes := Attributes{"href": n.Href}
	classes := Classes{"nav-link"}
	if n.Active {
		classes = append(classes, "active")
		attributes["aria-current"] = "page"
	}
	if n.Disabled {
		classes = append(classes, "disabled")
	}
	el := &Element{
		Name:    "li",
		Classes: Classes{"nav-item"},
		Items: Items{&Element{
			Name:       "a",
			Classes:    classes,
			Attributes: attributes,
			InnerHTML:  n.Title,
		}},
	}
	return el.Render()
}

// NavDropDown Navitem, but with dropdown items
type NavDropDown struct {
	Title string
	Items
}

func (n *NavDropDown) Render() Stringer {
	el := &Element{
		Name:    "li",
		Classes: Classes{"nav-item", "dropdown"},
		Items: Items{
			&Element{
				Name:    "a",
				Classes: Classes{"nav-link", "dropdown-toggle"},
				Attributes: Attributes{
					"href":           "#",
					"role":           "button",
					"data-bs-toggle": "dropdown",
					"aria-expanded":  "false",
				},
				InnerHTML: n.Title,
			},
			&Element{
				Name:    "ul",
				Classes: Classes{"dropdown-menu"},
				Items:   n.Items,
			},
		},
	}
	return el.Render()
}

// DropDownItem belongs to NavDropDown
type DropDownItem struct {
	Title    string
	Href     string
	Disabled bool
	Active   bool
}

func (n *DropDownItem) Render() Stringer {
	classes := Classes{"dropdown-item"}
	if n.Active {
		classes = append(classes, "active")
	}
	if n.Disabled {
		classes = append(classes, "disabled")
	}
	el := &Element{
		Name: "li",
		Items: Items{&Element{
			Name:       "a",
			Classes:    classes,
			Attributes: Attributes{"href": n.Href},
			InnerHTML:  n.Title,
		}},
	}
	return el.Render()
}

// DropDowndivider belongs to NavDropDown
type DropDowndivider struct {
	Title string
	Href  string
}

func (n *DropDowndivider) Render() Stringer {
	el := &Element{
		Name: "li",
		Items: Items{&Element{
			Name:    "hr",
			Classes: Classes{"dropdown-divider"},
		}},
	}
	return el.Render()
}

// NavBrand nav band
type NavBrand struct {
	Image *Image
	Title string
	Href  string
}

type Option[T any] func(*T)

// Verbosity sets Foo's verbosity level to v.
// func NavTitle[P Option](v string) P {
// 	return func(f *NavBrand) {
// 		f.Title = v
// 	}
// }

func NavTitle(title string) Option[NavBrand] {
	return navTitle[NavBrand](title)
}

func navTitle[T any](title string) Option[T] {
	return func(a *T) {
		switch x := any(a).(type) {
		case *NavBrand:
			x.Title = title
		}
	}
}

func NavImage(image string) Option[NavBrand] {
	return navImage[NavBrand](image)
}

func navImage[T any](image string) Option[T] {
	return func(a *T) {
		switch x := any(a).(type) {
		case *NavBrand:
			x.Image = &Image{
				Src:    image,
				Height: "20px",
			}
		}
	}
}

type Imager interface {
	SetImage(i *Image)
}

// func NavImage[T Option](v string) T {
// 	return func(f Imager) {
// 		f.SetImage(&Image{
// 			Src:    v,
// 			Height: "20px",
// 		})
// 	}
// }

// func OptionTemperature(t Celsius) func(*Foobar) error {
// 	return func(fb *Foobar) error {
// 		fb.temperature = t
// 		return nil
// 	}
// }
// func Append[T any](s []T, t ...T) []T {
func NewBrand(opts ...Option[NavBrand]) *NavBrand {
	fb := &NavBrand{
		Title: "{{.Nav.Title}}",
		Href:  "#",
	}
	for _, op := range opts {
		op(fb)
	}
	return fb
}

func (n *NavBrand) SetImage(i *Image) {
	n.Image = i
}

func (n *NavBrand) Render() Stringer {
	var el *Element
	if n.Href == "" {
		el = &Element{
			Name:    "span",
			Classes: Classes{"navbar-brand", "mb-0", "h1"},
		}
	} else {
		el = &Element{
			Name:       "a",
			Classes:    Classes{"navbar-brand"},
			Attributes: Attributes{"href": n.Href},
		}
	}
	el.Items = Items{}
	if n.Image != nil {
		el.Items = append(el.Items, n.Image)
	}
	el.Items = append(el.Items, RawHTML(n.Title))
	return el.Render()
}

// Nav ...
type Nav struct {
	ID string `json:"id,omitempty"` // how to auto generate
	// Title (optional) overwritten if Brand is set
	Title  template.HTML `json:"title,omitempty"`
	Brand  *NavBrand
	Items  Items `json:"items,omitempty"`
	Theme  Theme
	Search bool
	//
	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`

	Border  template.CSS      `json:"border,omitempty"`
	Style   string            `json:"style,omitempty"`
	Docked  string            `json:"docked,omitempty"` // top, bottom, left, right
	Classes []string          `json:"classes,omitempty"`
	Styles  map[string]string `json:"styles,omitempty"`
	Shadow  bool              `json:"shadow,omitempty"`
	//
	XType  string        `json:"xtype"`
	Layout string        `json:"layout,omitempty"`
	HTML   template.HTML `json:"html,omitempty"`
}

// Render ...
func (n *Nav) Render() Stringer {
	if n.Theme == "" {
		n.Theme = ThemeLight
	}

	if n.ID == "" {
		n.ID = fmt.Sprintf("navbarSupportedContent%d", rand.Intn(100))
	}

	var searchEl *Element
	if n.Search {
		searchEl = &Element{
			Name:       "form",
			Classes:    Classes{"d-flex"},
			Attributes: Attributes{"role": "search"},
			Items: Items{
				&FormControl{},
				&Button{
					Outline: true,
					Style:   ButtonSuccess,
					Type:    ButtonTypeSubmit,
					Text:    "Search",
				},
			},
		}
	}

	el := &Element{
		Name:       "nav",
		Classes:    Classes{"navbar", "navbar-expand-lg", "bg-body-tertiary"},
		Attributes: Attributes{"data-bs-theme": n.Theme},
		Items: Items{&Element{
			Classes: Classes{"container-fluid"},
			Items: Items{
				n.Brand,
				&Button{
					Classes: Classes{"navbar-toggler"},
					Items: Items{&Element{
						Name:    "span",
						Classes: Classes{"navbar-toggler-icon"},
						Attributes: Attributes{
							"type":           "button",
							"data-bs-toggle": "collapse",
							"data-bs-target": fmt.Sprintf("#%s", n.ID),
							"aria-controls":  n.ID,
							"aria-expanded":  "false",
							"aria-label":     "Toggle navigation",
						},
					}},
				},
				&Element{
					ID:      n.ID,
					Classes: Classes{"collapse", "navbar-collapse"},
					Items: Items{
						&Element{
							Name:    "ul",
							Classes: Classes{"navbar-nav", "me-auto", "mb-2", "mb-lg-0"},
							Items:   n.Items,
						},
						searchEl,
					},
				},
			},
		}},
	}
	return el.Render()
}
