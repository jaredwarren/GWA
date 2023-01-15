package gbt

func NewNavItem(opts ...Option[any]) *NavItem {
	fb := &NavItem{
		Active:   true,
		Disabled: false,
	}
	for _, op := range opts {
		op(any(fb))
	}
	return fb
}

// NavItem belongs to Nav
type NavItem struct {
	Image    *Image
	Title    string
	Href     string
	Icon     *Icon
	Disabled bool
	Active   bool
}

func (n *NavItem) SetTitle(t string) {
	n.Title = t
}

func (n *NavItem) SetHref(href string) {
	n.Href = href
}

func (n *NavItem) SetActive(active bool) {
	n.Active = active
}

func (n *NavItem) SetDisabled(disabled bool) {
	n.Disabled = disabled
}

func (n *NavItem) SetImage(i *Image) {
	n.Image = i
}

func (e *NavItem) SetIcon(i *Icon) {
	e.Icon = i
}

func (n *NavItem) Render() Stringer {
	attributes := Attributes{"href": n.Href}
	classes := Classes{"nav-link"}
	if n.Active {
		classes = append(classes, "active")
		attributes["aria-current"] = "page"
	}
	if n.Disabled {
		classes = append(classes, "disabled")
	}

	items := Items{}
	if n.Icon != nil {
		items = append(items, n.Icon)
	}
	items = append(items, RawHTML(n.Title))

	return NewElement(
		"li",
		WithClasses(Classes{"nav-item"}),
		WithItems(Items{
			NewElement(
				"a",
				WithAttributes(attributes),
				WithClasses(classes),
				WithItems(items),
			),
		}),
	).Render()
}
