package gbt

import (
	"fmt"
	"html/template"
)

type IconStyle string

var (
	IconOutlined IconStyle = "outlined"
	IconRounded  IconStyle = "rounded"
	IconSharp    IconStyle = "sharp"
)

// https://fonts.google.com/icons?icon.style=Outlined
// <span class="material-symbols-outlined">
// expand_more
// </span>
type Icon struct {
	Icon  string
	Style IconStyle
}

func (i *Icon) Render() Stringer {
	if i.Style == "" {
		i.Style = IconOutlined
	}

	class := fmt.Sprintf("material-symbols-%s", i.Style)

	ic := &Element{
		Name:       "span",
		Classes:    Classes{class},
		Attributes: Attributes{"style": "vertical-align: bottom;"},
		InnerHTML:  template.HTML(i.Icon),
	}
	return ic.Render()
}
