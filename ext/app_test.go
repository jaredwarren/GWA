package ext

import (
	"encoding/json"
	"testing"
)

func TestJSON(t *testing.T) {

	app := &Application{
		Name:        "my app",
		Controllers: []*Controller{},
		MainView: &Panel{
			Title:     "Panel Title!",
			Shadow:    true,
			Layout:    "hbox",
			HTML:      "test",
			IconClass: ".IconClass.",
			Classes:   []string{"a", "b"},
			Styles:    map[string]string{"style-a": "red"},
			Items: []Renderer{
				&Panel{
					HTML:   "My panel t..1",
					Docked: "top",
				},
				&Panel{
					HTML:   "My panel t..3",
					Docked: "left",
				},
				&Panel{
					HTML:   "My panel t..4",
					Docked: "bottom",
				},
				&Button{
					Text:    "Click Here",
					Handler: "btnClick",
				},
				&Button{
					Text: "2 Here",
					HandlerFn: func(id string) {

					},
				},
				&Form{
					Handler: "formSubmit",
					Items: []Renderer{
						&Fieldset{
							Legend: "Form Legend",
							Items: []Renderer{
								&Input{
									Label: "User Name:",
									Name:  "username",
									Type:  "text",
								},
								&Input{
									Label: "Send:",
									Name:  "submit",
									Type:  "submit",
								},
							},
						},
					},
				},
				&Tree{
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

	b, err := json.Marshal(app)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	if len(b) == 0 {
		t.Errorf("no data\n")
	}

	// test backwards

	app = &Application{}
	err = json.Unmarshal(b, app)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	if app.MainView.(*Panel).Items[0].(*Panel).HTML != "My panel t..1" {
		t.Errorf("items off\n")
	}

	if app.MainView.(*Panel).Styles["style-a"] != "red" {
		t.Errorf("mainpanel style wrong\n")
	}

}
