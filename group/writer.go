package group

import (
	"bytes"
	"io"
)

type Writer struct {
	delegate io.Writer
	sections chan *bytes.Buffer
}

func NewWriter(delegate io.Writer, limit int) *Writer {
	return &Writer{
		delegate: delegate,
		sections: make(chan *bytes.Buffer, limit),
	}
}

func (w *Writer) NewSection() io.Writer {
	var b bytes.Buffer
	w.sections <- &b
	return &b
}

func (w *Writer) Close() error {
	close(w.sections)
	for b := range w.sections {
		if _, err := io.Copy(w.delegate, b); err != nil {
			return err
		}
	}
	return nil
}
