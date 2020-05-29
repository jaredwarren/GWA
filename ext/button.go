package ext

import (
	"fmt"
	"html/template"
	"io"
)

var (
	buttonID = 0
)

// Button ...
type Button struct {
	XType     string        `json:"xtype"`
	ID        string        `json:"id,omitempty"`
	Text      template.HTML `json:"text,omitempty"`
	Handler   template.JS   `json:"handler,omitempty"`
	UI        string        `json:"ui,omitempty"` // TODO
	IconClass string        `json:"iconClass,omitempty"`
	Classes   Classes       `json:"classes,omitempty"`
	Styles    Styles        `json:"styles,omitempty"`
	Parent    Renderer      `json:"-"`
	HandlerFn Handler       `json:"-"`
}

// Render ...
func (b *Button) Render(w io.Writer) error {
	if b.ID == "" {
		b.ID = nextButtonID()
	}

	b.Styles = Styles{}

	if b.HandlerFn != nil {
		// TODO: fix id: remove '-'
		name := fmt.Sprintf("%s_click", "todo_")
		b.Handler = template.JS(name)
		go ui.Bind(name, b.HandlerFn)
	}

	// default classes
	b.Classes.Add("btn")

	// UI
	if b.UI == "" {
		b.UI = "primary"
	}
	b.Classes.Add(fmt.Sprintf("btn-%s", b.UI))

	// Attributes
	attrs := map[string]template.HTMLAttr{
		"id":    template.HTMLAttr(b.ID),
		"class": b.Classes.ToAttr(),
	}

	// TODO: add property for icon position (top: flex-direction column, ...)
	b.Styles["flex-direction"] = "row"
	b.Styles["display"] = "flex"

	if len(b.Styles) > 0 {
		attrs["style"] = b.Styles.ToAttr()
	}

	// Handler
	if b.Handler != "" {
		attrs["onclick"] = template.HTMLAttr(fmt.Sprintf("%s('%s')", b.Handler, b.ID))
	}

	items := Items{}

	// IconClass
	if b.IconClass != "" {
		items = append(items, &Element{
			Name: "i",
			Attributes: Attributes{
				"class": template.HTMLAttr(b.IconClass),
			},
		})
	}

	// Text
	if b.Text != "" {
		items = append(items, &RawHTML{template.HTML(b.Text)})
	}

	buttonEl := &Element{
		Name:       "button",
		Attributes: attrs,
		Items:      items,
	}

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
