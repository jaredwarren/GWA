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
	XType      string            `json:"xtype"`
	ID         string            `json:"id,omitempty"`
	ShowRoot   bool              `json:"showRoot,omitempty"`
	Root       *TreeNode         `json:"root,omitempty"`
	BranchIcon string            `json:"branchIcon,omitempty"`
	LeafIcon   string            `json:"leafIcon,omitempty"`
	Docked     string            `json:"docked,omitempty"`
	Classes    []string          `json:"classes,omitempty"`
	Styles     map[string]string `json:"styles,omitempty"`
	Parent     Renderer          `json:"-y"`
}

// Render ...
func (t *Tree) Render(w io.Writer) error {
	if t.ID == "" {
		t.ID = nextTreeID()
	}
	if t.Styles == nil {
		t.Styles = map[string]string{}
	}

	// default classes
	t.Classes = append(t.Classes, "x-tree")
	t.Classes = getClasses(t.Classes)

	t.Styles["border"] = "1px solid lightgrey"

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
	attrs := map[string]template.HTMLAttr{
		"id":    template.HTMLAttr(t.ID),
		"class": template.HTMLAttr(strings.Join(t.Classes, " ")),
	}
	if len(t.Styles) > 0 {
		attrs["style"] = styleToAttr(t.Styles)
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
	attrs := map[string]template.HTMLAttr{
		"id": template.HTMLAttr(tn.ID),
	}

	if len(tn.Children) > 0 {
		// add parent label
		leaf := &Element{
			Name: "span",
			Attributes: map[string]template.HTMLAttr{
				"class":   "parent",
				"onclick": "toggleNode(this)",
			},
			Items: Items{&Element{
				Name: "i",
				Attributes: map[string]template.HTMLAttr{
					"class": "fas fa-folder-open",
				},
			},
				&RawHTML{template.HTML(tn.Text)},
			},
		}
		items = append(items, leaf)

		// add sub-list
		list := &Element{
			Name:       "ul",
			Attributes: map[string]template.HTMLAttr{},
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
			Attributes: map[string]template.HTMLAttr{},
			Items:      Items{},
		}
		if tn.Handler != "" {
			leaf.Attributes["onclick"] = template.HTMLAttr(fmt.Sprintf("%s('%s')", tn.Handler, tn.ID))
		}

		if tn.IconClass != "" {
			leaf.Items = append(leaf.Items, &Element{
				Name: "i",
				Attributes: map[string]template.HTMLAttr{
					"class": template.HTMLAttr(tn.IconClass),
				},
			})
		}
		if tn.Text != "" {
			leaf.Items = append(leaf.Items, &RawHTML{template.HTML(tn.Text)})
		}

		items = append(items, leaf)
	}

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
