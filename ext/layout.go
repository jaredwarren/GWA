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
	}
}

// Layout ...
type Layout struct {
	ID      string // how to auto generate
	Type    string // absolute, accordion, border, card, tab, hbox, vbox
	Pack    string // start, end, center, space-between, space-arround, justify
	Align   string // start, end, center
	Reverse bool   // reverse direction
	Items   Items
	Classes []string
	Styles  map[string]string
}

// Render ...
func (l *Layout) Render(w io.Writer) error {
	if l.ID == "" {
		l.ID = nextInnerhtmlID()
	}

	// default classes
	classess := map[string]bool{
		"x-layout": true,
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
		if l.Reverse {
			styles["flex-direction"] = "column-reverse"
		} else {
			styles["flex-direction"] = "column"
		}
		classess["x-hbox"] = true
	}

	if l.Type == "vbox" {
		styles["display"] = "flex"

		if l.Reverse {
			styles["flex-direction"] = "row-reverse"
		} else {
			styles["flex-direction"] = "row"
		}
		classess["x-vbox"] = true
	}

	if l.Pack != "" {
		styles["justify-content"] = l.Pack
	} // default?

	switch l.Align {
	case "start":
		styles["align-items"] = "flex-start"
	case "end":
		styles["align-items"] = "flex-end"
	case "center":
		styles["align-items"] = "center"
	default:
		styles["align-items"] = "stretch"
	}

	nClasses := []string{}
	for k := range classess {
		nClasses = append(nClasses, k)
	}

	div := &DivContainer{
		ID:      l.ID,
		Styles:  styles,
		Classes: nClasses,
		Items:   l.Items, // don't layout again just copy items
	}
	return renderDiv(w, div)
}

// GetID ...
func (l *Layout) GetID() string {
	return l.ID
}

func nextLayoutID() string {
	id := fmt.Sprintf("layout-%d", layoutID)
	layoutID++
	return id
}
