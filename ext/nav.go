package ext

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strings"
)

var (
	navID = 0
)

// NewNav ...
func NewNav(title template.HTML) *Nav {
	return &Nav{
		ID:     nextNavID(),
		Title:  title,
		Border: template.CSS("1px solid lightgrey"),
		Docked: "top",
	}
}

type Stringer interface{}

// temp to test new interface
type NavItem interface {
	Render() Stringer
	// RenderString() interface{}
}

type NavBrand struct {
	Image NavItem
	Title string
	Href  string
}

func (n *NavBrand) Render() Stringer {
	t := template.New("nav")
	if n.Href == "" {
		t, _ = t.Parse(`<span class="navbar-brand mb-0 h1">
			{{if .Image}}{{.Image.Render}}{{end}}
		{{.Title}}
		</span>`)
	} else {
		t, _ = t.Parse(`<a class="navbar-brand" href="{{.Href}}">
		{{if .Image}}{{.Image.Render}}{{end}}
		{{.Title}}
	</a>`)
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, n)
	if err != nil {
		fmt.Println("[E] render head:", err)
	}

	return template.HTML(buf.String())
}

type Theme template.HTMLAttr

var (
	ThemeDark  Theme = "dark"
	ThemeLight Theme = "light"
)

// Nav ...
// TODO: I don't need all of this crap for container
type Nav struct {
	XType     string        `json:"xtype"`
	ID        string        `json:"id,omitempty"` // how to auto generate
	Title     template.HTML `json:"title,omitempty"`
	TitleLink string
	IconClass string    `json:"iconClass,omitempty"`
	Items     []NavItem `json:"items,omitempty"`
	Theme     Theme
	//
	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`

	Border  template.CSS      `json:"border,omitempty"`
	Style   string            `json:"style,omitempty"`
	Docked  string            `json:"docked,omitempty"` // top, bottom, left, right
	Classes []string          `json:"classes,omitempty"`
	Styles  map[string]string `json:"styles,omitempty"`
	Shadow  bool              `json:"shadow,omitempty"`
	//
	Layout string        `json:"layout,omitempty"`
	HTML   template.HTML `json:"html,omitempty"`
}

// Render ...
func (n *Nav) Render(w io.Writer) error {
	if n.ID == "" {
		n.ID = nextNavID()
	}

	navClasses := []string{
		"navbar",
		"navbar-expand-lg",
		"bg-body-tertiary",
	}
	navClass := template.HTMLAttr(strings.Join(navClasses, " "))

	if n.Theme == "" {
		n.Theme = ThemeLight
	}

	data := map[string]interface{}{
		"NavClass": navClass,
		"Items":    n.Items,
		"Theme":    n.Theme,
	}

	t := template.New("nav")
	t, _ = t.Parse(`
<nav class="navbar navbar-expand-lg bg-body-tertiary" data-bs-theme="{{.Theme}}">
  <div class="container-fluid">
	{{range $index, $item := .Items}}
	  {{$item.Render}}
	{{end}}
    <a class="navbar-brand" href="#">Navbar</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav me-auto mb-2 mb-lg-0">
        <li class="nav-item">
          <a class="nav-link active" aria-current="page" href="#">Home</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Link</a>
        </li>
        <li class="nav-item dropdown">
          <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
            Dropdown
          </a>
          <ul class="dropdown-menu">
            <li><a class="dropdown-item" href="#">Action</a></li>
            <li><a class="dropdown-item" href="#">Another action</a></li>
            <li><hr class="dropdown-divider"></li>
            <li><a class="dropdown-item" href="#">Something else here</a></li>
          </ul>
        </li>
        <li class="nav-item">
          <a class="nav-link disabled">Disabled</a>
        </li>
      </ul>
      <form class="d-flex" role="search">
        <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search">
        <button class="btn btn-outline-success" type="submit">Search</button>
      </form>
    </div>
  </div>
</nav>`)

	return t.Execute(w, data)

	return nil

	// default classes
	classess := map[string]bool{
		"navbar":           true,
		"navbar-expand-lg": true,
		"navbar-dark":      true,
		"bg-dark":          true,
	}
	// navbar navbar-expand-lg bg-body-tertiary

	// copy classes
	for _, c := range n.Classes {
		classess[c] = true
	}

	if n.Shadow {
		classess["x-shadow"] = true
	}

	// copy styles
	styles := Styles{}
	if len(n.Styles) > 0 {
		styles = n.Styles
	}

	// append new styles based on p's properties
	if n.Width != 0 && n.Docked != "top" && n.Docked != "bottom" {
		styles["width"] = fmt.Sprintf("%dpx", n.Width)
		classess["x-widthed"] = true
	}
	if n.Height != 0 && n.Docked != "left" && n.Docked != "right" {
		styles["height"] = fmt.Sprintf("%dpx", n.Height)
		classess["x-heighted"] = true
	}
	if n.Border != "" {
		styles["border"] = string(n.Border)
		classess["x-managed-border"] = true
	}

	// convert class back to array
	npClasses := []string{}
	for k := range classess {
		npClasses = append(npClasses, k)
	}

	items := Items{}

	// Title
	if n.Title != "" {
		title := &Innerhtml{
			HTML:    n.Title,
			Classes: Classes{"title"},
		}
		items = append(items, title)
	}

	// append rest of items
	if len(n.Items) > 0 {
		// items = append(items, n.Items...)
	}

	// update parent
	for _, i := range items {
		c, ok := i.(Child)
		if ok {
			c.SetParent(n)
		}
	}

	// Attributes
	attrs := map[string]template.HTMLAttr{
		"id":    template.HTMLAttr(n.ID),
		"class": template.HTMLAttr(strings.Join(npClasses, " ")),
	}
	if len(styles) > 0 {
		attrs["style"] = styles.ToAttr()
	}

	navEl := &Element{
		Name:       "nav",
		Attributes: attrs,
		Items:      items,
	}
	return navEl.Render(w)
}

// GetID ...
func (n *Nav) GetID() string {
	return n.ID
}

// GetDocked ...
func (n *Nav) GetDocked() string {
	return n.Docked
}

// SetStyle ...
func (n *Nav) SetStyle(key, value string) {
	if n.Styles == nil {
		n.Styles = map[string]string{}
	}
	n.Styles[key] = value
}

func nextNavID() string {
	id := fmt.Sprintf("nav-%d", navID)
	navID++
	return id
}
