package ext

import (
	"testing"
)

func TestNoDocked(t *testing.T) {
	i := []Renderer{
		&Panel{
			HTML: "test",
		},
	}
	outItems := layoutItems(i)
	if len(outItems) != len(i) {
		t.Errorf("%+v\n", outItems)
	}
}

func TestOneDocked(t *testing.T) {
	i := []Renderer{
		&Panel{
			HTML: "Test panel 0",
		},
		&Panel{
			HTML:   "Docked panel 1",
			Docked: "left",
		},
		&Panel{
			HTML: "test panel 2",
		},
	}
	outItems := layoutItems(i)
	// check that one layout item was returned
	if len(outItems) != 1 {
		t.Errorf("%+v\n", outItems)
	}
	layout, ok := outItems[0].(*Layout)
	if !ok {
		t.Errorf("Not layout:%+v", layout)
	}

	if len(layout.Items) != 2 {
		t.Errorf("layout items length wrong:%+v", layout)
	}

	dockedPanel := layout.Items[0].(*Panel)
	if dockedPanel.HTML != "Docked panel 1" {
		t.Errorf("wrong panel docked:%+v", dockedPanel)
	}

	body := layout.Items[1].(*Body)
	if len(body.Items) != 2 {
		t.Errorf("wrong body item:%+v", dockedPanel)
	}

}

func TestTowDocked(t *testing.T) {
	i := []Renderer{
		&Panel{
			HTML:   "Top Docked panel 0",
			Docked: "top",
		},
		&Panel{
			HTML:   "Left Docked panel 1",
			Docked: "left",
		},
		&Panel{
			HTML: "test panel 2",
		},
	}
	outItems := layoutItems(i)
	// check that one layout item was returned
	if len(outItems) != 1 {
		t.Errorf("%+v\n", outItems)
	}
	layout, ok := outItems[0].(*Layout)
	if !ok {
		t.Errorf("Not layout:%+v", layout)
	}

	// make sure layout has 2 items (layout, body)
	if len(layout.Items) != 2 {
		t.Errorf("layout items length wrong:%+v", layout)
	}

	// make sure first item is top docked panel
	dockedPanel := layout.Items[0].(*Panel)
	if dockedPanel.HTML != "Top Docked panel 0" {
		t.Errorf("wrong panel docked:%+v", dockedPanel)
	}

	// Make sure body only has 1 item (layout)
	body := layout.Items[1].(*Body)
	if len(body.Items) != 1 {
		t.Errorf("wrong body item:%+v", dockedPanel)
	}

	// Check docked 2
	layout2 := body.Items[0].(*Layout)
	if len(layout2.Items) != 2 {
		t.Errorf("wrong number of docked2 items:%+v", dockedPanel)
	}

	// make sure first item is top docked panel
	dockedPanel2 := layout2.Items[0].(*Panel)
	if dockedPanel2.HTML != "Left Docked panel 1" {
		t.Errorf("wrong panel docked:%+v", dockedPanel2)
	}

	// Make sure body2 only has 1 item (non docked panel)
	body2 := layout2.Items[1].(*Body)
	p2 := body2.Items[0].(*Panel)
	if p2.HTML != "test panel 2" {
		t.Errorf("wrong body2 item:%+v", dockedPanel)
	}
}
