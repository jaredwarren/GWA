package ext

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

var (
	tableID = 0
)

// Table ...
type Table struct {
	XType      string      `json:"xtype"`
	ID         string      `json:"id,omitempty"`
	Width      int         `json:"width,omitempty"`
	Height     int         `json:"height,omitempty"`
	HideHeader bool        `json:"hideHeader,omitempty"`
	Docked     string      `json:"docked,omitempty"`
	Classes    Classes     `json:"classes,omitempty"`
	Styles     Styles      `json:"styles,omitempty"`
	Parent     Renderer    `json:"-"`
	Columns    Columns     `json:"columns,omitempty"`
	Title      string      `json:"title,omitempty"`
	Data       Rows        `json:"data,omitempty"`
	Header     TableHeader `json:"header,omitempty"`
	Footer     TableFooter `json:"footer,omitempty"`
}

// Render ...
func (t *Table) Render(w io.Writer) error {
	if t.ID == "" {
		t.ID = nextTableID()
	}
	t.Classes.Add("pure-table")
	t.Classes.Add("pure-table-horizontal")

	// Body
	bodyItems := Items{}
	for _, r := range t.Data {
		row := &Element{
			Name:  "tr",
			Items: Items{},
		}

		if len(t.Header) > 0 {
			for _, hr := range t.Header {
				for _, hc := range hr {
					v, ok := r[hc.DataIndex]
					if !ok {
						continue
					}

					var cell *Element
					i, ok := v.(Renderer)
					if ok {
						cell = &Element{
							Name:  "td",
							Items: Items{i},
						}
					} else {
						cell = &Element{
							Name:      "td",
							Innerhtml: template.HTML(fmt.Sprintf("%v", v)),
						}
					}
					row.Items = append(row.Items, cell)
				}
			}
		} else {
			for _, v := range r {
				var cell *Element
				i, ok := v.(Renderer)
				if ok {
					cell = &Element{
						Name:  "td",
						Items: Items{i},
					}
				} else {
					cell = &Element{
						Name:      "td",
						Innerhtml: template.HTML(fmt.Sprintf("%v", v)),
					}
				}
				row.Items = append(row.Items, &Element{
					Name:  "td",
					Items: Items{cell},
				})
			}
		}
		bodyItems = append(bodyItems, row)
	}
	if len(bodyItems) == 0 {
		emptyTableEl := &Element{
			Name:      "div",
			Innerhtml: "No Data",
		}
		return emptyTableEl.Render(w)
	}
	body := &Element{
		Name:  "tbody",
		Items: bodyItems,
	}

	// HEAD
	thead := &Element{
		Name: "thead",
	}
	if !t.HideHeader {
		headerRows := Items{}
		if len(t.Header) > 0 {
			for _, hr := range t.Header {
				headerCells := Items{}
				for _, hc := range hr {
					headerCells = append(headerCells, &Element{
						Name:       "th",
						Innerhtml:  hc.Innerhtml,
						Items:      hc.Items,
						Attributes: hc.Attributes,
					})
				}
				headerRows = append(headerRows, &Element{
					Name:  "tr",
					Items: headerCells,
				})
			}
		} else if len(t.Data) > 0 {
			// create header from data[0]
			hr := t.Data[0]
			headerCells := Items{}
			for k := range hr {
				label := strings.Title(k)
				headerCells = append(headerCells, &Element{
					Name:      "th",
					Innerhtml: template.HTML(label),
				})
			}
			headerRows = append(headerRows, &Element{
				Name:  "tr",
				Items: headerCells,
			})
		} // else no data, no header

		thead.Items = headerRows
	}

	// FOOT
	footerRows := Items{}
	if len(t.Footer) > 0 {
		for _, hr := range t.Footer {
			footerCells := Items{}
			for _, hc := range hr {
				footerCells = append(footerCells, &Element{
					Name:       "th",
					Innerhtml:  hc.Innerhtml,
					Attributes: hc.Attributes,
				})
			}
			footerRows = append(footerRows, &Element{
				Name:  "tr",
				Items: footerCells,
			})
		}
	}
	tfoot := &Element{
		Name:  "tfoot",
		Items: footerRows,
	}

	navEl := &Element{
		Name: "table",
		Attributes: Attributes{
			"id":    template.HTMLAttr(t.ID),
			"class": t.Classes.ToAttr(),
			"style": t.Styles.ToAttr(),
		},
		Items: Items{
			thead,
			body,
			tfoot,
		},
	}
	return navEl.Render(w)
}

// GetID ...
func (t *Table) GetID() string {
	return t.ID
}

// SetParent ...
func (t *Table) SetParent(p Renderer) {
	t.Parent = p
}

// GetDocked ...
func (t *Table) GetDocked() string {
	return t.Docked
}

// SetStyle ...
func (t *Table) SetStyle(key, value string) {
	if t.Styles == nil {
		t.Styles = map[string]string{}
	}
	t.Styles[key] = value
}

/**
* Header
 */

// TableHeader ..
type TableHeader []HeaderRow

// HeaderRow ...
type HeaderRow []HeaderCell

// HeaderCell ...
type HeaderCell struct {
	Innerhtml  template.HTML
	Attributes map[string]template.HTMLAttr
	DataIndex  string
	Items      Items
}

/**
* Footer
 */

// TableFooter ..
type TableFooter []FooterRow

// FooterRow ...
type FooterRow []FooterCell

// FooterCell ...
type FooterCell struct {
	Innerhtml  template.HTML
	Attributes map[string]template.HTMLAttr
	DataIndex  string
	Items      Items
}

/**
* Columns
 */

// Columns ...
type Columns []*Column

// Column ...
type Column struct {
	Text      string
	DataIndex string
	Width     int
}

/**
* Row
 */

// Rows ...
type Rows []Row

// Row ...
type Row map[string]interface{}

// Cell ...
type Cell struct {
	Innerhtml  template.HTML
	Attributes map[string]template.HTMLAttr
}

/**
* ???
 */

func nextTableID() string {
	id := fmt.Sprintf("table-%d", tableID)
	tableID++
	return id
}
