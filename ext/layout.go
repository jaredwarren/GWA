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
type ItemList Items

// Layout ...
type Layout struct {
	ID      string // how to auto generate
	Type    string // absolute, accordion, border, card, tab, hbox, vbox
	Pack    string // start, end, center, space-between, space-arround, justify
	Align   string // start, end, center, stretch
	Items   Items
	Classes []string
	Styles  map[string]string
}

// Render ...
func (l *Layout) Render(w io.Writer) error {
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

		// TODO: fix width and height!!!!
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

	// don't layout again just copy items
	n.Items = l.Items

	return render(w, "layout", n)
}

func nextLayoutID() string {
	id := fmt.Sprintf("%d", layoutID)
	layoutID++
	return id
}
