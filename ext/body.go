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
	fmt.Print("  render Body:")
	n := &Body{}
	if b.ID != "" {
		n.ID = b.ID
	} else {
		n.ID = nextInnerhtmlID()
	}
	fmt.Println(n.ID) // show id
	// n.Items = b.Items
	n.Items = LayoutItems(b.Items)
	return render(w, "body", n)
}

func nextBodyID() string {
	id := fmt.Sprintf("%d", bodyID)
	bodyID++
	return id
}
