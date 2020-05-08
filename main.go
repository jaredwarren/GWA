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
		"\n☐ make header 'docked'",
		"\n☐ fix things so they work in test\n",
	}
)

func main() {
	fmt.Println("TODO:", TODO)

	// testJSON()
	serve()
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
