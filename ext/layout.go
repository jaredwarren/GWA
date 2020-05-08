package ext

import (
	"fmt"
	"io"
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
func (l *Layout) Render(w io.Writer) error {
	fmt.Println("  render layout")
	nl, _ := l.Build()
	return render(w, "layout", nl)
}

// Build copys info to a new panel
func (l *Layout) Build() (Renderer, error) {
	n := &Layout{}
	if l.ID != "" {
		n.ID = l.ID
	} else {
		n.ID = nextInnerhtmlID()
	}

	// default classes
	// TODO: add class based on type, pack, align
	classess := map[string]bool{
		"x-dock":           true,
		"x-dock-vertical":  true,
		"x-managed-border": true,
	}
	// add classses from p, no duplicates
	for _, c := range l.Classes {
		if _, ok := classess[c]; !ok {
			classess[c] = true
		}
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

	// copy styles from og
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

		// TODO: fix width asnd height!!!!
		classess["x-hbox"] = true
		// styles["width"] = "100%" // or something like this!!
	}

	if l.Type == "vbox" {
		styles["display"] = "flex"
		if l.Pack == "start" {
			styles["flex-direction"] = "row"
		} else if l.Pack == "end" {
			styles["flex-direction"] = "row-reverse"
		}

		classess["x-vbox"] = true
		// .Styles["height"] = "100%" // or something like this!!
	}
	n.Styles = styles

	n.Classes = []string{}
	for k := range classess {
		n.Classes = append(n.Classes, k)
	}

	// TODO: might have to check all items, if docked?

	// TODO: parse l.Style and add to styles
	// l.Styles = styles

	return n, nil
}

// Debug ...
func (l *Layout) Debug() {

}

func nextLayoutID() string {
	id := fmt.Sprintf("%d", layoutID)
	layoutID++
	return id
}
