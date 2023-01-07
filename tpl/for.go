package tpl

// Use this if app produes a template instead of raw html
type For struct {
}

func (f *For) Render() string {
	return `{{range $index, $item := .Items}}
		  {{$item.Render}}
		{{end}}`
}
