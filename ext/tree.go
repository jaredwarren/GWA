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
	t.Classes = t.GetClasses()

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

// GetClasses ...
func (t *Tree) GetClasses() []string {
	// default classes
	classess := map[string]bool{}
	// copy classes
	for _, c := range t.Classes {
		if _, ok := classess[c]; !ok {
			classess[c] = true
		}
	}
	// convert class back to array
	npClasses := []string{}
	for k := range classess {
		npClasses = append(npClasses, k)
	}
	return npClasses
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

func buildTree(i interface{}) *Tree {
	ii := i.(map[string]interface{})

	p := &Tree{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if showRoot, ok := ii["showRoot"]; ok {
		p.ShowRoot = showRoot.(bool)
	}

	if branchIcon, ok := ii["branchIcon"]; ok {
		p.BranchIcon = branchIcon.(string)
	}

	if leafIcon, ok := ii["leafIcon"]; ok {
		p.LeafIcon = leafIcon.(string)
	}

	if docked, ok := ii["docked"]; ok {
		p.Docked = docked.(string)
	}

	if c, ok := ii["classes"]; ok {
		jclass := c.([]interface{})
		classes := make([]string, len(jclass))
		for i, cl := range jclass {
			classes[i] = cl.(string)
		}
		p.Classes = classes
	}

	if s, ok := ii["styles"]; ok {
		jclass := s.(map[string]interface{})
		styles := map[string]string{}
		for i, cl := range jclass {
			styles[i] = cl.(string)
		}
		p.Styles = styles
	}

	if root, ok := ii["root"]; ok {
		p.Root = buildTreeNode(root)
	}

	return p
}

// TreeNode ...
type TreeNode struct {
	XType     string      `json:"xtype"`
	ID        string      `json:"id,omitempty"`
	Text      string      `json:"text,omitempty"`
	Handler   template.JS `json:"handler,omitempty"`
	Collapsed bool        `json:"collapsed,omitempty"`
	Leaf      bool        `json:"leaf,omitempty"`
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
		// "class": template.HTMLAttr(strings.Join(t.Classes, " ")),
	}
	// if len(t.Styles) > 0 {
	// 	attrs["style"] = styleToAttr(t.Styles)
	// }

	if len(tn.Children) > 0 {

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

	return renderTemplate(w, "treenode", tn)
}

// GetID ...
func (tn *TreeNode) GetID() string {
	return tn.ID
}

func buildTreeNode(i interface{}) *TreeNode {
	ii := i.(map[string]interface{})

	p := &TreeNode{}
	if ID, ok := ii["id"]; ok {
		p.ID = ID.(string)
	}

	if collapsed, ok := ii["collapsed"]; ok {
		p.Collapsed = collapsed.(bool)
	}

	if leaf, ok := ii["leaf"]; ok {
		p.Leaf = leaf.(bool)
	}

	if text, ok := ii["text"]; ok {
		p.Text = text.(string)
	}

	if iconClass, ok := ii["iconClass"]; ok {
		p.IconClass = iconClass.(string)
	}

	items := []*TreeNode{}
	if ii, ok := ii["children"]; ok {
		is := ii.([]interface{})
		for _, i := range is {
			item := buildTreeNode(i)
			items = append(items, item)
		}
	}
	p.Children = items

	return p
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
