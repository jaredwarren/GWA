package ext

import (
	"fmt"
	"io"
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
func (b *Body) Render(w io.Writer) error {
	nb, _ := b.Build()
	return render(w, "body", nb)
}

// Build copys info to a new panel
func (b *Body) Build() (Renderer, error) {
	n := &Body{}
	if b.ID != "" {
		n.ID = b.ID
	} else {
		n.ID = nextInnerhtmlID()
	}
	return n, nil
}

// Debug ...
func (b *Body) Debug() {}

func nextBodyID() string {
	id := fmt.Sprintf("%d", bodyID)
	bodyID++
	return id
}
