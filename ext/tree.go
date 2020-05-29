package ext

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

var (
	treeID     = 0
	treeNodeID = 0
)

// Tree ...
type Tree struct {
	XType      string    `json:"xtype"`
	ID         string    `json:"id,omitempty"`
	ShowRoot   bool      `json:"showRoot,omitempty"`
	Root       *TreeNode `json:"root,omitempty"`
	Width      int       `json:"width,omitempty"`
	Height     int       `json:"height,omitempty"`
	BranchIcon string    `json:"branchIcon,omitempty"`
	LeafIcon   string    `json:"leafIcon,omitempty"`
	ParentIcon string    `json:"parentIcon,omitempty"`
	Docked     string    `json:"docked,omitempty"`
	Classes    Classes   `json:"classes,omitempty"`
	Styles     Styles    `json:"styles,omitempty"`
	Parent     Renderer  `json:"-"`
}

// Render ...
func (t *Tree) Render(w io.Writer) error {
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
	if t.ShowRoot {
		items = append(items, t.Root)
	} else {
		for _, i := range t.Root.Children {
			items = append(items, i)
		}
	}

	// Attributes
	attrs := Attributes{
		"id":    template.HTMLAttr(t.ID),
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
	return navEl.Render(w)
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
	Search    bool        `json:"search,omitempty"`
	IconClass string      `json:"iconClass,omitempty"`
	Children  []*TreeNode `json:"children,omitempty"`
}

// Render ...
func (tn *TreeNode) Render(w io.Writer) error {
	if tn.ID == "" {
		tn.ID = nextTreeNodeID()
	}
	items := Items{}

	// Attributes
	attrs := Attributes{
		"id": template.HTMLAttr(tn.ID),
	}

	classes := []string{}

	if len(tn.Children) > 0 {
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
				&RawHTML{template.HTML(tn.Text)},
			},
		}
		items = append(items, leaf)

		// Search
		if tn.Search {
			search := &Element{
				Name: "div",
				Attributes: Attributes{
					"style": "display: flex; flex-direction: row; align-items: center;",
				},
				Items: Items{&Element{
					Name: "i",
					Attributes: Attributes{
						"class": "fas fa-folder-open",
					},
				},
					&RawHTML{template.HTML(tn.Text)},
				},
			}
			items = append(items, search)
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
	} else {
		leaf := &Element{
			Name:       "span",
			Attributes: Attributes{},
			Items:      Items{},
		}
		if tn.Handler != "" {
			leaf.Attributes["onclick"] = template.HTMLAttr(fmt.Sprintf("%s('%s')", tn.Handler, tn.ID))
		}

		if tn.IconClass != "" {
			leaf.Items = append(leaf.Items, &Element{
				Name: "i",
				Attributes: Attributes{
					"class": template.HTMLAttr(tn.IconClass),
				},
			})
		}

		// add leaf text
		if tn.Text != "" {
			leaf.Items = append(leaf.Items, &RawHTML{template.HTML(tn.Text)})
		}

		items = append(items, leaf)

		// test other stuff in leav
		items = append(items, &Element{
			Name: "span",
			Items: Items{
				&Element{
					Name: "a",
					Attributes: Attributes{
						"href": template.HTMLAttr("#"),
					},
					Innerhtml: template.HTML(`<i class="fas fa-key"></i>`),
				},
				&Element{
					Name: "a",
					Attributes: Attributes{
						"href": template.HTMLAttr("#"),
					},
					Innerhtml: template.HTML(`<i class="fas fa-info-circle"></i>`),
				},
			},
		})

		classes = append(classes, "leaf")
	}

	attrs["class"] = template.HTMLAttr(strings.Join(classes, " "))

	tnEl := &Element{
		Name:       "li",
		Attributes: attrs,
		Items:      items,
	}
	return tnEl.Render(w)
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
