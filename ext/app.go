package ext

import (
	"bytes"
	"encoding/json"
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
	XType       string        `json:"xtype"`
	Name        string        `json:"name"`
	MainView    Renderer      `json:"mainview"`
	Width       int           `json:"width,omitempty"`
	Height      int           `json:"height,omitempty"`
	Controllers []*Controller `json:"-"`
	Exit        chan error    `json:"-"`
	service     *Service      `json:"-"`
	cwd         string        `json:"-"`
	ui          lorca.UI      `json:"-"`
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

// // MarshalJSON ...
// func (a *Application) MarshalJSON() ([]byte, error) {
// 	result := map[string]interface{}{}
// 	e := reflect.ValueOf(a).Elem()
// 	for i := 0; i < e.NumField(); i++ {
// 		varName := lowerInitial(e.Type().Field(i).Name)
// 		if e.Field(i).CanInterface() {
// 			result[varName] = e.Field(i).Interface()
// 		}
// 	}
// 	return json.Marshal(result)
// 	return json.Marshal(&struct {
// 		XType    string   `json:"xtype"`
// 		Width    int      `json:"width,omitempty"`
// 		Height   int      `json:"height,omitempty"`
// 		Name     string   `json:"name"`
// 		MainView Renderer `json:"mainview"`
// 	}{
// 		XType:    "app",
// 		Name:     a.Name,
// 		Width:    a.Width,
// 		Height:   a.Height,
// 		MainView: a.MainView,
// 	})
// }

// UnmarshalJSON ...
func (a *Application) UnmarshalJSON(data []byte) error {
	var jApp map[string]interface{}
	if err := json.Unmarshal(data, &jApp); err != nil {
		fmt.Println(err)
		return err
	}

	if xtype, ok := jApp["xtype"]; !ok || xtype != "app" {
		return fmt.Errorf("Root must be app")
	}

	if name, ok := jApp["name"]; ok {
		a.Name = name.(string)
	}

	// fmt.Printf("%+v\n", a)
	mainview, ok := jApp["mainview"]
	if !ok || mainview == nil {
		return fmt.Errorf("mainview missing")
	}

	a.MainView = addChild(mainview)

	fmt.Printf("%+v\n", a.MainView)

	// for k, v := range mainview.(map[string]interface{}) {
	// 	fmt.Printf("  %s: %+v\n", k, v)
	// }
	return nil

	// b.Price, _ = v[0].(string)
	// b.Size, _ = v[1].(string)
	// b.NumOrders = int(v[2].(float64))

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
		return buildButton(i)
	case "form":
		return buildForm(i)
	case "input":
		return buildInput(i)
	case "fieldset":
		return buildFieldset(i)
	case "tree":
		return buildTree(i)
	case "treenode":
		return buildTreeNode(i)
	}
	fmt.Printf("[error] %+v\n", xtype)
	return nil
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
