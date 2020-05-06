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
		ID: nextContainerID(),
		// Width:  300,
		// Height: 200,
		// Border: template.CSS("1px solid lightgrey"),
	}
}

// Container ...
// TODO: I don't need all of this crap for container
type Container struct {
	ID string // how to auto generate
	// Title     string
	// IconClass string
	// Layout    string
	HTML template.HTML
	// Width     int // float?
	// Height    int // float?
	Items []Renderer
	// RenderTo  string // type???
	// Header    *Header
	// Border    template.CSS
	Docked string // top, bottom, left, right, ''
	// Flex      int
	// Style     string
	// Shadow    bool
	Classes []string
	Styles  map[string]string
	Body    *Body
}

// Render ...
func (c *Container) Render() template.HTML {
	fmt.Println("  render container")

	if c.ID == "" {
		c.ID = nextContainerID()
	}

	// default classes
	c.Classes = []string{
		"x-container",
		"x-container",
		"x-component",
		"x-noborder-trbl",
		"x-header-position-top",
		"x-root",
	}

	// HTML
	if c.HTML != "" {
		c.Items = append(c.Items, &Container{
			HTML: c.HTML,
		})
	}

	// ITEMS
	if len(c.Items) == 0 {
		// debug add dummy html
		html := c.HTML
		if html == "" {
			html = ":("
		}
		c.Items = []Renderer{
			&Container{
				HTML: html,
			},
		}
	}

	// BODY
	if c.Body == nil {
		c.Body = NewBody(c.Items)
	}

	styles := map[string]string{}

	// TODO: might have to check all items, if docked?

	// TODO: parse c.Style and add to styles
	c.Styles = styles

	// body
	// x-container-body x-body-wrap-el x-container-body-wrap-el x-container-body-wrap-el x-component-body-wrap-el
	// x-body-wrap-el x-container-body-wrap-el x-container-body-wrap-el x-component-body-wrap-el

	return render("container", c)
}

// Debug ...
func (c *Container) Debug() {}

func nextContainerID() string {
	id := fmt.Sprintf("%d", containerID)
	containerID++
	return id
}

// GetDocked ...
func (c *Container) GetDocked() string {
	return c.Docked
}
