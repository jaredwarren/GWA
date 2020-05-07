package main

import (
	"fmt"

	"github.com/jaredwarren/app"
	"github.com/jaredwarren/goext/service"
)

func main() {
	serve()
}

func serve() {

	// data := []byte(`{
	// 	"title":  "Panel",
	// 	"html": "<div>asdf</div",
	// 	"asdf": "asdf",
	// 	"items": [
	// 		{
	// 			"title":  "Panel",
	// 			"html": "<div>asdf</div"
	// 		}
	// 	]
	// }`)

	// var np interface{}

	// err := json.Unmarshal(data, &np)
	// if err != nil {
	// 	fmt.Println("[E]", err)
	// }

	// // assume root is panel
	// root := &ext.Panel{}
	// e := reflect.ValueOf(root).Elem()
	// for i := 0; i < e.NumField(); i++ {
	// 	varName := e.Type().Field(i).Name
	// 	varType := e.Type().Field(i).Type
	// 	// varValue := e.Field(i).Interface()
	// 	// fmt.Printf("%v %v %v\n", varName, varType, varValue)
	// 	for k, v := range np.(map[string]interface{}) {
	// 		if strings.ToLower(varName) == strings.ToLower(k) {
	// 			fmt.Printf("%+v -> %+v.(%s)\n", k, v, varType)
	// 		}

	// 	}

	// }
	// return

	//
	//
	//
	//

	// fmt.Printf("%+v\n", e)

	// pp := np.(*ext.Panel)
	// fmt.Printf("%+v\n", pp)

	// return

	// fmt.Printf("%+v\n", np)

	// TODO: make way to import json

	// panel := &ext.Panel{
	// 	Title:  "Panel",
	// 	Width:  300,
	// 	Height: 200,
	// 	Shadow: true,
	// 	// Scrollable: true,
	// 	// Padding: 20,
	// 	Layout: "hbox",
	// 	// Defaults: []interface{}{
	// 	// 	Flex: 1,
	// 	// },
	// 	HTML: "test",
	// 	Items: []ext.Renderer{
	// 		&ext.Panel{
	// 			HTML:   "My panel text...1",
	// 			Docked: "left",
	// 			Style:  "",
	// 			Flex:   1, // because default doesn't work
	// 		},
	// 		&ext.Panel{
	// 			HTML:   "My panel text...2",
	// 			Docked: "left",
	// 			Style:  "",
	// 			Flex:   1, // because default doesn't work
	// 		},
	// 	},
	// }
	// panel.Render()

	// panel.Debug()

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
