package ext

import (
	"encoding/json"
	"testing"
)

func TestPanelJSON(t *testing.T) {
	// test from panel to json
	op := &Panel{
		HTML:        "My panel t..1",
		Docked:      "top",
		Collapsable: true,
		Classes:     []string{"class-a", "class-b"},
		Controller:  &Controller{},
	}
	b, err := json.Marshal(op)
	if err != nil {
		t.Errorf("%+v\n", err)
	}
	if len(b) == 0 {
		t.Errorf("no data\n")
	}

	// test from json to panel
	panel := &Panel{}
	err = json.Unmarshal(b, panel)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	if op.HTML != panel.HTML {
		t.Errorf("HTML wrong\n")
	}

	if op.Docked != panel.Docked {
		t.Errorf("Docked wrong\n")
	}

	if op.Collapsable != panel.Collapsable {
		t.Errorf("Collapsable wrong\n")
	}

	if op.Classes[0] != panel.Classes[0] {
		t.Errorf("Class 0 wrong\n")
	}
}

func TestNoDocked(t *testing.T) {
	// i := []Renderer{
	// 	&Panel{
	// 		HTML: "test",
	// 	},
	// }
	// outItems := layoutItems(i)
	// if len(outItems) != len(i) {
	// 	t.Errorf("%+v\n", outItems)
	// }
}

func TestOneDocked(t *testing.T) {
	// i := []Renderer{
	// 	&Panel{
	// 		HTML: "Test panel 0",
	// 	},
	// 	&Panel{
	// 		HTML:   "Docked panel 1",
	// 		Docked: "left",
	// 	},
	// 	&Panel{
	// 		HTML: "test panel 2",
	// 	},
	// }
	// outItems := layoutItems(i)
	// // check that one layout item was returned
	// if len(outItems) != 1 {
	// 	t.Errorf("%+v\n", outItems)
	// }
	// layout, ok := outItems[0].(*Layout)
	// if !ok {
	// 	t.Errorf("Not layout:%+v", layout)
	// }

	// if len(layout.Items) != 2 {
	// 	t.Errorf("layout items length wrong:%+v", layout)
	// }

	// dockedPanel := layout.Items[0].(*Panel)
	// if dockedPanel.HTML != "Docked panel 1" {
	// 	t.Errorf("wrong panel docked:%+v", dockedPanel)
	// }

	// body := layout.Items[1].(*Body)
	// if len(body.Items) != 2 {
	// 	t.Errorf("wrong body item:%+v", dockedPanel)
	// }
}

func TestTowDocked(t *testing.T) {
	// i := []Renderer{
	// 	&Panel{
	// 		HTML:   "Top Docked panel 0",
	// 		Docked: "top",
	// 	},
	// 	&Panel{
	// 		HTML:   "Left Docked panel 1",
	// 		Docked: "left",
	// 	},
	// 	&Panel{
	// 		HTML: "test panel 2",
	// 	},
	// }
	// outItems := layoutItems(i)
	// // check that one layout item was returned
	// if len(outItems) != 1 {
	// 	t.Errorf("%+v\n", outItems)
	// }
	// layout, ok := outItems[0].(*Layout)
	// if !ok {
	// 	t.Errorf("Not layout:%+v", layout)
	// }

	// // make sure layout has 2 items (layout, body)
	// if len(layout.Items) != 2 {
	// 	t.Errorf("layout items length wrong:%+v", layout)
	// }

	// // make sure first item is top docked panel
	// dockedPanel := layout.Items[0].(*Panel)
	// if dockedPanel.HTML != "Top Docked panel 0" {
	// 	t.Errorf("wrong panel docked:%+v", dockedPanel)
	// }

	// // Make sure body only has 1 item (layout)
	// body := layout.Items[1].(*Body)
	// if len(body.Items) != 1 {
	// 	t.Errorf("wrong body item:%+v", dockedPanel)
	// }

	// // Check docked 2
	// layout2 := body.Items[0].(*Layout)
	// if len(layout2.Items) != 2 {
	// 	t.Errorf("wrong number of docked2 items:%+v", dockedPanel)
	// }

	// // make sure first item is top docked panel
	// dockedPanel2 := layout2.Items[0].(*Panel)
	// if dockedPanel2.HTML != "Left Docked panel 1" {
	// 	t.Errorf("wrong panel docked:%+v", dockedPanel2)
	// }

	// // Make sure body2 only has 1 item (non docked panel)
	// body2 := layout2.Items[1].(*Body)
	// p2 := body2.Items[0].(*Panel)
	// if p2.HTML != "test panel 2" {
	// 	t.Errorf("wrong body2 item:%+v", dockedPanel)
	// }
}

func TestRender(t *testing.T) {
	// p := &Panel{
	// 	Title: "main",
	// 	Items: []Renderer{
	// 		&Panel{
	// 			HTML:   "Test panel 0",
	// 			Width:  200,
	// 			Height: 200,
	// 			Shadow: true,
	// 			Border: "1px solid red",
	// 		},
	// 		&Panel{
	// 			HTML:   "Docked panel 1",
	// 			Docked: "left",
	// 		},
	// 		&Panel{
	// 			HTML: "test panel 2",
	// 		},
	// 	},
	// }
	// np, err := p.Build()
	// if err != nil {
	// 	t.Error("[E]", err)
	// }
	// if np == nil {
	// 	t.Errorf("no np")
	// }

	// Debug(np)
	// t.Errorf("no np")
	// return

	// for j, i := range np.(*Panel).Items {
	// 	fmt.Printf(" (%d)---> %+v\n", j, i)
	// }

	// layout := np.(*Panel).Items[0].(*Layout)
	// if len(layout.Items) != 2 {
	// 	t.Errorf("wrong number of items:%+v", layout)
	// }

	// header := layout.Items[0].(*Header)
	// if header.Title != "main" {
	// 	t.Errorf("wrong header:%+v", header)
	// }

	// dockedPanel1 := layout.Items[1].(*Body).Items[0].(*Layout).Items[0].(*Panel).HTML
	// if dockedPanel1 != "Docked panel 1" {
	// 	t.Errorf("wrong docked item:%+v", header)
	// }

	// bodyPanel0 := layout.Items[1].(*Body).Items[0].(*Layout).Items[1].(*Body).Items[0].(*Panel)

	// if bodyPanel0.Width != 200 {
	// 	t.Errorf("wrong width:%+v", layout.Items[1].(*Body).Items[0].(*Layout).Items[1].(*Body).Items[0].(*Panel))
	// }
	// if layout.Items[1].(*Body).Items[0].(*Layout).Items[1].(*Body).Items[0].(*Panel).Height != 200 {
	// 	t.Errorf("wrong height:%+v", layout.Items[1].(*Body).Items[0].(*Layout).Items[1].(*Body).Items[0].(*Panel))
	// }
	// if layout.Items[1].(*Body).Items[0].(*Layout).Items[1].(*Body).Items[0].(*Panel).Border != "1px solid red" {
	// 	t.Errorf("wrong border:%+v", layout.Items[1].(*Body).Items[0].(*Layout).Items[1].(*Body).Items[0].(*Panel))
	// }
	// if layout.Items[1].(*Body).Items[0].(*Layout).Items[1].(*Body).Items[0].(*Panel).Shadow != true {
	// 	t.Errorf("wrong shadow:%+v", layout.Items[1].(*Body).Items[0].(*Layout).Items[1].(*Body).Items[0].(*Panel))
	// }
	// fmt.Printf("%+v\n", layout.Items[1].(*Body).Items[0].(*Layout).Items[1].(*Body).Items[0].(*Panel))
	// t.Error("")
}
