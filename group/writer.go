package group

import (
	"bytes"
	"fmt"
	"io"
	"sync/atomic"
)

type Writer struct {
	w        io.Writer    // synchronized by active
	sections []*section   //
	offset   int          // [0, offset) is written, synchronized by active
	active   atomic.Int32 // [0, active) is closed
}

func NewWriter(w io.Writer, n int) *Writer {
	var (
		g        = &Writer{w: w}
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
	// It's possible for the sections to not be "drained" completely, even if
	// all of them are closed, depending on the order they're closed in.
	for _, s := range g.sections[g.active.Load():] {
		if err := s.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (g *Writer) drain() error {
	for _, s := range g.sections[g.offset : g.active.Load()+1] {
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
	closed atomic.Bool
}

func (s *section) Write(p []byte) (int, error) {
	if s.closed.Load() {
		panic(fmt.Sprintln("section: already closed", s.i))
	}
	switch active := s.g.active.Load(); {
	case s.i > active:
		return s.b.Write(p)
	case s.i == active:
		if err := s.g.drain(); err != nil {
			return 0, err
		}
		return s.g.w.Write(p)
	default: // s.i < active
		panic(fmt.Sprintln("section: impossible", s.i, active))
	}
}

func (s *section) Close() error {
	switch active := s.g.active.Load(); {
	case s.i > active:
		s.closed.Store(true)
		return nil
	case s.i == active:
		if err := s.g.drain(); err != nil {
			return err
		}
		s.closed.Store(true)
		active++
		// Try to advance the active section as much as possible, so we continue
		// to forward writes to the underlying writer in realtime.
		// Note: it's possible for the next open section's close to race with
		// this and "slip" away from us, so this is very much a best effort.
		for i, s := range s.g.sections[active:] {
			if !s.closed.Load() {
				active += int32(i)
				break
			}
		}
		s.g.active.Store(active)
		return nil
	case s.closed.Load():
		return nil
	default: // s.i < active
		panic(fmt.Sprintln("section: impossible", s.i, active))
	}
}
