package gbt

type Image struct {
	Src    string
	Alt    string
	Width  string
	Height string
	Classes
	Attributes
}

func NewImage(src string, opts ...Option[any]) *Image {
	if src == "" {
		return nil
	}
	fb := &Image{
		Src: src,
	}
	for _, op := range opts {
		op(any(fb))
	}
	return fb
}

func (i *Image) SetSize(w, h string) {
	if w != "" {
		i.Width = w
	}
	if h != "" {
		i.Height = h
	}
}

func (i *Image) SetAttributes(a Attributes) {
	i.Attributes = a
}

func (i *Image) SetClasses(c Classes) {
	i.Classes = c
}

func (i *Image) Render() Stringer {
	if i.Attributes == nil {
		i.Attributes = Attributes{}
	}
	attr := i.Attributes
	attr["src"] = i.Src
	if i.Width != "" {
		attr["width"] = i.Width
	}
	if i.Height != "" {
		attr["height"] = i.Height
	}
	if i.Alt != "" {
		attr["alt"] = i.Alt
	}
	el := NewElement(
		"img",
		WithAttributes(attr),
		WithClasses(i.Classes),
	)
	return el.Render()
}
