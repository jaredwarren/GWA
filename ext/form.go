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
	Docked    string            `json:"docked,omitempty"`
	Classes   []string          `json:"classes,omitempty"`
	Styles    map[string]string `json:"styles,omitempty"`
	Items     Items             `json:"items,omitempty"`
	Action    string            `json:"action,omitempty"`
	Method    string            `json:"method,omitempty"`
	Handler   string            `json:"handler,omitempty"`
	HandlerFn FormHandler       `json:"-"`
	// TODO: success/fail handler, how to push info back to front?
}

// Render ...
func (f *Form) Render(w io.Writer) error {
	if f.ID == "" {
		f.ID = nextFormID()
	}

	//
	if f.Method == "" {
		f.Method = "post"
	}

	//
	if f.Action == "" {
		f.Action = fmt.Sprintf("/submit/%s", f.Handler)
	}

	// Default styles
	if f.Styles == nil {
		f.Styles = map[string]string{}
	}
	f.Classes = append(f.Classes, "x-form")
	f.Styles["border"] = "1px solid red"
	return renderTemplate(w, "form", f)
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

func buildForm(i interface{}) *Form {
	ii := i.(map[string]interface{})

	p := &Form{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if docked, ok := ii["docked"]; ok {
		p.Docked = docked.(string)
	}

	if action, ok := ii["action"]; ok {
		p.Action = action.(string)
	}

	if method, ok := ii["method"]; ok {
		p.Method = method.(string)
	}

	if handler, ok := ii["handler"]; ok {
		p.Handler = handler.(string)
	}

	if c, ok := ii["classes"]; ok {
		jclass := c.([]interface{})
		classes := make([]string, len(jclass))
		for i, cl := range jclass {
			classes[i] = cl.(string)
		}
		p.Classes = classes
	}

	if s, ok := ii["styles"]; ok {
		jclass := s.(map[string]interface{})
		styles := map[string]string{}
		for i, cl := range jclass {
			styles[i] = cl.(string)
		}
		p.Styles = styles
	}

	items := Items{}
	if ii, ok := ii["items"]; ok {
		is := ii.([]interface{})
		for _, i := range is {
			item := addChild(i)
			items = append(items, item)
		}
	}

	p.Items = items

	return p
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
	XType   string            `json:"xtype"`
	ID      string            `json:"id,omitempty"`
	Docked  string            `json:"docked,omitempty"`
	Classes []string          `json:"classes,omitempty"`
	Styles  map[string]string `json:"styles,omitempty"`
	Legend  template.HTML     `json:"legend,omitempty"`
	Items   Items             `json:"items,omitempty"`
	Parent  Renderer          `json:"-"`
}

// Render ...
func (f *Fieldset) Render(w io.Writer) error {
	// if f.ID == "" {
	// 	f.ID = nextFormID()
	// }

	// if f.Styles == nil {
	// 	f.Styles = map[string]string{}
	// }
	// f.Classes = append(f.Classes, "x-form")
	// f.Styles["border"] = "1px solid lightgrey"

	// TODO: add legend to items
	items := Items{}
	if f.Legend != "" {
		// for now just use raw html
		items = append(items, &Element{
			Name:  "legend",
			Items: Items{&RawHTML{f.Legend}},
		})
	}
	// TODO: layout items
	for _, i := range f.Items {

		// append Label if any
		ii, ok := i.(*Input)
		if ok && ii.Label != "" {
			// input.ID might not be set yet
			if ii.ID == "" {
				ii.ID = nextInputID()
			}

			items = append(items, &Element{
				Name:  "label",
				Items: Items{&RawHTML{ii.Label}},
				Attributes: map[string]template.HTMLAttr{
					"for": template.HTMLAttr(ii.ID),
				},
			})
		}

		items = append(items, i)
	}

	div := &DivContainer{
		// ID:      p.ID,
		// Classes: npClasses,
		// Styles:  styles,
		Items: items,
	}

	return render(w, `<fieldset>
			{{range $item := $.Items}}
			{{Render $item}}
			{{end}}</fieldset>`, div)
}

// GetID ...
func (f *Fieldset) GetID() string {
	return f.ID
}

func buildFieldset(i interface{}) *Fieldset {
	ii := i.(map[string]interface{})

	p := &Fieldset{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if docked, ok := ii["docked"]; ok {
		p.Docked = docked.(string)
	}

	if c, ok := ii["classes"]; ok {
		jclass := c.([]interface{})
		classes := make([]string, len(jclass))
		for i, cl := range jclass {
			classes[i] = cl.(string)
		}
		p.Classes = classes
	}

	if s, ok := ii["styles"]; ok {
		jclass := s.(map[string]interface{})
		styles := map[string]string{}
		for i, cl := range jclass {
			styles[i] = cl.(string)
		}
		p.Styles = styles
	}

	items := []Renderer{}
	if ii, ok := ii["items"]; ok {
		is := ii.([]interface{})
		for _, i := range is {
			item := addChild(i)
			items = append(items, item)
		}
	}

	if legend, ok := ii["legend"]; ok {
		p.Legend = template.HTML(legend.(string))
	}

	p.Items = items
	return p
}

/*
 * Input
 */

// Input ...
type Input struct {
	XType        string                       `json:"xtype"`
	ID           string                       `json:"id,omitempty"`
	Classes      []string                     `json:"classes,omitempty"`
	Styles       map[string]string            `json:"styles,omitempty"`
	Type         string                       `json:"type,omitempty"`
	Name         string                       `json:"name,omitempty"`
	Value        string                       `json:"value,omitempty"`
	Attributes   map[string]template.HTMLAttr `json:"attributes,omitempty"`
	Form         string                       `json:"form,omitempty"`
	Disabled     bool                         `json:"disabled,omitempty"`
	Autofocus    bool                         `json:"autofocus,omitempty"`
	Autocomplete string                       `json:"autocomplete,omitempty"`
	Label        template.HTML                `json:"label,omitempty"`
	Parent       Renderer                     `json:"-"`
}

// Render ...
func (i *Input) Render(w io.Writer) error {
	if i.ID == "" {
		i.ID = nextInputID()
	}

	// if f.Styles == nil {
	// 	f.Styles = map[string]string{}
	// }
	// f.Classes = append(f.Classes, "x-form")
	// f.Styles["border"] = "1px solid lightgrey"

	// TODO: validate attributes based on type

	if len(i.Attributes) == 0 {
		i.Attributes = map[string]template.HTMLAttr{}
	}

	// ID
	i.Attributes["id"] = template.HTMLAttr(i.ID)

	// Type
	if i.Type == "" {
		// required, default?
		panic("wtf")
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

	e := &Element{
		Name:        "input",
		Attributes:  i.Attributes,
		SelfClosing: false,
	}
	return e.Render(w)
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

func buildInput(i interface{}) *Input {
	ii := i.(map[string]interface{})

	p := &Input{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if t, ok := ii["type"]; ok {
		p.Type = t.(string)
	}

	if name, ok := ii["name"]; ok {
		p.Name = name.(string)
	}

	if value, ok := ii["value"]; ok {
		p.Value = value.(string)
	}

	if form, ok := ii["form"]; ok {
		p.Form = form.(string)
	}

	if disabled, ok := ii["disabled"]; ok {
		p.Disabled = disabled.(bool)
	}

	if autofocus, ok := ii["autofocus"]; ok {
		p.Autofocus = autofocus.(bool)
	}

	if autocomplete, ok := ii["autocomplete"]; ok {
		p.Autocomplete = autocomplete.(string)
	}

	if label, ok := ii["label"]; ok {
		p.Label = template.HTML(label.(string))
	}

	if c, ok := ii["classes"]; ok {
		jclass := c.([]interface{})
		classes := make([]string, len(jclass))
		for i, cl := range jclass {
			classes[i] = cl.(string)
		}
		p.Classes = classes
	}

	if s, ok := ii["styles"]; ok {
		jclass := s.(map[string]interface{})
		styles := map[string]string{}
		for i, cl := range jclass {
			styles[i] = cl.(string)
		}
		p.Styles = styles
	}

	if s, ok := ii["attributes"]; ok {
		jclass := s.(map[string]interface{})
		attributes := map[string]template.HTMLAttr{}
		for i, cl := range jclass {
			attributes[i] = template.HTMLAttr(cl.(string))
		}
		p.Attributes = attributes
	}

	return p
}
