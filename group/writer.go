package group

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"sync/atomic"

	"github.com/avamsi/ergo"
)

type Writer struct {
	w        io.Writer    // synchronized by active
	n        int32        //
	sections []*section   //
	offset   int          // [0, offset) is written, synchronized by active
	m        sync.Mutex   // synchronizes active and section.closed
	active   atomic.Int32 // [0, active) is closed
}

func NewWriter(w io.Writer, n int) *Writer {
	var (
		g        = &Writer{w: w, n: int32(n)}
		sections = make([]*section, n)
	)
	for i := range sections {
		sections[i] = &section{g: g, i: int32(i)}
	}
	g.sections = sections
	return g
}

func (g *Writer) Section(i int) io.WriteCloser {
	return g.sections[i]
}

func (g *Writer) Close() error {
	for _, s := range g.sections[g.active.Load():] {
		if err := s.Close(); err != nil {
			return err
		}
	}
	// It's possible for the sections to not be drained completely, even if
	// all of them are closed, depending on the order they're closed in.
	return g.drain(g.n)
}

func (g *Writer) drain(till int32) error {
	for _, s := range g.sections[g.offset:till] {
		if _, err := s.b.WriteTo(g.w); err != nil {
			return err
		}
		g.offset++ // (including g.w above) synchronized by active
	}
	return nil
}

type section struct {
	g      *Writer
	i      int32
	b      bytes.Buffer
	closed bool
}

func (s *section) Write(p []byte) (int, error) {
	if s.closed {
		ergo.Panic("section: already closed", s.i)
	}
	switch active := s.g.active.Load(); {
	case s.i > active:
		return s.b.Write(p)
	case s.i == active:
		if err := s.g.drain(s.i + 1); err != nil {
			return 0, err
		}
		return s.g.w.Write(p)
	default: // s.i < active
		panic(fmt.Sprint("section: impossible", s.i, active))
	}
}

func (s *section) Close() error {
	s.g.m.Lock()
	unlock := true
	defer func() {
		if unlock {
			s.g.m.Unlock()
		}
	}()
	switch active := s.g.active.Load(); {
	case s.i > active:
		s.closed = true
		return nil
	case s.i == active:
		s.g.m.Unlock()
		if err := s.g.drain(s.i + 1); err != nil {
			unlock = false
			return err
		}
		s.g.m.Lock()
		s.closed = true
		active++
		for active < s.g.n && s.g.sections[active].closed {
			active++
		}
		s.g.active.Store(active)
		return nil
	case s.closed:
		return nil
	default: // s.i < active
		panic(fmt.Sprint("section: impossible", s.i, active))
	}
}
