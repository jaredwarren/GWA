package gbt

import (
	"fmt"
	"html/template"
)

type ButtonStyle string

var (
	ButtonPrimary   ButtonStyle = "primary"
	ButtonSecondary ButtonStyle = "secondary"
	ButtonSuccess   ButtonStyle = "success"
	ButtonDanger    ButtonStyle = "danger"
	ButtonWarning   ButtonStyle = "warning"
	ButtonInfo      ButtonStyle = "info"
	ButtonLight     ButtonStyle = "light"
	ButtonDark      ButtonStyle = "dark"
	ButtonLink      ButtonStyle = "link"
)

type ButtonType string

var (
	ButtonTypeSubmit ButtonType = "submit"
	ButtonTypeReset  ButtonType = "reset"
	ButtonTypeButton ButtonType = "button"
)

// Button ...
type Button struct {
	ID           string
	Text         template.HTML
	Handler      template.JS
	Icon         *Icon
	IconPosition string
	Style        ButtonStyle
	Outline      bool
	Disabled     bool
	Href         string
	Type         string  //
	OnClick      string  // TODO:
	HandlerFn    Handler // TODO:
	Badge        *Badge
	Size
	Classes
	Attributes
}

// Render ...
func (b *Button) Render() Stringer {
	if b.Attributes == nil {
		b.Attributes = Attributes{}
	}

	if b.Style == "" {
		b.Style = ButtonPrimary
	}

	name := "button"
	if b.Href != "" {
		name = "a"
		b.Attributes["href"] = b.Href
	}

	// Icon and Text
	items := Items{}
	if b.Icon != nil {
		if b.IconPosition == "right" {
			items = append(items, &Element{
				Name:      "span",
				InnerHTML: template.HTML(b.Text),
			}, b.Icon)
		} else {
			items = append(items, b.Icon, &Element{
				Name:      "span",
				InnerHTML: template.HTML(b.Text),
			})
		}
	} else {
		items = append(items, &Element{
			Name:      "span",
			InnerHTML: template.HTML(b.Text),
		})
	}

	// Badge
	if b.Badge != nil {
		items = append(items, b.Badge)
	}

	// Style
	b.Classes = append(b.Classes, "btn")
	style := b.Style
	if b.Outline {
		style = "btn-outline-" + style
	} else {
		style = "btn-" + style
	}
	b.Classes = append(b.Classes, string(style))

	// size
	if b.Size != "" {
		b.Classes = append(b.Classes, fmt.Sprintf("btn-%s", b.Size))
	}

	// Type
	if b.Type != "" {
		b.Attributes["type"] = b.Type
	}

	btn := &Element{
		Name:       name,
		Attributes: b.Attributes,
		Classes:    b.Classes,
		Items:      items,
	}
	return btn.Render()

}

// Badge ...
type Badge struct {
	Text  template.HTML
	Style ButtonStyle
}

func (b *Badge) Render() Stringer {
	if b.Style == "" {
		b.Style = ButtonSecondary
	}

	classes := Classes{
		"badge",
		fmt.Sprintf("text-bg-%s", b.Style),
	}

	ic := &Element{
		Name:      "span",
		Classes:   classes,
		InnerHTML: b.Text,
	}
	return ic.Render()
}
