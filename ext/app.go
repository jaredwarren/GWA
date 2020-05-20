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
	"strings"

	"github.com/gorilla/mux"
	"github.com/zserge/lorca"
)

var (
	// for now make the stupid thing global
	web *Service
	ui  lorca.UI
)

// Service basic web service
type Service struct {
	Name   string
	Mux    *mux.Router
	Exit   chan error
	Server *http.Server
	Home   *url.URL
	// Config      *WebConfig
}

// Application ...
type Application struct {
	Name string
	// Schemas map[string]Schema
	// Using   string // seledted schema
	MainView    Renderer
	Controllers []*Controller

	Width  int
	Height int

	Exit chan error

	service *Service
	cwd     string

	ui lorca.UI
}

// Launch ...
func (a *Application) Launch() error {
	a.Exit = make(chan error)
	// setup web service
	addr := "127.0.0.1:8083" // TODO: find open port
	// sudo lsof -i tcp:8083
	// kill -9 45590
	u, _ := url.Parse(fmt.Sprintf("http://%s", addr))
	web = &Service{
		Name: a.Name,
		Home: u,
	}

	// Interrupt handler (ctrl-c)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		done := <-signalChan
		a.Exit <- fmt.Errorf("%s", done)
	}()

	// Start Server
	web.Mux = mux.NewRouter()
	web.Mux.HandleFunc("/static/{filename:[a-zA-Z0-9\\.\\-\\_\\/]*}", FileServer)
	web.Mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// favicon
		// brew install imagemagick
		// convert -background none static/db.svg -define icon:auto-resize static/favicon.ico
		w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext("static/favicon.ico")))
		http.ServeFile(w, r, "static/favicon.ico")
	})
	web.Mux.HandleFunc("/health-check", HealthCheck).Methods("GET", "HEAD")
	web.Mux.HandleFunc("/", a.Home).Methods("GET")
	web.Mux.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		a.Close()
	}).Methods("GET")
	web.Server = &http.Server{
		Addr:    addr,
		Handler: web.Mux,
	}
	go func() {
		// TODO: add https, stuff...
		fmt.Printf("HTTP server listening on %q\n", addr)
		err := web.Server.ListenAndServe()
		if err != nil {
			fmt.Println("[E] http:", err)
		}
		a.Exit <- err
	}()
	a.service = web

	// Setup UI
	uiArgs := []string{}
	width := 500
	if a.Width > 0 {
		width = a.Width
	}
	height := 500
	if a.Height > 0 {
		height = a.Height
	}

	var err error
	ui, err = lorca.New("", "", width, height, uiArgs...)
	if err != nil {
		fmt.Println("[E] ui.New", err)
	}
	a.ui = ui

	// TODO:
	// use ui.Bind("counterAdd", c.Add)? for call backs???

	// run ui
	err = ui.Load(a.service.Home.String())
	if err != nil {
		fmt.Println("[E] ui.load", err)
	}

	go func() {
		<-ui.Done()
		a.Exit <- fmt.Errorf("UI Closed")
	}()

	done := <-a.Exit
	a.Close()
	return done
}

// Home ...
func (a *Application) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HOME")
	err := a.Render(w)
	if err != nil {
		fmt.Println("[[E]]:", err)
	}
}

// Close ...
func (a *Application) Close() error {
	fmt.Print("CLOSE")
	var err error
	if a.service != nil && a.service.Server != nil {
		a.service.Server.Close()
	}
	if a.ui != nil {
		err = a.ui.Close()
	}
	return err
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
	items := Items{}
	// set controllers ui, so ui.Bind works
	for _, c := range a.Controllers {
		c.ui = a.ui
		items = append(items, c)
	}

	// render main view
	items = append(items, a.MainView)
	div := &DivContainer{
		ID:      fmt.Sprintf("app"),
		Classes: []string{"x-viewport"},
		Styles:  map[string]string{},
		Items:   items,
	}
	buf := new(bytes.Buffer)
	err := renderDiv(buf, div)
	if err != nil {
		fmt.Println("[E] render:", err)
	}

	// render full html
	return renderTemplate(w, "base", &struct {
		Title string
		Body  template.HTML
	}{
		Title: a.Name,
		Body:  template.HTML(buf.String()),
	})
}

// Update ...
func (a *Application) Update(r Renderer) error {
	buf := new(bytes.Buffer)
	err := r.Render(buf)
	if err != nil {
		return err
	}
	js := strings.Replace(buf.String(), "\n", "", -1)
	js = strings.Replace(js, "'", "\\'", -1)
	js = fmt.Sprintf(`document.getElementById('%s').outerHTML = '%s';`, r.GetID(), js)
	a.ui.Eval(js)
	return nil
}

// Find ...
func (a *Application) Find(id string) Renderer {
	return find(id, a.MainView)
}

func find(id string, node Renderer) Renderer {
	ni, ok := node.(Parent)
	if !ok {
		return nil
	}
	items := ni.GetChildren()
	for _, i := range items {
		if i.GetID() == id {
			return i
		}
		r := find(id, i)
		if r != nil {
			return r
		}
	}
	return nil
}

// Dockable item that can be docked
type Dockable interface {
	GetDocked() string
}

// Child ...
type Child interface {
	SetParent(p Renderer)
}

// Parent ...
type Parent interface {
	GetChildren() Items
}
