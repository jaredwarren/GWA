package gbt

import (
	"fmt"
	"html/template"
)

var (
	treeID     = 0
	treeNodeID = 0
)

// Tree ...
type Tree struct {
	XType      string        `json:"xtype"`
	ID         string        `json:"id,omitempty"`
	Title      template.HTML `json:"title,omitempty"`
	ShowRoot   bool          `json:"showRoot,omitempty"`
	Header     *Header       `json:"header,omitempty"`
	Root       *TreeNode     `json:"root,omitempty"`
	Width      int           `json:"width,omitempty"`
	Height     int           `json:"height,omitempty"`
	BranchIcon string        `json:"branchIcon,omitempty"`
	LeafIcon   string        `json:"leafIcon,omitempty"`
	ParentIcon string        `json:"parentIcon,omitempty"`
	Docked     string        `json:"docked,omitempty"`
	Classes    Classes       `json:"classes,omitempty"`
	Styles     Styles        `json:"styles,omitempty"`
	Parent     Renderer      `json:"-"`
}

// Render ...
func (t *Tree) Render() Stringer {
	if t.ID == "" {
		t.ID = nextTreeID()
	}

	if t.Styles == nil {
		t.Styles = Styles{}
	}
	t.Styles["border"] = "1px solid lightgrey"
	if t.Width != 0 && t.Docked != "top" && t.Docked != "bottom" {
		t.Styles["width"] = fmt.Sprintf("%dpx", t.Width)
	}
	// what if I want height to be 0px?
	if t.Height != 0 && t.Docked != "left" && t.Docked != "right" {
		t.Styles["height"] = fmt.Sprintf("%dpx", t.Height)
	}

	// default classes
	t.Classes = append(t.Classes, "x-tree")

	//
	items := Items{}

	// HEADER
	var header *Header
	if t.Title != "" {
		if t.Header == nil {
			// TODO: I don't like the "NewHeader" fn
			header = NewHeader(t.Title)
		} else if t.Header.Title == "" {
			header = t.Header
			header.Title = t.Title
			header.Docked = "top"
		} // else header is all set, ignore Title attribute
	} else if t.Header != nil {
		header = t.Header
	} // else assume no header

	// append header as docked item[0]
	if header != nil {
		if header.Docked == "" {
			header.Docked = "top"
		}
		items = append(items, header)
	}

	if t.ShowRoot {
		items = append(items, t.Root)
	} else {
		for _, i := range t.Root.Children {
			items = append(items, i)
		}
	}

	// Attributes
	attrs := Attributes{
		"id":    t.ID,
		"class": t.Classes.ToAttr(),
	}
	if len(t.Styles) > 0 {
		attrs["style"] = t.Styles.ToAttr()
	}

	navEl := &Element{
		Name:       "ul",
		Attributes: attrs,
		Items:      items,
	}
	return navEl.Render()
}

// GetID ...
func (t *Tree) GetID() string {
	return t.ID
}

// SetParent ...
func (t *Tree) SetParent(p Renderer) {
	t.Parent = p
}

// GetDocked ...
func (t *Tree) GetDocked() string {
	return t.Docked
}

// SetStyle ...
func (t *Tree) SetStyle(key, value string) {
	if t.Styles == nil {
		t.Styles = map[string]string{}
	}
	t.Styles[key] = value
}

// TreeNode ...
type TreeNode struct {
	XType     string      `json:"xtype"`
	ID        string      `json:"id,omitempty"`
	Text      string      `json:"text,omitempty"`
	Handler   template.JS `json:"handler,omitempty"`
	Collapsed bool        `json:"collapsed,omitempty"`
	Leaf      bool        `json:"leaf,omitempty"`
	Search    *Search     `json:"search,omitempty"`
	IconClass string      `json:"iconClass,omitempty"`
	Children  []*TreeNode `json:"children,omitempty"`
	Items     Items       `json:"items,omitempty"`
}

// Render ...
func (tn *TreeNode) Render() Stringer {
	if tn.ID == "" {
		tn.ID = nextTreeNodeID()
	}

	if len(tn.Children) > 0 {
		return tn.RenderParent()
	}

	return tn.RenderLeaf()
}

// RenderParent ...
func (tn *TreeNode) RenderParent() Stringer {
	items := Items{}

	// Attributes
	attrs := Attributes{
		"id": tn.ID,
	}

	// add parent label
	leaf := &Element{
		Name: "span",
		Attributes: Attributes{
			"class":   "parent",
			"onclick": "toggleNode(this)",
		},
		Items: Items{&Element{
			Name: "i",
			Attributes: Attributes{
				"class": "fas fa-folder-open",
			},
		},
			RawHTML(template.HTML(tn.Text)),
		},
	}
	items = append(items, leaf)

	// Search
	if tn.Search != nil {
		items = append(items, tn.Search)
	}

	// add sub-list
	list := &Element{
		Name:       "ul",
		Attributes: Attributes{},
		Items:      Items{},
	}
	if tn.Collapsed {
		list.Attributes["class"] = "collapsed"
	} else {
		list.Attributes["class"] = "expanded"
	}
	for _, i := range tn.Children {
		list.Items = append(list.Items, i)
	}
	items = append(items, list)

	tnEl := &Element{
		Name:       "li",
		Attributes: attrs,
		Items:      items,
	}
	return tnEl.Render()
}

// RenderLeaf ...
func (tn *TreeNode) RenderLeaf() Stringer {
	// Attributes
	attrs := Attributes{
		"id": tn.ID,
	}
	leaf := &Element{
		Name: "span",
		Attributes: Attributes{
			"class": "x-hbox",
		},
		Items: Items{},
	}
	if tn.Handler != "" {
		leaf.Attributes["onclick"] = fmt.Sprintf("%s('%s')", tn.Handler, tn.ID)
	}

	// if tn.IconClass != "" {
	leaf.Items = append(leaf.Items, &Element{
		Name: "i",
		Attributes: Attributes{
			"class": tn.IconClass,
		},
	})
	// }

	// add leaf text
	if tn.Text != "" {
		leaf.Items = append(leaf.Items, &Element{
			Name: "span",
			Attributes: Attributes{
				"style": "align-self:center", // TODO: fix this!!!
			},
			InnerHTML: template.HTML(tn.Text),
		})

	}

	// items := Items{leaf}
	// Add other items
	extra := &Element{
		Name: "span",
		Attributes: Attributes{
			"class": "x-hbox",
		},
		Items: Items{},
	}
	for _, i := range tn.Items {
		extra.Items = append(extra.Items, i)
	}

	classes := Classes{"leaf"}
	attrs["class"] = classes.ToAttr()

	tnEl := &Element{
		Name:       "li",
		Attributes: attrs,
		Items:      Items{leaf, extra},
	}
	return tnEl.Render()
}

// GetID ...
func (tn *TreeNode) GetID() string {
	return tn.ID
}

func nextTreeID() string {
	id := fmt.Sprintf("tree-%d", treeID)
	treeID++
	return id
}

func nextTreeNodeID() string {
	id := fmt.Sprintf("tree-node-%d", treeNodeID)
	treeNodeID++
	return id
}

/**
*
 */

// Search ...
type Search struct {
	XType   string      `json:"xtype"`
	ID      string      `json:"id,omitempty"`
	Text    string      `json:"text,omitempty"`
	Handler template.JS `json:"handler,omitempty"`
}

// Render ...
func (s *Search) Render() Stringer {
	// if tn.ID == "" {
	// 	tn.ID = nextTreeNodeID()
	// }

	search := &Element{
		Name: "div",
		Attributes: Attributes{
			"style": "display: flex; flex-direction: row; align-items: center;",
		},
		Items: Items{
			&Input{},
			&Button{
				Classes: Classes{"button-xsmall", "pure-button"},
				// Styles:    Styles{"visibility": "hidden"},
				Handler: "clearSearch",
				// IconClass: "fad fa-times-circle",
			},
		},
	}

	return search.Render()
}

// GetID ...
func (s *Search) GetID() string {
	return s.ID
}
