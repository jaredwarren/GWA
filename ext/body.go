package ext

import (
	"fmt"
	"html/template"
)

var (
	bodyID = 0
)

// NewBody ...
func NewBody(items []Renderer) *Body {
	return &Body{
		ID:    nextBodyID(),
		Items: items,
	}

}

// Body ...
type Body struct {
	ID    string
	Items []Renderer
}

// Render ...
func (b *Body) Render() template.HTML {
	if b.ID == "" {
		b.ID = nextBodyID()
	}

	return render("body", b)
}

// Debug ...
func (b *Body) Debug() {}

func nextBodyID() string {
	id := fmt.Sprintf("%d", bodyID)
	bodyID++
	return id
}
