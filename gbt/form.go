package gbt

import (
	"fmt"
	"html/template"
	"math/rand"
)

// Form ...
type Form struct {
	Items
	Submit *Button
	// TODO: form attributes
}

func (f *Form) Render() Stringer {
	if f.Submit != nil {
		f.Items = append(f.Items, f.Submit)
	}
	fr := &Element{
		Name:  "form",
		Items: f.Items,
	}
	return fr.Render()
}

// Fieldset
type Fieldset struct {
	Legend template.HTML
	Items
}

func (f *Fieldset) Render() Stringer {
	items := Items{}
	if f.Legend != "" {
		items = append(items, &Element{
			Name:      "legend",
			InnerHTML: f.Legend,
		})
	}
	items = append(items, f.Items...)
	fs := &Element{
		Name:  "fieldset",
		Items: items,
	}
	return fs.Render()
}

// FormLabel
type FormLabel struct {
	Label template.HTML
	ForID string
}

func (f *FormLabel) Render() Stringer {
	if f.Label == "" {
		return ""
	}
	l := &Element{
		Name: "label",
		Attributes: Attributes{
			"for": f.ForID,
		},
		Classes:   Classes{"form-label"},
		InnerHTML: f.Label,
	}
	return l.Render()
}

type Size string

var (
	SizeSmall Size = "sm"
	SizeLarge Size = "lg"
)

// FormControl ...
// Alias for Input
type FormControl struct {
	ID     string
	Size   Size // "", "sm", "lg" -> class="form-control form-control-lg"
	Type   InputType
	HelpID string
	Attributes
	Classes
}

func (f *FormControl) Render() Stringer {
	if f.Attributes == nil {
		f.Attributes = Attributes{}
	}
	f.Attributes["aria-describedby"] = f.HelpID

	f.Classes = append(f.Classes, "form-control")
	switch f.Size {
	case SizeSmall:
		f.Classes = append(f.Classes, "form-control-sm")
	case SizeLarge:
		f.Classes = append(f.Classes, "form-control-lg")
	}

	i := &Input{
		ID:         f.ID,
		Type:       f.Type,
		Classes:    f.Classes,
		Attributes: f.Attributes,
	}
	return i.Render()
}

// FormText
type FormText struct {
	ID   string
	Text template.HTML
}

func (f *FormText) Render() Stringer {
	if f.Text == "" {
		return ""
	}
	e := &Element{
		ID:        f.ID,
		Classes:   Classes{"form-text"},
		InnerHTML: f.Text,
	}
	return e.Render()
}

// FormCheck inline, switch, togglebtn (style), radio button(style),
type FormCheck struct{}

func (f *FormCheck) Render() Stringer {
	return ""
}

// FormCheckInput radio/checkbox
type FormCheckInput struct{}

func (f *FormCheckInput) Render() Stringer {
	return ""
}

// FormCheckLabel
type FormCheckLabel struct{}

func (f *FormCheckLabel) Render() Stringer {
	return ""
}

// FormSelect
type FormSelect struct{}

func (f *FormSelect) Render() Stringer {
	return ""
}

//
// custom form
//

// FormEmail ...
type FormEmail struct {
	ID          string
	Name        string
	Placeholder string
	Value       string
	Required    bool
	ReadOnly    bool
	Multiple    bool
	Label       template.HTML
	HelpText    template.HTML
}

func (f *FormEmail) Render() Stringer {
	if f.Name == "" {
		f.Name = "email" // possible conflict?
	}
	if f.ID == "" {
		f.ID = fmt.Sprintf("%s-%d", f.Name, rand.Intn(100))
	}

	items := Items{}
	if f.Label != "" {
		items = append(items, &FormLabel{
			Label: f.Label,
			ForID: f.ID,
		})
	}

	helpID := fmt.Sprintf("%s-help-%d", f.ID, rand.Intn(100))

	attributes := Attributes{}
	if f.Value != "" {
		attributes["value"] = f.Value
	}
	if f.Placeholder != "" {
		attributes["placeholder"] = f.Placeholder
	}
	if f.Required {
		attributes["required"] = f.Required
	}
	if f.ReadOnly {
		attributes["readonly"] = f.ReadOnly
	}
	if f.Multiple {
		attributes["multiple"] = f.Multiple
	}

	attributes["name"] = f.Name // should always exist

	items = append(items, &FormControl{
		ID:         f.ID,
		Type:       InputTypeEmail,
		HelpID:     helpID,
		Attributes: attributes,
	})

	if f.HelpText != "" {
		items = append(items, &FormText{
			ID:   helpID,
			Text: f.HelpText,
		})
	}

	d := &Element{
		Classes: Classes{"mb-3"},
		Items:   items,
	}
	return d.Render()
}

type AutoComplete string

var (
	AutoCompleteOn      AutoComplete = "on"
	AutoCompleteOff     AutoComplete = "off"
	AutoCompleteCurrent AutoComplete = "current-password"
	AutoCompleteNew     AutoComplete = "new-password"
)

type InputMode string

var (
	InputModeNone    InputMode = "none"
	InputModeText    InputMode = "text"
	InputModeDecimal InputMode = "decimal"
	InputModeNumeric InputMode = "numeric"
	InputModeTel     InputMode = "tel"
	InputModeSearch  InputMode = "search"
	InputModeEmail   InputMode = "email"
	InputModeURL     InputMode = "url"
)

// FormPassword ...
type FormPassword struct {
	ID          string
	Name        string
	Placeholder string
	Required    bool
	ReadOnly    bool
	MinLength   string // use string because it allows "" to omit empty
	MaxLength   string
	Size        string
	Pattern     string
	AutoComplete
	InputMode
	Label    template.HTML
	HelpText template.HTML
	Value    string // ??
}

func (f *FormPassword) Render() Stringer {
	if f.Name == "" {
		f.Name = "email" // possible conflict?
	}
	if f.ID == "" {
		f.ID = fmt.Sprintf("%s-%d", f.Name, rand.Intn(100))
	}

	items := Items{}
	if f.Label != "" {
		items = append(items, &FormLabel{
			Label: f.Label,
			ForID: f.ID,
		})
	}

	helpID := fmt.Sprintf("%s-help-%d", f.ID, rand.Intn(100))

	attributes := Attributes{}
	if f.Value != "" {
		attributes["value"] = f.Value
	}
	if f.Placeholder != "" {
		attributes["placeholder"] = f.Placeholder
	}
	if f.Required {
		attributes["required"] = f.Required
	}
	if f.ReadOnly {
		attributes["readonly"] = f.ReadOnly
	}
	if f.AutoComplete != "" {
		attributes["autocomplete"] = f.AutoComplete
	}
	if f.InputMode != "" {
		attributes["inputmode"] = f.InputMode
	}
	if f.Pattern != "" {
		attributes["pattern"] = f.Pattern
	}
	if f.MinLength != "" {
		attributes["minlength"] = f.MinLength
	}
	if f.MaxLength != "" {
		attributes["maxlength"] = f.MaxLength
	}
	if f.Size != "" {
		attributes["size"] = f.Size
	}

	attributes["name"] = f.Name // should always exist

	items = append(items, &FormControl{
		ID:         f.ID,
		Type:       InputTypeEmail,
		HelpID:     helpID,
		Attributes: attributes,
	})

	if f.HelpText != "" {
		items = append(items, &FormText{
			ID:   helpID,
			Text: f.HelpText,
		})
	}

	d := &Element{
		Classes: Classes{"mb-3"},
		Items:   items,
	}
	return d.Render()
}
