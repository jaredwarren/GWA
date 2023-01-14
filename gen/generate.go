package gen

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	. "github.com/dave/jennifer/jen"
	"github.com/jaredwarren/goext/gbt"
	"golang.org/x/net/html"
)

func Generate(rawHtml string) {
	rawHtml = strings.TrimSpace(rawHtml)
	tkn := html.NewTokenizer(strings.NewReader(rawHtml))

	// e := getObj(tkn)
	e := getObj2(tkn)
	f, _ := os.Create("./out.go")
	defer f.Close()

	fmt.Fprintf(f, "%#v", e)
}

var TagNameRegex = regexp.MustCompile(`<(.+?)[ >]`)

func getName(t html.Token) string {
	parts := TagNameRegex.FindStringSubmatch(t.String())
	if len(parts) < 2 {
		fmt.Printf("~~~~~~~~~~~~~~~\n %+v\n\n", parts)
		return "ERROR"
	}
	return parts[1]
}

func getAttrs(a []html.Attribute) (gbt.Attributes, gbt.Classes) {
	att := gbt.Attributes{}
	class := gbt.Classes{}
	for _, aa := range a {
		if aa.Key == "class" {
			for _, v := range strings.Split(aa.Val, " ") {
				if v != "" {
					class = append(class, v)
				}
			}
		} else {
			att[aa.Key] = aa.Val
		}
	}
	return att, class
}

func getObj(tkn *html.Tokenizer) gbt.Renderer {
	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			return nil
		case html.StartTagToken:
			t := tkn.Token()
			name := getName(t)

			att, class := getAttrs(t.Attr)

			e := &gbt.Element{
				Name:       name,
				Attributes: att,
				Classes:    class,
			}

			// Some self closing nodes aren't formatted correctly
			if gbt.IsSelfClosing(name) {
				return e
			}

			// Append children
			for {
				i := getObj(tkn)
				if i == nil {
					break
				}
				e.Items = append(e.Items, i)
			}
			return e
		case html.TextToken:
			t := tkn.Token()
			return gbt.RawHTML(strings.TrimSpace(t.String()))
		case html.EndTagToken:
			return nil
		case html.SelfClosingTagToken:
			t := tkn.Token()
			name := getName(t)
			att, class := getAttrs(t.Attr)
			e := &gbt.Element{
				Name:       name,
				Attributes: att,
				Classes:    class,
			}
			return e
		case html.CommentToken:
			return nil
		default:
			fmt.Printf("~~~~~~~~~~~~~~~\n %+v\n\n", tt)
			t := tkn.Token()
			fmt.Printf("~~~~~~~~~~~~~~~t:\n %+v\n\n", t)
			fmt.Printf("~~~~~~~~~~~~~~~att:\n %+v\n\n", t.Attr)
			return nil
		}
	}
}

// jen
func getAttrs2(a []html.Attribute) (Code, Code) {

	att := Dict{}
	class := []Code{}
	for _, aa := range a {
		if aa.Key == "class" {
			for _, v := range strings.Split(aa.Val, " ") {
				if v != "" {
					class = append(class, Lit(v))
				}
			}
		} else {
			att[Lit(aa.Key)] = Lit(aa.Val)
		}
	}
	return Id("gbt.Attributes").Values(att), Id("gbt.Classes").Values(class...)
}

func getObj2(tkn *html.Tokenizer) Code {
	for {
		tt := tkn.Next()
		switch tt {
		case html.ErrorToken:
			return nil
		case html.StartTagToken:
			t := tkn.Token()
			name := getName(t)

			att, class := getAttrs2(t.Attr)

			// Some self closing nodes aren't formatted correctly
			if gbt.IsSelfClosing(name) {
				e := Op("&").Id("gbt.Element").Values(Dict{
					Id("Name"):       Lit(name),
					Id("Attributes"): att,
					Id("Classes"):    class,
				})
				return e
			}

			// Append children
			items := []Code{}
			for {
				i := getObj2(tkn)
				if i == nil {
					break
				}
				items = append(items, i)
			}

			e := Op("&").Id("gbt.Element").Values(Dict{
				Id("Name"):       Lit(name),
				Id("Attributes"): att,
				Id("Classes"):    class,
				Id("Items"):      Id("gbt.Items").Values(items...),
			})

			return e
		case html.TextToken:
			t := tkn.Token()
			str := strings.TrimSpace(t.String())
			if str != "" {
				e := Id("gbt.RawHTML").Call(Lit(str))
				return e
			}
		case html.EndTagToken:
			return nil
		case html.SelfClosingTagToken:
			t := tkn.Token()
			name := getName(t)
			att, class := getAttrs2(t.Attr)
			e := Op("&").Id("gbt.Element").Values(Dict{
				Id("Name"):       Lit(name),
				Id("Attributes"): att,
				Id("Classes"):    class,
			})
			return e
		case html.CommentToken:
			return nil
		default:
			fmt.Printf("~~~~~~~~~~~~~~~\n %+v\n\n", tt)
			t := tkn.Token()
			fmt.Printf("~~~~~~~~~~~~~~~t:\n %+v\n\n", t)
			fmt.Printf("~~~~~~~~~~~~~~~att:\n %+v\n\n", t.Attr)
			return nil
		}
	}
}
