package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/jaredwarren/goext/ext"
)

var (
	// TODO ...
	TODO = []string{
		"\n☐ json to panel (use xtype)",         // big problem with functions
		"\n☐ make header 'docked'",              // I think this is done, but header template needs cleaned up
		"\n☐ fix things so they work in test\n", // almost done
		// "\n☐ make 'app' class that's full-screen, merge with native app\n",
		// "\n☐ figure a good way to load item from file\n", // done?
		"\n☐ figure a way to re-attach to web service (keep running in backgorund (might have to be saparate app))\n",
		"\n☐ Figure a way to update single panel without page reload!!!\n",
		"\n☐ figure way to make controller work -> pass ui to controller, ui.bind?\n",
		"\n☐ store (get data from ui.bind->ui.eval? or ajax, something...)\n",
		"\n☐ fix handler problem with type (might have to make all the same!), wonder if I can override json marshaller?, if not then what?\n",
		"\n☐ create FORM and figure good way to submit to controller\n",
		"\n☐ create multiple sessions/instances of app\n",
		"\n☐ create multiple windows\n",
		"\n☐ save app state, have to do manually\n",
		"\n☐ Look for template e.g. panel.html\n",
		"\n☐ \n",
		"\n☐ replace all woff2 in pro.min.css https://kit-pro.fontawesome.com/releases/v5.13.0/webfonts/pro-fa-brands-400-5.12.0.woff2\n",
		"\n☐ \n",
	}

	// this is here to show that objects can be in a saparate file
	mainController = &ext.Controller{
		Handlers: ext.Handlers{
			"btnClick": func(id string) {
				fmt.Print("Button Clicked:")
				fmt.Printf("   %+v\n", id)

				// Button update test
				btn := app.Find(id)
				if btn != nil {
					btn.(*ext.Button).Text = "Clicked!!!"
					app.Update(btn)
				}

				// Update Tree Test
				t := app.Find("tree-0")
				if t != nil {
					t.(*ext.Tree).Root.Text = "UPDATED"
					app.Update(t)
				}
			},
		},
		FormHandlers: ext.FormHandlers{
			"formSubmit": func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("submit....")
			},
		},
	}
	app *ext.Application
)

func main() {
	// fmt.Println("TODO:", TODO)

	// other2()
	// return

	// // //
	app = buildFromJSON()
	done := app.Launch()
	if done != nil {
		fmt.Println("Something Happened, Bye!", done)
	} else {
		fmt.Println("Good Bye!")
	}
	return

	//na

	// // //

	app = &ext.Application{
		Name: "my app",
		Controllers: []*ext.Controller{
			mainController,
		},
		MainView: &ext.Panel{
			Title:  "Panel Title!",
			Shadow: true,
			Layout: "hbox",
			HTML:   "test",
			Items: []ext.Renderer{
				&ext.Panel{
					HTML:   "My panel text...1",
					Docked: "top",
				},
				&ext.Panel{
					HTML:   "My panel text...3",
					Docked: "left",
				},
				&ext.Panel{
					HTML:   "My panel text...4",
					Docked: "bottom",
				},
				&ext.Button{
					Text:    "Click Here",
					Handler: "btnClick",
				},
				&ext.Button{
					Text: "2 Here",
					HandlerFn: func(id string) {
						fmt.Print("Button 2 Clicked:")
						fmt.Printf("   %+v\n", id)

						// Button update test
						btn := app.Find(id)
						if btn != nil {
							btn.(*ext.Button).Text = "Clicked!!!"
							app.Update(btn)
						}

						// Update Tree Test
						t := app.Find("tree-0")
						if t != nil {
							t.(*ext.Tree).Root.Text = "UPDATED"
							app.Update(t)
						}
					},
				},
				&ext.Form{
					// Text:    "Click Here",
					// Handler: "btnClick",
					// Method: "post",
					// Action: "submit",
					Handler: "formSubmit",
					// Handler: func(w http.ResponseWriter, r *http.Request) {
					// 	fmt.Println("submit....")
					// },

					Items: []ext.Renderer{
						&ext.Fieldset{
							Legend: "Form Legend",
							Items: []ext.Renderer{
								&ext.Input{
									Label: "User Name:",
									Name:  "username",
									Type:  "text",
								},
								&ext.Input{
									Label: "Send:",
									Name:  "submit",
									Type:  "submit",
								},
							},
						},
					},
				},
				&ext.Tree{
					Docked:     "right",
					ShowRoot:   true,
					BranchIcon: "",
					LeafIcon:   "",
					Root: &ext.TreeNode{
						Text: "root",
						// IconClass: "fas fa-folder-open",
						Children: []*ext.TreeNode{
							&ext.TreeNode{
								Text:      "c1",
								IconClass: "fas fa-fighter-jet",
							},
							&ext.TreeNode{
								Text:      "c2",
								IconClass: "fad fa-acorn",
							},
							&ext.TreeNode{
								Text:      "c3",
								IconClass: "fad fa-arrow-alt-from-right",
							},
							&ext.TreeNode{
								Text:      "c4",
								IconClass: "fad fa-tree-palm",
							},
							&ext.TreeNode{
								Text: "c2",
								Children: []*ext.TreeNode{
									&ext.TreeNode{
										Text:     "c2c1",
										Children: []*ext.TreeNode{},
									},
								},
							},
							&ext.TreeNode{
								Text: "c3",
								Children: []*ext.TreeNode{
									&ext.TreeNode{
										Text:     "c3c1",
										Children: []*ext.TreeNode{},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// app -> json
	if false {
		b, err := json.MarshalIndent(app, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
		}
		ioutil.WriteFile("./app.json", b, 0644)
		return
	}

	done := app.Launch()
	if done != nil {
		fmt.Println("Something Happened, Bye!", done)
	} else {
		fmt.Println("Good Bye!")
	}
}

func other() {
	f, err := os.Open("static/css/pro.min.css")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	// \(\.\.\/webfonts\/([a-z\-0-9\.])\.woff2\)

	// src:url\((.+?)\)
	// ../webfonts/pro-fa-light-300-5.0.1.woff2
	// re := regexp.MustCompile(`\.\.\/webfonts\/(.+?)\.woff2`)
	re := regexp.MustCompile(`url\((.+?)\)`)

	line := 1
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for scanner.Scan() {
		text := scanner.Text()
		// fmt.Println(text)
		matches := re.FindAllString(text, -1)
		if len(matches) > 0 {
			fmt.Println(matches)
		}

		// if strings.Contains(scanner.Text(), "yourstring") {
		// 	return line, nil
		// }

		line++
		// if line > 100 {
		// 	return
		// }
	}

	fmt.Println(line)

	if err := scanner.Err(); err != nil {
		// Handle the error
	}
}

func other2() {
	data, _ := ioutil.ReadFile("static/css/pro.min.css")
	/* ... omitted error check..and please add ... */
	/* find index of newline */
	file := string(data)
	/* func Split(s, sep string) []string */
	temp := strings.Split(file, "\n")

	// re := regexp.MustCompile(`url\((.+?)\)`)
	re := regexp.MustCompile(`url\((.+?)\.woff2\)`)

	for _, item := range temp {
		// matches := re.FindAllString(item, -1)
		matches := re.FindAllStringSubmatch(item, -1)
		if len(matches) > 0 {
			relPath := matches[0][1]

			url := strings.Replace(relPath, "../", "https://kit-pro.fontawesome.com/releases/v5.13.0/", -1) + ".woff2"
			file := strings.Replace(relPath, "../webfonts/", "static/webfonts/", -1) + ".woff2"

			if !fileExists(file) {
				fmt.Println("DOWN", file, url)
				// err := DownloadFile(file, url)
				// if err != nil {
				// 	fmt.Println(err)
				// }
			}
			// https://kit-pro.fontawesome.com/releases/v5.13.0/

			// time.Sleep(3 * time.Second)
		}

		// https://kit-pro.fontawesome.com/releases/v5.13.0/webfonts/pro-fa-solid-900-5.11.2
		// https://kit-pro.fontawesome.com/releases/v5.13.0/webfonts/pro-fa-light-300-5.0.9.woff2

	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func buildFromJSON() *ext.Application {
	dat, err := ioutil.ReadFile("./app.json")
	na := &ext.Application{}
	err = json.Unmarshal(dat, na)
	if err != nil {
		fmt.Println(err)
	}

	na.Controllers = []*ext.Controller{
		mainController,
	}

	b, err := json.MarshalIndent(na, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	ioutil.WriteFile("./app2.json", b, 0644)
	return nil

	d := na.Launch()
	if d != nil {
		fmt.Println("Something Happened, Bye!", d)
	} else {
		fmt.Println("Good Bye!")
	}

	return na
}
