package gbt

type Titler interface {
	SetTitle(string)
}

func WithTitle[F func(any)](title string) F {
	return func(a any) {
		x, ok := any(a).(Titler)
		if ok {
			x.SetTitle(title)
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
		}
	}
}

// Linker set an HREF value
type Iconer interface {
	SetIcon(i string)
}

func WithIcon[F func(any)](i string) F {
	return func(a any) {
		x, ok := any(a).(Iconer)
		if ok {
			x.SetIcon(i)
		}
	}
}
