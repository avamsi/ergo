package group

import (
	"bytes"
	"io"
)

type Writer struct {
	w        io.Writer
	sections []*bytes.Buffer
	closed   bool
}

func NewWriter(w io.Writer, n int) *Writer {
	return &Writer{
		w:        w,
		sections: make([]*bytes.Buffer, 0, n),
	}
}

func (g *Writer) AddSection() io.Writer {
	if g.closed {
		panic("group.Writer: already closed")
	}
	var b bytes.Buffer
	g.sections = append(g.sections, &b)
	return &b
}

func (g *Writer) Close() error {
	if g.closed {
		return nil
	}
	for _, b := range g.sections {
		if _, err := b.WriteTo(g.w); err != nil {
			return err
		}
	}
	g.closed = true
	return nil
}
