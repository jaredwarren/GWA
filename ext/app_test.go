package ext

import (
	"encoding/json"
	"testing"
)

func TestJSON(t *testing.T) {
	op := &Application{
		XType:       "app",
		Name:        "my app",
		Controllers: []*Controller{},
		MainView: &Panel{
			XType:     "panel",
			Title:     "Panel Title!",
			Shadow:    true,
			Layout:    "hbox",
			HTML:      "test",
			IconClass: ".IconClass.",
			Classes:   []string{"a", "b"},
			Styles:    map[string]string{"style-a": "red"},
			Items: []Renderer{
				&Panel{
					XType:  "panel",
					HTML:   "My panel t..1",
					Docked: "top",
				},
				&Panel{
					XType:  "panel",
					HTML:   "My panel t..3",
					Docked: "left",
				},
				&Panel{
					XType:  "panel",
					HTML:   "My panel t..4",
					Docked: "bottom",
				},
				&Button{
					XType:   "button",
					Text:    "Click Here",
					Handler: "btnClick",
				},
				&Button{
					XType: "button",
					Text:  "2 Here",
					HandlerFn: func(id string) {

					},
				},
				&Form{
					XType:   "form",
					Handler: "formSubmit",
					Items: []Renderer{
						&Fieldset{
							XType:  "fieldset",
							Legend: "Form Legend",
							Items: []Renderer{
								&Input{
									XType: "input",
									Label: "User Name:",
									Name:  "username",
									Type:  "text",
								},
								&Input{
									XType: "input",
									Label: "Send:",
									Name:  "submit",
									Type:  "submit",
								},
							},
						},
					},
				},
				&Tree{
					XType:      "tree",
					Docked:     "right",
					ShowRoot:   true,
					BranchIcon: "",
					LeafIcon:   "",
					Root: &TreeNode{
						Text: "root",
						Children: []*TreeNode{
							&TreeNode{
								Text:      "c1",
								IconClass: "fas fa-fighter-jet",
							},
							&TreeNode{
								Text:      "c2",
								IconClass: "fad fa-acorn",
							},
							&TreeNode{
								Text:      "c3",
								IconClass: "fad fa-arrow-alt-from-right",
							},
							&TreeNode{
								Text:      "c4",
								IconClass: "fad fa-tree-palm",
							},
							&TreeNode{
								Text: "c2",
								Children: []*TreeNode{
									&TreeNode{
										Text:     "c2c1",
										Children: []*TreeNode{},
									},
								},
							},
							&TreeNode{
								Text: "c3",
								Children: []*TreeNode{
									&TreeNode{
										Text:     "c3c1",
										Children: []*TreeNode{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	b, err := json.Marshal(op)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	if len(b) == 0 {
		t.Errorf("no data\n")
	}

	// test backwards

	newApp := &Application{}
	err = json.Unmarshal(b, newApp)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	if newApp.MainView.(*Panel).Title != op.MainView.(*Panel).Title {
		t.Errorf("main view panel\n")
	}

	if newApp.MainView.(*Panel).Items[0].(*Panel).HTML != op.MainView.(*Panel).Items[0].(*Panel).HTML {
		t.Errorf("items off\n")
	}

	if newApp.MainView.(*Panel).Styles["style-a"] != op.MainView.(*Panel).Styles["style-a"] {
		t.Errorf("mainpanel style wrong\n")
	}

}
