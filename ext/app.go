package ext

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/zserge/lorca"
)

// Application ...
type Application struct {
	Name string
	// Schemas map[string]Schema
	// Using   string // seledted schema
	MainView Renderer

	Width  int
	Height int

	Exit chan error

	service *Service
	cwd     string
	ui      lorca.UI
}

// Launch ...
func (a *Application) Launch() error {

	// setup web service
	addr := "127.0.0.1:8083" // TODO: find open port
	u, _ := url.Parse(fmt.Sprintf("http://%s", addr))
	app := &Service{
		Name: a.Name,
		Exit: make(chan error),
		Home: u,
	}

	// Interrupt handler (ctrl-c)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		done := <-signalChan
		app.Exit <- fmt.Errorf("%s", done)
	}()

	// Start Server
	app.Mux = mux.NewRouter()
	app.Mux.HandleFunc("/static/{filename:[a-zA-Z0-9\\.\\-\\_\\/]*}", FileServer)
	app.Mux.HandleFunc("/health-check", HealthCheck).Methods("GET", "HEAD")
	app.Mux.HandleFunc("/", a.Home).Methods("GET")
	app.Server = &http.Server{
		Addr:    addr,
		Handler: app.Mux,
	}
	go func() {
		// TODO: add https, stuff...
		fmt.Printf("HTTP server listening on %q\n", addr)
		app.Exit <- app.Server.ListenAndServe()
	}()

	a.service = app

	done := <-a.Exit
	return done
}

// Home serves a file with mime type header
func (a *Application) Home(w http.ResponseWriter, r *http.Request) {
	err := a.Render(w)
	if err != nil {
		fmt.Println("[[E]]:", err)
	}
}

// FileServer serves a file with mime type header
func FileServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["filename"]
	w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(file)))
	http.ServeFile(w, r, "./static/"+file)
}

// HealthCheck return ok
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// Render ...
func (a *Application) Render(w io.Writer) error {
	buf := new(bytes.Buffer)

	// render main view
	div := &DivContainer{
		ID:      fmt.Sprintf("app"),
		Classes: []string{"x-viewport"},
		Items:   Items{a.MainView},
	}
	err := renderDiv(buf, div)
	if err != nil {
		fmt.Println("[E] render:", err)
	}

	// render full html
	return render(w, "base", &struct {
		Title string
		Body  template.HTML
	}{
		Title: a.Name,
		Body:  template.HTML(buf.String()),
	})
}

// TabPanel ...
type TabPanel struct {
	// Name   string
	// Tables []Table
}

// List ...
type List struct {
	Title   string
	Store   *Store
	Columns []*Column
}

// Column ...
type Column struct {
	Text      string
	DataIndex string
	Width     int
}

// Store ...
type Store struct {
}

// Items ...
type Items []Renderer

// Render ...
func (i Items) Render(w io.Writer) error {
	for _, item := range i {
		err := item.Render(w)
		if err != nil {
			return err
		}
	}
	return nil
}

// Renderer an item that can be rendered
type Renderer interface {
	Render(w io.Writer) error
}

// Render a template
func renderTemplate(w io.Writer, t string, data interface{}) error {
	funcMap := template.FuncMap{
		"Render": func(item Renderer) template.HTML {
			buf := new(bytes.Buffer)
			item.Render(w)
			return template.HTML(buf.String())
		},
	}
	tpl, err := template.New("base").Funcs(funcMap).ParseFiles(fmt.Sprintf("templates/%s.html", t))
	if err != nil {
		fmt.Printf("[E] %s parse error:%s\n", t, err)
	}
	templates := template.Must(tpl, err)
	return templates.ExecuteTemplate(w, "base", data)
}

// Dockable item that can be docked
type Dockable interface {
	GetDocked() string
}

// Debug print crap out
func Debug(p Renderer) {
	d(p, 0)
}

func d(p Renderer, depth int) {
	pd(depth)
	typeof := reflect.TypeOf(p).String()

	switch typeof {
	case "*ext.Panel":
		fmt.Print("| ", "Panel", p.(*Panel).ID)
		fmt.Println("  html:", p.(*Panel).HTML)
		pd(depth)
		fmt.Println("  style:", p.(*Panel).Styles)
		for _, i := range p.(*Panel).Items {
			d(i, depth+1)
		}
	case "*ext.Innerhtml":
		fmt.Print("| ", "Innerhtml", p.(*Innerhtml).ID)
		fmt.Println("  html:", p.(*Innerhtml).HTML)
	case "*ext.Layout":
		fmt.Print("| ", "Layout", p.(*Layout).ID)
		fmt.Println(":", p.(*Layout).Type)
		for _, i := range p.(*Layout).Items {
			d(i, depth+1)
		}
	case "*ext.Body":
		fmt.Print("| ", "Body", p.(*Body).ID)
		fmt.Println("")
		for _, i := range p.(*Body).Items {
			d(i, depth+1)
		}
	case "*ext.Header":
		fmt.Print("| ", "Header", p.(*Header).ID)
		fmt.Println("::", p.(*Header).Title)
	default:
		fmt.Print("| ?", typeof)
		fmt.Println(" ?")
	}
}

// Print depth spaces
func pd(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
}

// LayoutItems ...
func LayoutItems(oi Items) Items {
	// if there's only one item there's nothing to layout
	if len(oi) < 2 {
		return oi
	}

	bodyItems := Items{}
	layout := &Layout{
		Items: Items{},
	}
	var di Renderer
	for _, i := range oi {
		// already find docked item, append rest to body and move on
		if di != nil {
			bodyItems = append(bodyItems, i)
			continue
		}

		// Look for docked item
		// if not dockable add to items, else add to body
		d, ok := i.(Dockable)
		if ok {
			docked := d.GetDocked()
			if docked != "" {
				switch docked {
				case "top":
					layout.Type = "hbox"
					layout.Pack = "start"
					// i goes first
				case "bottom":
					layout.Type = "hbox"
					layout.Pack = "end"
					// i goes last
				case "left":
					layout.Type = "vbox"
					layout.Pack = "start"
					// i goes first
				case "right":
					layout.Type = "vbox"
					layout.Pack = "end"
					// i goes last
				default:
					// what to do
				}
				layout.Align = "start" // should always be start?
				di = i
			} else {
				bodyItems = append(bodyItems, i)
			}
		} else {
			bodyItems = append(bodyItems, i)
		}
	}

	// Nothing to layout
	if di != nil {
		layout.Items = Items{di}
		// Add body items to layout
		if len(bodyItems) > 0 {
			layout.Items = append(layout.Items, NewBody(bodyItems))
		} // else what to do????? add a blank one?
	} else {
		return oi
	}

	// add rest
	return Items{layout}
}

func renderDiv(w io.Writer, data *DivContainer) error {
	return render(w, `<div id="{{.ID}}" class="{{range $c:= $.Classes}}{{$c}} {{end}}" style="{{range $k, $s:= $.Styles}}{{$k}}:{{$s}}; {{end}}">
			{{range $item := $.Items}}
			{{Render $item}}
			{{end}}</div>`, data)
}

func render(w io.Writer, t string, data interface{}) error {
	tpl, err := template.New("base").Funcs(template.FuncMap{
		"Render": func(item Renderer) template.HTML {
			buf := new(bytes.Buffer)
			item.Render(w)
			return template.HTML(buf.String())
		},
	}).Parse(t)
	if err != nil {
		fmt.Printf("[E] parse error:%s\n", err)
	}

	templates := template.Must(tpl, err)
	err = templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Printf("[E] execute error:%s\n", err)
	}
	return err
}

// DivContainer gerneric div container
type DivContainer struct {
	ID            string
	Items         Items
	ContainerType string
	Classes       []string
	Styles        map[string]string
}
