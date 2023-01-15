package gbt

import "fmt"

type Option[T any] func(T)

type Titler interface {
	SetTitle(string)
}

func WithTitle[F func(any)](title string) F {
	return func(a any) {
		x, ok := any(a).(Titler)
		if ok {
			x.SetTitle(title)
		} else {
			fmt.Printf("[E] not Titler(%T):%+v\n", a, a)
		}
	}
}

// Linker set an HREF value
type Linker interface {
	SetHref(href string)
}

func WithHref[F func(any)](href string) F {
	return func(a any) {
		x, ok := any(a).(Linker)
		if ok {
			x.SetHref(href)
		} else {
			fmt.Printf("[E] not Linker(%T):%+v\n", a, a)
		}
	}
}

// Linker set an HREF value
type Imager interface {
	SetImage(i *Image)
}

func WithImage[F func(any)](i *Image) F {
	return func(a any) {
		x, ok := any(a).(Imager)
		if ok {
			x.SetImage(i)
		} else {
			fmt.Printf("[E] not Imager(%T):%+v\n", a, a)
		}
	}
}

// Attributer ...
type Attributer interface {
	SetAttributes(a Attributes)
}

func WithAttributes[F func(any)](aa Attributes) F {
	return func(a any) {
		x, ok := any(a).(Attributer)
		if ok {
			x.SetAttributes(aa)
		} else {
			fmt.Printf("[E] not Attributer(%T):%+v\n", a, a)
		}
	}
}

// Classer ...
type Classer interface {
	SetClasses(c Classes)
}

func WithClasses[F func(any)](c Classes) F {
	return func(a any) {
		x, ok := any(a).(Classer)
		if ok {
			x.SetClasses(c)
		} else {
			fmt.Printf("[E] not Classer(%T):%+v\n", a, a)
		}
	}
}

// HTMLer ...
type HTMLer interface {
	SetInnerHTML(html Stringer)
}

func WithHTML[F func(any)](html Stringer) F {
	return func(a any) {
		x, ok := any(a).(HTMLer)
		if ok {
			x.SetInnerHTML(html)
		} else {
			fmt.Printf("[E] not HTMLer(%T):%+v\n", a, a)
		}
	}
}

// Itemer ...
type Itemer interface {
	SetItems(items Items)
}

func WithItems[F func(any)](items Items) F {
	return func(a any) {
		x, ok := any(a).(Itemer)
		if ok {
			x.SetItems(items)
		} else {
			fmt.Printf("[E] not Itemer(%T):%+v\n", a, a)
		}
	}
}

// Styler ...
type Styler interface {
	SetStyle(style string)
}

func WithStyle[F func(any)](style string) F {
	return func(a any) {
		x, ok := any(a).(Styler)
		if ok {
			x.SetStyle(style)
		} else {
			fmt.Printf("[E] not Styler(%T):%+v\n", a, a)
		}
	}
}

// Iconer Icon
type Iconer interface {
	SetIcon(i *Icon)
}

func WithIcon[F func(any)](i string, opts ...Option[any]) F {
	return func(a any) {
		x, ok := any(a).(Iconer)
		if ok {
			x.SetIcon(NewIcon(i, opts...))
		} else {
			fmt.Printf("[E] not Iconer(%T):%+v\n", a, a)
		}
	}
}

// TODO: replace with `WithImage`
func NavImage[F func(any)](image string) F {
	return func(a any) {
		x, ok := any(a).(*NavBrand)
		if ok {
			x.Image = &Image{
				Src:    image,
				Height: "20px",
			}
		}
	}
}

// Sizer Icon
type Sizer interface {
	SetSize(w, h string)
}

func WithSize[F func(any)](w, h string) F {
	return func(a any) {
		x, ok := any(a).(Sizer)
		if ok {
			x.SetSize(w, h)
		} else {
			fmt.Printf("[E] not Sizer(%T):%+v\n", a, a)
		}
	}
}
