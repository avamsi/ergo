package group

import (
	"bytes"
	"io"
)

type Writer struct {
	w        io.Writer
	sections []*section
	closed   bool
}

func NewWriter(w io.Writer, n int) *Writer {
	return &Writer{
		w:        w,
		sections: make([]*section, 0, n),
	}
}

func (g *Writer) AddSection() io.Writer {
	if g.closed {
		panic("group: already closed")
	}
	s := &section{g: g}
	g.sections = append(g.sections, s)
	return s
}

func (g *Writer) Close() error {
	for _, b := range g.sections {
		if err := b.close(); err != nil {
			return err
		}
	}
	g.closed = true
	return nil
}

type section struct {
	g      *Writer
	b      bytes.Buffer
	closed bool
}

func (s *section) Write(p []byte) (int, error) {
	if s.closed {
		panic("section: already closed")
	}
	return s.b.Write(p)
}

func (s *section) close() error {
	if _, err := s.b.WriteTo(s.g.w); err != nil {
		return err
	}
	s.closed = true
	return nil
}
