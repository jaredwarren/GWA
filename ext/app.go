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

	"github.com/gorilla/mux"
	"github.com/zserge/lorca"
)

// Service basic web service
type Service struct {
	Name   string
	Mux    *mux.Router
	Exit   chan error
	Server *http.Server
	// Controllers []Controller
	Home *url.URL
	// Config      *WebConfig
}

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
	fmt.Println("LAUNCH")
	a.Exit = make(chan error)
	// setup web service
	addr := "127.0.0.1:8083" // TODO: find open port
	// sudo lsof -i tcp:8083
	// kill -9 45590
	u, _ := url.Parse(fmt.Sprintf("http://%s", addr))
	web := &Service{
		Name: a.Name,
		Home: u,
	}

	// Interrupt handler (ctrl-c)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		done := <-signalChan
		fmt.Println("ctrl-c", done)
		a.Exit <- fmt.Errorf("%s", done)
	}()

	// Start Server
	web.Mux = mux.NewRouter()
	web.Mux.HandleFunc("/static/{filename:[a-zA-Z0-9\\.\\-\\_\\/]*}", FileServer)
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
		fmt.Println("[E] http:", err)
		a.Exit <- err
	}()
	a.service = web

	// fmt.Println("Waiting for exit")
	// done := <-a.Exit
	// fmt.Println("DONE:", done)
	// a.Close()
	// return done

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

	ui, err := lorca.New("", "", width, height, uiArgs...)
	if err != nil {
		fmt.Println("[E] ui.New", err)
	}
	a.ui = ui

	// TODO:
	// use ui.Bind("counterAdd", c.Add)? for call backs???

	// run ui
	fmt.Println("LOAD:", a.service.Home.String())
	err = ui.Load(a.service.Home.String())
	if err != nil {
		fmt.Println("[E] ui.load", err)
	}

	go func() {
		fmt.Println("waiting for ui")
		x := <-ui.Done()
		fmt.Printf("X:%+v\n", x)
		a.Exit <- fmt.Errorf("UI Closed")
	}()

	fmt.Println("Waiting for exit")
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
	// render main view
	div := &DivContainer{
		ID:      fmt.Sprintf("app"),
		Classes: []string{"x-viewport"},
		Items:   Items{a.MainView},
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

// Dockable item that can be docked
type Dockable interface {
	GetDocked() string
}

// Child item that can be docked
type Child interface {
	SetParent(p Renderer)
}
