package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/jaredwarren/app"
	"github.com/jaredwarren/goext/ext"
	"github.com/jaredwarren/goext/service"
)

var (
	// TODO ...
	TODO = []string{
		"\n☐ json to panel (use xtype)",
		"\n☐ make header 'docked'", // I think this is done, but header template needs cleaned up
		"\n☐ fix things so they work in test\n",
		"\n☐ make 'app' class that's full-screen, merge with native app\n",
	}
)

func main() {
	fmt.Println("TODO:", TODO)

	// testJSON()
	serve()
	// db()
}

// TODO:

func serve() {
	// return
	conf := &app.WebConfig{
		Host: "127.0.0.1",
		Port: 8083,
	}
	a := app.NewWeb(conf)

	service.Register(a)

	d := <-a.Exit
	fmt.Printf("Done:%+v\n", d)
}

func testJSON() {
	data := []byte(`{
		"title":  "Panel",
		"html": "<div>asdf</div",
		"asdf": "asdf",
		"items": [
			{
				"title":  "Panel",
				"html": "<div>asdf</div"
			}
		]
	}`)

	var np interface{}

	err := json.Unmarshal(data, &np)
	if err != nil {
		fmt.Println("[E]", err)
	}

	// assume root is panel
	root := &ext.Panel{}
	e := reflect.ValueOf(root).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		// varValue := e.Field(i).Interface()
		// fmt.Printf("%v %v %v\n", varName, varType, varValue)
		for k, v := range np.(map[string]interface{}) {
			if strings.ToLower(varName) == strings.ToLower(k) {
				fmt.Printf("%+v -> %+v.(%s)\n", k, v, varType)
			}

		}

	}
}

func db() {
	// p := &ext.Panel{
	// 	Title: "main",
	// 	Items: []ext.Renderer{
	// 		&ext.Panel{
	// 			HTML:   "Test panel 0",
	// 			Width:  200,
	// 			Height: 200,
	// 			Shadow: true,
	// 			Border: "1px solid red",
	// 		},
	// 		&ext.Panel{
	// 			HTML:   "Docked panel 1",
	// 			Docked: "left",
	// 		},
	// 		&ext.Panel{
	// 			HTML: "test panel 2",
	// 		},
	// 	},
	// }
	p := &ext.Panel{
		Title:  "Header Title",
		Width:  300,
		Height: 200,
		Shadow: true,
		// Scrollable: true,
		// Padding: 20,
		Layout: "hbox",
		// Defaults: []interface{}{
		// 	Flex: 1,
		// },
		HTML: "test",
		Items: ext.Items{
			&ext.Panel{
				HTML:   "My panel text...1",
				Docked: "top",
				// Style:  "",
				Flex: 1, // because default doesn't work
			},
			&ext.Panel{
				HTML:   "My panel text...2",
				Docked: "right",
				// Style:  "",
				Flex: 1, // because default doesn't work
			},
		},
	}
	fmt.Printf("%+v\n", p)
	// ext.Debug(p)

	// d, err := json.Marshal(np)
	// if err != nil {
	// 	fmt.Printf("%+v\n", err)
	// }
	// fmt.Printf("%+v\n", string(d))
}
