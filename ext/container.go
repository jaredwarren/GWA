package ext

import (
	"fmt"
	"html/template"
)

var (
	containerID = 0
)

// NewContainer ...
func NewContainer() *Container {
	return &Container{
		ID:     nextContainerID(),
		Width:  300,
		Height: 200,
		Border: template.CSS("1px solid lightgrey"),
	}
}

// Container ...
// TODO: I don't need all of this crap for container
type Container struct {
	ID        string // how to auto generate
	Title     string
	IconClass string
	Layout    string
	HTML      template.HTML
	Width     int // float?
	Height    int // float?
	Items     []Renderer
	RenderTo  string // type???
	Header    *Header
	Border    template.CSS
	Docked    string // top, bottom, left, right
	Flex      int
	Style     string
	Shadow    bool
	Classes   []string
	Styles    map[string]string
	Body      string
}

// Render ...
func (p *Container) Render() template.HTML {
	fmt.Println("  render container")
	if p.Header == nil {
		p.Header = &Header{}
		if p.Title != "" {
			p.Header.Title = p.Title
		}
	}

	if p.ID == "" {
		p.ID = nextContainerID()
	}

	// default classes
	p.Classes = []string{
		"x-container",
		"x-container",
		"x-component",
		"x-noborder-trbl",
		"x-header-position-top",
		"x-root",
	}

	if p.Shadow {
		p.Classes = append(p.Classes, "x-shadow")
	}

	styles := map[string]string{}
	if p.Width != 0 { // what if I want width to be 0px?
		styles["width"] = fmt.Sprintf("%dpx", p.Width)
		p.Classes = append(p.Classes, "x-widthed")
	}
	if p.Height != 0 { // what if I want height to be 0px?
		styles["height"] = fmt.Sprintf("%dpx", p.Height)
		p.Classes = append(p.Classes, "x-heighted")
	}
	if p.Border != "" {
		styles["border"] = string(p.Border)
		p.Classes = append(p.Classes, "x-managed-border")
	}

	// TODO: might have to check all items, if docked?

	// TODO: parse p.Style and add to styles
	p.Styles = styles

	// body
	// x-container-body x-body-wrap-el x-container-body-wrap-el x-container-body-wrap-el x-component-body-wrap-el
	// x-body-wrap-el x-container-body-wrap-el x-container-body-wrap-el x-component-body-wrap-el

	return render("container", p)
}

func nextContainerID() string {
	id := fmt.Sprintf("%d", containerID)
	containerID++
	return id
}
