package gbt

import (
	"fmt"
	"html/template"
	"math/rand"
)

// temp to test new interface
type INavItem interface {
	Render() Stringer
}

// NavItem belongs to Nav
type NavItem struct {
	Image    *Image
	Title    string
	Href     string
	Disabled bool
	Active   bool
}

func (n *NavItem) Render() Stringer {
	return renderToHTML(`<li class="nav-item"><a class="nav-link {{if .Active}}active{{end}} {{if .Disabled}}disabled{{end}}" aria-current="page" href="{{.Href}}">{{.Title}}</a></li>`, n)
}

// NavDropDown Navitem, but with dropdown items
type NavDropDown struct {
	Title string
	Items []INavItem
}

func (n *NavDropDown) Render() Stringer {
	return renderToHTML(`<li class="nav-item dropdown">
          <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
            {{.Title}}
          </a>
          <ul class="dropdown-menu">
		  	{{range $index, $item := .Items}}
				{{$item.Render}}
			{{end}}
          </ul>
        </li>`, n)
}

// DropDownItem belongs to NavDropDown
type DropDownItem struct {
	Title    string
	Href     string
	Disabled bool
	Active   bool
}

func (n *DropDownItem) Render() Stringer {
	return renderToHTML(`<li><a class="dropdown-item {{if .Active}}active{{end}} {{if .Disabled}}disabled{{end}}" href="{{.Href}}">{{.Title}}</a></li>`, n)
}

// DropDowndivider belongs to NavDropDown
type DropDowndivider struct {
	Title string
	Href  string
}

func (n *DropDowndivider) Render() Stringer {
	return template.HTML(`<li><hr class="dropdown-divider"></li>`)
}

// NavBrand nav band
type NavBrand struct {
	Image *Image
	Title string
	Href  string
}

func (n *NavBrand) Render() Stringer {
	var tmp string
	if n.Href == "" {
		tmp = `<span class="navbar-brand mb-0 h1">
			{{if .Image}}{{.Image.Render}}{{end}}
		{{.Title}}
		</span>`
	} else {
		tmp = `<a class="navbar-brand" href="{{.Href}}">
		{{if .Image}}{{.Image.Render}}{{end}}
		{{.Title}}
	</a>`
	}

	return renderToHTML(tmp, n)
}

// Nav ...
type Nav struct {
	ID string `json:"id,omitempty"` // how to auto generate
	// Title (optional) overwritten if Brand is set
	Title  template.HTML `json:"title,omitempty"`
	Brand  *NavBrand
	Items  []INavItem `json:"items,omitempty"`
	Theme  Theme
	Search bool
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
	XType  string        `json:"xtype"`
	Layout string        `json:"layout,omitempty"`
	HTML   template.HTML `json:"html,omitempty"`
}

// Render ...
func (n *Nav) Render() Stringer {
	if n.Theme == "" {
		n.Theme = ThemeLight
	}

	if n.ID == "" {
		n.ID = fmt.Sprintf("navbarSupportedContent%d", rand.Intn(100))
	}

	return renderToHTML(`
<nav class="navbar navbar-expand-lg bg-body-tertiary" data-bs-theme="{{.Theme}}">
  <div class="container-fluid">
	{{if .Brand}}
		{{.Brand.Render}}
	{{end}}
	{{if gt (len .Items) 0}}
		<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#{{.ID}}" aria-controls="{{.ID}}" aria-expanded="false" aria-label="Toggle navigation">
      	  <span class="navbar-toggler-icon"></span>
    	</button>
		<div class="collapse navbar-collapse" id="{{.ID}}">
		  <ul class="navbar-nav me-auto mb-2 mb-lg-0">
			{{range $index, $item := .Items}}
			  {{$item.Render}}
			{{end}}
		  </ul>
			{{if .Search}}
			  <form class="d-flex" role="search">
				<input class="form-control me-2" type="search" placeholder="Search" aria-label="Search">
				<button class="btn btn-outline-success" type="submit">Search</button>
			  </form>
			{{end}}
		</div>
	{{end}}
  </div>
</nav>`, n)
}
