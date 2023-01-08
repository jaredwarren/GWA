package gbt

type InputType string

var (
	InputTypeButton        InputType = "button"
	InputTypeCheckbox      InputType = "checkbox"
	InputTypeColor         InputType = "color"
	InputTypeDate          InputType = "date"
	InputTypeDatetimeLocal InputType = "datetime-local"
	InputTypeEmail         InputType = "email"
	InputTypeFile          InputType = "file"
	InputTypeImage         InputType = "image"
	InputTypeMonth         InputType = "month"
	InputTypeNumber        InputType = "number"
	InputTypePassword      InputType = "password"
	InputTypeRadio         InputType = "radio"
	InputTypeRange         InputType = "range"
	InputTypeReset         InputType = "reset"
	InputTypeSearch        InputType = "search"
	InputTypeSubmit        InputType = "submit"
	InputTypeTel           InputType = "tel"
	InputTypeText          InputType = "text"
	InputTypeTime          InputType = "time"
	InputTypeURL           InputType = "url"
	InputTypeWeek          InputType = "week"
)

// Input ...
type Input struct {
	ID         string
	Type       InputType
	Classes    Classes
	Attributes Attributes
}

// Render ...
func (i *Input) Render() Stringer {
	// maybe need to override name for textarea
	e := &Element{
		Name:       "input",
		Attributes: i.Attributes,
		Classes:    i.Classes,
	}
	return e.Render()
}
