package ext

import (
	"fmt"
	"io"
)

var (
	bodyID = 0
)

// NewBody ...
func NewBody(items Items) *Body {
	return &Body{
		ID:    nextBodyID(),
		Items: items,
	}

}

// Body ...
type Body struct {
	ID    string
	Items Items
}

// Render ...
func (b *Body) Render(w io.Writer) error {
	if b.ID == "" {
		b.ID = nextInnerhtmlID()
	}

	div := &DivContainer{
		ID: b.ID,
		Classes: []string{
			"x-panel-body",
			"x-body-wrap-el",
			"x-panel-body-wrap-el",
			"x-container-body-wrap-el",
			"x-component-body-wrap-el",
		},
		Items: LayoutItems(b.Items),
	}
	return div.Render(w)
}

// GetID ...
func (b *Body) GetID() string {
	return b.ID
}

func nextBodyID() string {
	id := fmt.Sprintf("body-%d", bodyID)
	bodyID++
	return id
}
