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
	conf := &app.WebConfig{
		Host: "127.0.0.1",
		Port: 8083,
	}
	a := app.NewWeb(conf)

	service.Register(a)

	d := <-a.Exit
	fmt.Printf("Done:%+v\n", d)
}
