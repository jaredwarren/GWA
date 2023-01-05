package ext

import (
	"fmt"
	"html/template"
	"io"
)

var (
	formID  = 0
	inputID = 0
)

// Form ...
type Form struct {
	XType     string            `json:"xtype"`
	ID        string            `json:"id,omitempty"`
	Parent    Renderer          `json:"-"`
	Width     int               `json:"width,omitempty"`
	Height    int               `json:"height,omitempty"`
	Layout    string            `json:"layout,omitempty"`
	Resizable bool              `json:"resizable,omitempty"`
	Resize    string            `json:"resize,omitempty"`
	Border    template.CSS      `json:"border,omitempty"`
	Docked    string            `json:"docked,omitempty"` // top, bottom, left, right, ''
	Flex      string            `json:"flex,omitempty"`
	Shadow    bool              `json:"shadow,omitempty"`
	Classes   Classes           `json:"classes,omitempty"`
	Styles    map[string]string `json:"styles,omitempty"`
	Items     Items             `json:"items,omitempty"`
	Action    string            `json:"action,omitempty"`
	Method    string            `json:"method,omitempty"`
	Handler   string            `json:"handler,omitempty"`
	HandlerFn FormHandler       `json:"-"`
	// TODO: success/fail handler, how to push info back to front?
}

func (f *Form) RenderString() string {
	return "TODO: Application"
}

// Render ...
func (f *Form) Render(w io.Writer) error {
	if f.ID == "" {
		f.ID = nextFormID()
	}

	// TODO: find a way to make form use same logic as panel, so there isn't duplicate code

	// default to post method
	if f.Method == "" {
		f.Method = "post"
	}

	// default action
	if f.Action == "" {
		f.Action = fmt.Sprintf("/submit/%s", f.Handler)
	}

	// default classes
	classess := map[string]bool{
		"x-form": true,
	}
	// copy classes
	for _, c := range f.Classes {
		if _, ok := classess[c]; !ok {
			classess[c] = true
		}
	}
	if f.Shadow {
		classess["x-shadow"] = true
	}

	// copy styles
	styles := Styles{}
	if len(f.Styles) > 0 {
		styles = f.Styles
	}

	// append new styles based on p's properties
	// what if I want width to be 0px?
	if f.Width != 0 && f.Docked != "top" && f.Docked != "bottom" {
		styles["width"] = fmt.Sprintf("%dpx", f.Width)
		classess["x-widthed"] = true
	}
	// what if I want height to be 0px?
	if f.Height != 0 && f.Docked != "left" && f.Docked != "right" {
		styles["height"] = fmt.Sprintf("%dpx", f.Height)
		classess["x-heighted"] = true
	}
	if f.Border != "" {
		styles["border"] = string(f.Border)
		classess["x-managed-border"] = true
	}

	if f.Layout == "hbox" {
		styles["flex-direction"] = "row"
	} else {
		styles["flex-direction"] = "column"
	}

	if f.Flex != "" {
		styles["flex"] = f.Flex
	}

	if f.Resizable {
		styles["resize"] = "both"
	} else if f.Resize != "" {
		styles["resize"] = f.Resize
	}

	// convert class back to array
	npClasses := []string{}
	for k := range classess {
		npClasses = append(npClasses, k)
	}

	// ITEMS
	items := Items{}
	for _, i := range f.Items {
		c, ok := i.(Child)
		if ok {
			c.SetParent(f)
		}
		items = append(items, i)
	}

	// Attributes
	attrs := map[string]template.HTMLAttr{
		"id":       template.HTMLAttr(f.ID),
		"class":    f.Classes.ToAttr(),
		"action":   template.HTMLAttr(f.Action),
		"method":   template.HTMLAttr(f.Method),
		"onsubmit": template.HTMLAttr(fmt.Sprintf("submitForm('%s', event)", f.ID)),
	}
	if len(styles) > 0 {
		attrs["style"] = styles.ToAttr()
	}

	navEl := &Element{
		Name:       "form",
		Attributes: attrs,
		Items:      LayoutItems(items),
	}
	return navEl.Render(w)
}

// GetID ...
func (f *Form) GetID() string {
	return f.ID
}

// SetParent ...
func (f *Form) SetParent(p Renderer) {
	f.Parent = p
}

// GetDocked ...
func (f *Form) GetDocked() string {
	return f.Docked
}

// SetStyle ...
func (f *Form) SetStyle(key, value string) {
	if f.Styles == nil {
		f.Styles = map[string]string{}
	}
	f.Styles[key] = value
}

func nextFormID() string {
	id := fmt.Sprintf("form-%d", formID)
	formID++
	return id
}

/*
 * Fieldset
 */

// Fieldset ...
type Fieldset struct {
	XType     string            `json:"xtype"`
	ID        string            `json:"id,omitempty"`
	Docked    string            `json:"docked,omitempty"`
	Classes   []string          `json:"classes,omitempty"`
	Styles    map[string]string `json:"styles,omitempty"`
	Resizable bool              `json:"resizable,omitempty"`
	Resize    string            `json:"resize,omitempty"`
	Legend    template.HTML     `json:"legend,omitempty"`
	Items     Items             `json:"items,omitempty"`
	Parent    Renderer          `json:"-"`
}

func (f *Fieldset) RenderString() string {
	return "TODO: Application"
}

// Render ...
func (f *Fieldset) Render(w io.Writer) error {
	items := Items{}
	if f.Legend != "" {
		items = append(items, &Element{
			Name:  "legend",
			Items: Items{&RawHTML{f.Legend}},
		})
	}

	// copy styles
	styles := Styles{}
	if len(f.Styles) > 0 {
		styles = f.Styles
	}
	if f.Resizable {
		styles["resize"] = "both"
		styles["overflow"] = "auto" // needs to be anything but "visible"
	} else if f.Resize != "" {
		styles["resize"] = f.Resize
		styles["overflow"] = "auto"
	}

	for _, i := range f.Items {
		c, ok := i.(Child)
		if ok {
			c.SetParent(f)
		}
		items = append(items, i)
	}

	navEl := &Element{
		Name: "fieldset",
		Attributes: Attributes{
			"style": styles.ToAttr(),
		},
		Items: LayoutItems(items),
	}
	return navEl.Render(w)
}

// GetID ...
func (f *Fieldset) GetID() string {
	return f.ID
}

/*
 * Input
 */

// Input ...
type Input struct {
	XType        string        `json:"xtype"`
	ID           string        `json:"id,omitempty"`
	Classes      Classes       `json:"classes,omitempty"`
	Styles       Styles        `json:"styles,omitempty"`
	Type         string        `json:"type,omitempty"`
	Name         string        `json:"name,omitempty"`
	Value        string        `json:"value,omitempty"`
	Attributes   Attributes    `json:"attributes,omitempty"`
	Events       Events        `json:"events,omitempty"`
	Data         Data          `json:"data,omitempty"`
	Form         string        `json:"form,omitempty"`
	Disabled     bool          `json:"disabled,omitempty"`
	Autofocus    bool          `json:"autofocus,omitempty"`
	Autocomplete string        `json:"autocomplete,omitempty"`
	Label        template.HTML `json:"label,omitempty"`
	Parent       Renderer      `json:"-"`
}

func (i *Input) RenderString() string {
	return "TODO: Application"
}

// Render ...
func (i *Input) Render(w io.Writer) error {
	if i.ID == "" {
		i.ID = nextInputID()
	}

	// TODO: validate attributes based on type

	if len(i.Attributes) == 0 {
		i.Attributes = map[string]template.HTMLAttr{}
	}

	// ID
	i.Attributes["id"] = template.HTMLAttr(i.ID)

	// Type
	if i.Type == "" {
		// required, default text??
		panic("type required")
	}
	i.Attributes["type"] = template.HTMLAttr(i.Type)

	// Name
	if i.Name != "" {
		i.Attributes["name"] = template.HTMLAttr(i.Name)
	}

	// Value
	if i.Value != "" {
		i.Attributes["value"] = template.HTMLAttr(i.Value)
	}

	// Disabled
	if i.Disabled {
		i.Attributes["disabled"] = ""
	}

	// Autofocus
	if i.Autofocus {
		i.Attributes["autofocus"] = ""
	}

	// Autocomplete
	if i.Autocomplete != "" {
		i.Attributes["autocomplete"] = template.HTMLAttr(i.Autocomplete)
	}

	// Form
	if i.Form != "" {
		i.Attributes["form"] = template.HTMLAttr(i.Form)
	} // TODO: else get from parent form?

	// Add Events
	eattr := i.Events.ToAttr()
	for n, e := range eattr {
		i.Attributes[n] = e
	}

	// Add Data
	dAttrs := i.Data.ToAttr()
	for n, v := range dAttrs {
		i.Attributes[n] = v
	}

	// Label
	items := Items{}
	if i.Label != "" {
		items = append(items, &Element{
			Name:       "label",
			Attributes: Attributes{"for": template.HTMLAttr(i.ID)},
			Innerhtml:  i.Label,
		})
	}

	e := &Element{
		Name:       "input",
		Attributes: i.Attributes,
	}

	// override name for textarea
	if i.Type == "textarea" {
		e.Name = "textarea"
	}

	items = append(items, e)
	return items.Render(w)
}

func nextInputID() string {
	id := fmt.Sprintf("input-%d", inputID)
	inputID++
	return id
}

// GetID ...
func (i *Input) GetID() string {
	return i.ID
}
