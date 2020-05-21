package ext

import (
	"encoding/json"
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

// MarshalJSON ...
func (t *Tree) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		XType      string            `json:"xtype"`
		ID         string            `json:"id,omitempty"`
		ShowRoot   bool              `json:"showRoot,omitempty"`
		Root       *TreeNode         `json:"root,omitempty"`
		BranchIcon string            `json:"branchIcon,omitempty"`
		LeafIcon   string            `json:"leafIcon,omitempty"`
		Docked     string            `json:"docked,omitempty"`
		Classes    []string          `json:"classes,omitempty"`
		Styles     map[string]string `json:"styles,omitempty"`
	}{
		XType:      "tree",
		ID:         t.ID,
		ShowRoot:   t.ShowRoot,
		Root:       t.Root,
		BranchIcon: t.BranchIcon,
		LeafIcon:   t.LeafIcon,
		Docked:     t.Docked,
		Classes:    t.Classes,
		Styles:     t.Styles,
	})
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

// MarshalJSON ...
func (tn *TreeNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		XType     string      `json:"xtype"`
		ID        string      `json:"id,omitempty"`
		Text      string      `json:"text,omitempty"`
		Collapsed bool        `json:"collapsed,omitempty"`
		Leaf      bool        `json:"leaf,omitempty"`
		IconClass string      `json:"iconClass,omitempty"`
		Children  []*TreeNode `json:"children,omitempty"`
	}{
		XType:     "treenode",
		ID:        tn.ID,
		Text:      tn.Text,
		Collapsed: tn.Collapsed,
		Leaf:      tn.Leaf,
		IconClass: tn.IconClass,
		Children:  tn.Children,
	})
}

func nextTreeID() string {
	id := fmt.Sprintf("tree-%d", treeID)
	treeID++
	return id
}
