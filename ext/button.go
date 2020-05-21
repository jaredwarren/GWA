package ext

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
)

var (
	buttonID = 0
)

// Button ...
type Button struct {
	ID        string
	Text      template.HTML
	Handler   template.JS
	UI        string // TODO
	IconClass string
	Parent    Renderer
	HandlerFn Handler
}

// Render ...
func (b *Button) Render(w io.Writer) error {
	if b.ID == "" {
		b.ID = nextButtonID()
	}

	if b.HandlerFn != nil {
		// TODO: fix id: remove '-'
		name := fmt.Sprintf("%s_click", "todo_")
		b.Handler = template.JS(name)
		go ui.Bind(name, b.HandlerFn)
	}

	return renderTemplate(w, "button", b)
}

// GetID ...
func (b *Button) GetID() string {
	return b.ID
}

// SetParent ...
func (b *Button) SetParent(p Renderer) {
	b.Parent = p
}

// MarshalJSON ...
func (b *Button) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		XType     string        `json:"xtype"`
		ID        string        `json:"id,omitempty"`
		Text      template.HTML `json:"text,omitempty"`
		Handler   template.JS   `json:"handler,omitempty"`
		UI        string        `json:"ui,omitempty"`
		IconClass string        `json:"iconClass,omitempty"`
	}{
		XType:     "button",
		ID:        b.ID,
		Text:      b.Text,
		Handler:   b.Handler,
		UI:        b.UI,
		IconClass: b.IconClass,
	})
}

func nextButtonID() string {
	id := fmt.Sprintf("button-%d", buttonID)
	buttonID++
	return id
}

func buildButton(i interface{}) *Button {
	ii := i.(map[string]interface{})

	p := &Button{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if IconClass, ok := ii["iconClass"]; ok {
		p.IconClass = IconClass.(string)
	}

	if ui, ok := ii["ui"]; ok {
		p.UI = ui.(string)
	}

	if text, ok := ii["text"]; ok {
		p.Text = template.HTML(text.(string))
	}

	if handler, ok := ii["handler"]; ok {
		p.Handler = template.JS(handler.(string))
	}

	return p
}
