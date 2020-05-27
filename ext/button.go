package ext

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

var (
	buttonID = 0
)

// Button ...
type Button struct {
	XType     string            `json:"xtype"`
	ID        string            `json:"id,omitempty"`
	Text      template.HTML     `json:"text,omitempty"`
	Handler   template.JS       `json:"handler,omitempty"`
	UI        string            `json:"ui,omitempty"` // TODO
	IconClass string            `json:"iconClass,omitempty"`
	Classes   []string          `json:"classes,omitempty"`
	Styles    map[string]string `json:"styles,omitempty"`
	Parent    Renderer          `json:"-"`
	HandlerFn Handler           `json:"-"`
}

// Render ...
func (b *Button) Render(w io.Writer) error {
	if b.ID == "" {
		b.ID = nextButtonID()
	}

	if b.UI == "" {
		b.UI = "primary"
	}

	if b.HandlerFn != nil {
		// TODO: fix id: remove '-'
		name := fmt.Sprintf("%s_click", "todo_")
		b.Handler = template.JS(name)
		go ui.Bind(name, b.HandlerFn)
	}

	// default classes
	classess := map[string]bool{
		"btn": true,
	}
	if b.UI != "" {
		bui := fmt.Sprintf("btn-%s", b.UI)
		classess[bui] = true
	}
	// copy classes
	for _, c := range b.Classes {
		if _, ok := classess[c]; !ok {
			classess[c] = true
		}
	}
	// convert class back to array
	npClasses := []string{}
	for k := range classess {
		npClasses = append(npClasses, k)
	}

	// Attributes
	attrs := map[string]template.HTMLAttr{
		"id":    template.HTMLAttr(b.ID),
		"class": template.HTMLAttr(strings.Join(npClasses, " ")),
	}

	// Handler
	if b.Handler != "" {
		attrs["onclick"] = template.HTMLAttr(fmt.Sprintf("%s('%s')", b.Handler, b.ID))
	}

	buttonEl := &Element{
		Name:       "button",
		Attributes: attrs,
	}

	html := template.HTML("")

	// IconClass
	// TODO: add icon position
	if b.IconClass != "" {
		// TODO: convert iconClass and Text to Items
		html = template.HTML(fmt.Sprintf("<i class=%q></i>", b.IconClass))
	}

	// Text
	if b.Text != "" {
		html += template.HTML(b.Text)
	}

	buttonEl.Items = Items{&RawHTML{html}}

	return buttonEl.Render(w)
}

// GetID ...
func (b *Button) GetID() string {
	return b.ID
}

// SetParent ...
func (b *Button) SetParent(p Renderer) {
	b.Parent = p
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
