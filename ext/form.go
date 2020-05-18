package ext

import (
	"fmt"
	"io"
)

var (
	formID = 0
)

// Form ...
type Form struct {
	ID       string
	ShowRoot bool
	Parent   Renderer
	Root     *FormNode
	Docked   string
	Classes  []string
	Styles   map[string]string
}

// Render ...
func (t *Form) Render(w io.Writer) error {
	if t.ID == "" {
		t.ID = nextFormID()
	}
	if t.Styles == nil {
		t.Styles = map[string]string{}
	}
	t.Classes = append(t.Classes, "x-form")
	t.Styles["border"] = "1px solid lightgrey"
	return renderTemplate(w, "form", t)
}

// GetID ...
func (t *Form) GetID() string {
	return t.ID
}

// SetParent ...
func (t *Form) SetParent(p Renderer) {
	t.Parent = p
}

// GetDocked ...
func (t *Form) GetDocked() string {
	return t.Docked
}

// FormNode ...
type FormNode struct {
	ID        string
	Text      string
	Expanded  bool
	Leaf      bool
	IconClass string
	Children  []*FormNode
}

// Render ...
func (tn *FormNode) Render(w io.Writer) error {
	// if tn.ID == "" {
	// 	tn.ID = nextFormID()
	// }
	// copy styles
	// styles := map[string]string{}
	// if len(t.Styles) > 0 {
	// 	styles = t.Styles
	// }

	return renderTemplate(w, "formnode", tn)
}

// GetID ...
func (tn *FormNode) GetID() string {
	return tn.ID
}

func nextFormID() string {
	id := fmt.Sprintf("form-%d", formID)
	formID++
	return id
}
