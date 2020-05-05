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
// TODO: I don't need all of this crap for container
type Body struct {
	ID        string
	Title     string
	IconClass string
	Layout    string
	HTML      template.HTML
	Width     int // float?
	Height    int // float?
	Items     []Renderer
	RenderTo  string // type???
	Body      *Body
	Border    template.CSS
	Docked    string // top, bottom, left, right
	Flex      int
	Style     string
	Classes   []string
}

// Render ...
func (b *Body) Render() template.HTML {
	if b.ID == "" {
		b.ID = nextBodyID()
	}

	// // default classes
	// b.Classes = []string{
	// 	"x-body-wrap-el",
	// 	"x-panel-body-wrap-el",
	// 	"x-container-body-wrap-el",
	// 	"x-component-body-wrap-e",
	// }

	return render("body", b)
}

func nextBodyID() string {
	id := fmt.Sprintf("%d", bodyID)
	bodyID++
	return id
}
