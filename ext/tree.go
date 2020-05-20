package ext

import (
	"fmt"
	"io"
)

var (
	treeID = 0
)

// Tree ...
type Tree struct {
	ID         string
	ShowRoot   bool
	Parent     Renderer
	Root       *TreeNode
	BranchIcon string
	LeafIcon   string
	Docked     string
	Classes    []string
	Styles     map[string]string
}

// Render ...
func (t *Tree) Render(w io.Writer) error {
	if t.ID == "" {
		t.ID = nextTreeID()
	}
	if t.Styles == nil {
		t.Styles = map[string]string{}
	}
	t.Classes = append(t.Classes, "x-tree")
	t.Styles["border"] = "1px solid lightgrey"
	return renderTemplate(w, "tree", t)
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

// TreeNode ...
type TreeNode struct {
	ID   string
	Text string
	// Expanded  bool
	Collapsed bool
	Leaf      bool
	IconClass string
	Children  []*TreeNode
}

// Render ...
func (tn *TreeNode) Render(w io.Writer) error {
	// if tn.ID == "" {
	// 	tn.ID = nextTreeID()
	// }
	// copy styles
	// styles := map[string]string{}
	// if len(t.Styles) > 0 {
	// 	styles = t.Styles
	// }

	return renderTemplate(w, "treenode", tn)
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
