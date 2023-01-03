package ext

import "io"

type Head struct {
	Items Items
}

type HeadElement Renderer

func (h *Head) Render(w io.Writer) error {
	// TODO: wrap in head once i clean up base.html
	for _, hi := range h.Items {
		err := hi.Render(w)
		if err != nil {
			return err
		}
	}
	return nil
}
