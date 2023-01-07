package gbt

import "html/template"

type Card struct {
	Header *CardHeader
	Footer *CardFooter
	Body   Items
	Classes
	Styles
}

func (c *Card) Render() Stringer {
	return renderToHTML(`<div class="card {{.Classes.Render}}">{{.Header.Render}}{{.Body.Render}}{{.Footer.Render}}</div>`, c)
}

type CardHeader struct {
	Body Items
	Classes
}

func (c *CardHeader) Render() Stringer {
	if len(c.Body) == 0 {
		return ""
	}
	e := &Element{
		Classes: Classes{"card-header"},
		Items:   c.Body,
	}
	return e.Render()
}

type CardFooter struct {
	Body Items
	Classes
}

func (c *CardFooter) Render() Stringer {
	if len(c.Body) == 0 {
		return ""
	}
	e := &Element{
		Classes: Classes{"card-footer"},
		Items:   c.Body,
	}
	return e.Render()
}

type CardBody struct {
	Title string
	Text  string
	Body  Items
	Classes
}

func (c *CardBody) Render() Stringer {
	i := Items{}
	if c.Title != "" {
		i = append(i, &CardTitle{
			Text: template.HTML(c.Title),
		})
	}
	if c.Text != "" {
		i = append(i, &CardText{
			Text: template.HTML(c.Text),
		})
	}
	c.Body = append(i, c.Body...)
	if len(c.Body) == 0 {
		return ""
	}
	e := &Element{
		Classes: Classes{"card-body"},
		Items:   c.Body,
	}
	return e.Render()
}

type CardTitle struct {
	Text template.HTML
	Classes
}

func (c *CardTitle) Render() Stringer {
	e := &Element{
		Classes:   Classes{"card-title"},
		InnerHTML: c.Text,
	}
	return e.Render()
}

type CardText struct {
	Text template.HTML
	Classes
}

func (c *CardText) Render() Stringer {
	e := &Element{
		Name:      "p",
		Classes:   Classes{"card-text"},
		InnerHTML: c.Text,
	}
	return e.Render()
}

type CardLink struct {
	Href string
	Text string
	Classes
}

func (c *CardLink) Render() Stringer {
	e := &Element{
		Name:       "a",
		Classes:    Classes{"card-link"},
		Attributes: Attributes{"href": c.Href},
		InnerHTML:  template.HTML(c.Text),
	}
	return e.Render()
}
