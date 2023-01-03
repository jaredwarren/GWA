package v2

import "fmt"

type App struct {
	Children []*Child
}

func (a *App) Render() {
	fmt.Println("Render")
	// TODO: for each child render inside "a"
}

type Child struct {
}

func Load() {
	a := &App{}
	a.Render()
	// generate base html
	//
}

// https://asdf.com/home
// -> return base html
//
