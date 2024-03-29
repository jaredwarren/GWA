package gbt

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/gorilla/mux"
)

var (
	// for now make the stupid thing global
	web *Service
)

// Service basic web service
type Service struct {
	Name   string
	Mux    *mux.Router
	Exit   chan error
	Server *http.Server
	Home   *url.URL
}

// Application ...
type Application struct {
	// TODO: add xtype when marshal/unmarshal
	XType       string        `json:"xtype"`
	Name        string        `json:"name"`
	MainView    Renderer      `json:"mainview"`
	Nav         *Nav          `json:"nav,omitempty"`
	Head        *Head         `json:"head"`
	Width       int           `json:"width,omitempty"`
	Height      int           `json:"height,omitempty"`
	Controllers []*Controller `json:"-"`
	Exit        chan error    `json:"-"`
	service     *Service
}

// NewApp ...
func NewApp(name string) *Application {
	return &Application{
		XType: "app",
		Name:  name,
	}
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
		fmt.Printf("HTTP server listening on http://%s\n", addr)
		err := web.Server.ListenAndServe()
		if err != nil {
			fmt.Println("[E] http:", err)
		}
		a.Exit <- err
	}()
	a.service = web

	done := <-a.Exit
	a.Close()
	return done
}

// Home ...
func (a *Application) Home(w http.ResponseWriter, r *http.Request) {
	h := a.Render()
	_, err := fmt.Fprint(w, h)
	if err != nil {
		fmt.Println("[E] print err:", err)
	}
}

// Close ...
func (a *Application) Close() error {
	var err error
	if a.service != nil && a.service.Server != nil {
		a.service.Server.Close()
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

func (a *Application) ToHTML(path string) error {
	h := a.Render()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprint(f, h)
	if err != nil {
		return err
	}
	return nil
}

// Render ...
func (a *Application) Render() Stringer {
	items := Items{}
	// set controllers ui, so ui.Bind works
	for _, c := range a.Controllers {
		items = append(items, c)
	}

	// render main view
	items = append(items, a.MainView)
	// TODO: wrap items in new "body" element
	div := &DivContainer{
		ID:      "app",
		Classes: []string{"x-viewport"},
		Styles:  map[string]string{},
		Items:   items,
	}

	if a.Head.Title == "" {
		a.Head.Title = a.Name
	}

	return renderToHTML(`<!DOCTYPE html>
<html lang="en">
<head>
    {{.Head.Render}}
</head>
<body>
    {{.Nav.Render}}
    {{.Body.Render}}
</body>
</html>`, map[string]any{
		"Head": a.Head,
		"Nav":  a.Nav,
		"Body": div,
	})
}

// Update ...
func (a *Application) Update(r Renderer) error {
	// buf := new(bytes.Buffer)
	// err := r.Render(buf)
	// if err != nil {
	// 	return err
	// }
	// js := strings.Replace(buf.String(), "\n", "", -1)
	// js = strings.Replace(js, "'", "\\'", -1)
	// js = fmt.Sprintf(`document.getElementById('%s').outerHTML = '%s';`, r.GetID(), js)
	return nil
}

// Find ...
func (a *Application) Find(id string) Renderer {
	return find(id, a.MainView)
}

// UnmarshalJSON ...
func (a *Application) UnmarshalJSON(data []byte) error {
	var jApp map[string]interface{}
	if err := json.Unmarshal(data, &jApp); err != nil {
		fmt.Println(err)
		return err
	}

	if xtype, ok := jApp["xtype"]; !ok || xtype != "app" {
		return fmt.Errorf("root must be app")
	}

	if name, ok := jApp["name"]; ok {
		a.Name = name.(string)
	}

	mainview, ok := jApp["mainview"]
	if !ok || mainview == nil {
		return fmt.Errorf("mainview missing")
	}

	a.MainView = addChild(mainview)
	return nil
}

// TODO: move this to json file
//  - finish building other comps

func addChild(i interface{}) Renderer {
	xtype, ok := i.(map[string]interface{})["xtype"]
	if !ok {
		fmt.Printf("[error] %+v\n", i)
		return nil
	}
	switch xtype {
	case "panel":
		return buildPanel(i)
	case "button":
		// return buildButton(i)
	case "form":
		return buildForm(i)
	case "input":
		return buildInput(i)
	case "fieldset":
		return buildFieldset(i)
	case "tree":
		return buildTree(i)
	// case "treenode":
	// 	return buildTreeNode(i)
	default:
		fmt.Printf("Unknown Type: %+v\n", xtype)
	}
	return nil
}

func find(id string, node Renderer) Renderer {
	panic("oops")
	// ni, ok := node.(Parent)
	// if !ok {
	// 	return nil
	// }
	// items := ni.GetChildren()
	// for _, i := range items {
	// 	if i.GetID() == id {
	// 		return i
	// 	}
	// 	r := find(id, i)
	// 	if r != nil {
	// 		return r
	// 	}
	// }
	// return nil
}

// Dockable item that can be docked
type Dockable interface {
	GetDocked() string
	SetStyle(key, value string)
}

// Child ...
type Child interface {
	SetParent(p Renderer)
}

// Parent ...
type Parent interface {
	GetChildren() Items
}
