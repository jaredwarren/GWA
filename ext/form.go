package ext

import (
	"encoding/json"
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
	ID        string
	Parent    Renderer
	Docked    string
	Classes   []string
	Styles    map[string]string
	Items     Items
	Action    string
	Method    string
	Handler   string
	HandlerFn FormHandler
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

// MarshalJSON ...
func (f *Form) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		XType   string            `json:"xtype"`
		ID      string            `json:"id,omitempty"`
		Docked  string            `json:"docked,omitempty"`
		Classes []string          `json:"classes,omitempty"`
		Styles  map[string]string `json:"styles,omitempty"`
		Items   Items             `json:"items,omitempty"`
		Action  string            `json:"action,omitempty"`
		Method  string            `json:"method,omitempty"`
		Handler string            `json:"handler,omitempty"`
	}{
		XType:   "form",
		ID:      f.ID,
		Docked:  f.Docked,
		Classes: f.Classes,
		Styles:  f.Styles,
		Items:   f.Items,
		Action:  f.Action,
		Method:  f.Method,
		Handler: f.Handler,
	})
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
	ID      string
	Parent  Renderer
	Docked  string
	Classes []string
	Styles  map[string]string
	Legend  template.HTML
	Items   Items
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
			Name:      "legend",
			InnerHTML: f.Legend,
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
				Name:      "label",
				InnerHTML: ii.Label,
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

/*
 * Input
 */

// Input ...
type Input struct {
	ID           string
	Parent       Renderer
	Classes      []string
	Styles       map[string]string
	Type         string
	Name         string
	Value        string
	Attributes   map[string]template.HTMLAttr
	Form         string
	Disabled     bool
	Autofocus    bool
	Autocomplete string
	Label        template.HTML
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
