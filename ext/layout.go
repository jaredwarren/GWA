package ext

import (
	"fmt"
	"html/template"
)

var (
	layoutID = 0
)

// NewLayout ...
func NewLayout() *Layout {
	return &Layout{
		ID: nextLayoutID(),
		// Width:  300,
		// Height: 200,
		// Border: template.CSS("1px solid lightgrey"),
	}
}

// ItemList ...
type ItemList []Renderer

// Layout ...
type Layout struct {
	ID      string // how to auto generate
	Type    string // absolute, accordion, border, card, tab, hbox, vbox
	Pack    string // start, end, center, space-between, space-arround, justify
	Align   string // start, end, center, stretch
	Items   []Renderer
	Classes []string
	Styles  map[string]string
}

// Render ...
func (l *Layout) Render() template.HTML {
	fmt.Println("  render layout")

	if l.ID == "" {
		l.ID = nextLayoutID()
	}

	// default classes
	// TODO: add class based on type, pack, align
	l.Classes = []string{
		"x-dock",
		"x-dock-vertical",
		"x-managed-border",
		// "x-layout",
		// "x-layout",
		// "x-component",
		// "x-noborder-trbl",
		// "x-header-position-top",
		// "x-root",
	}

	// TODO: if pack is end flex-end?

	// // HTML
	// if l.HTML != "" {
	// 	l.Items = append(l.Items, &Layout{
	// 		HTML: l.HTML,
	// 	})
	// }

	// // ITEMS
	// if len(l.Items) == 0 {
	// 	// debug add dummy html
	// 	html := l.HTML
	// 	if html == "" {
	// 		html = ":("
	// 	}
	// 	l.Items = []Renderer{
	// 		&Layout{
	// 			HTML: html,
	// 		},
	// 	}
	// }

	// // BODY
	// if l.Body == nil {
	// 	l.Body = NewBody(l.Items)
	// }

	styles := map[string]string{}
	if len(l.Styles) > 0 {
		styles = l.Styles
	}
	if l.Type == "hbox" {
		styles["display"] = "flex"
		if l.Pack == "start" {
			styles["flex-direction"] = "column"
		} else if l.Pack == "end" {
			styles["flex-direction"] = "column-reverse"
		}
	}

	if l.Type == "vbox" {
		styles["display"] = "flex"
		if l.Pack == "start" {
			styles["flex-direction"] = "row"
		} else if l.Pack == "end" {
			styles["flex-direction"] = "row-reverse"
		}
	}

	// TODO: might have to check all items, if docked?

	// TODO: parse l.Style and add to styles
	l.Styles = styles

	// body
	// x-layout-body x-body-wrap-el x-layout-body-wrap-el x-layout-body-wrap-el x-component-body-wrap-el
	// x-body-wrap-el x-layout-body-wrap-el x-layout-body-wrap-el x-component-body-wrap-el

	return render("layout", l)
}

// Debug ...
func (l *Layout) Debug() {

}

func nextLayoutID() string {
	id := fmt.Sprintf("%d", layoutID)
	layoutID++
	return id
}
