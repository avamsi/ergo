package group

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestWriterSimple(t *testing.T) {
	var (
		b bytes.Buffer
		w = NewWriter(&b, 5)
	)
	for i := 0; i < 5; i++ {
		fmt.Fprintln(w.AddSection(), i)
	}
	if err := w.Close(); err != nil {
		t.Error(err)
	}
	var (
		got  = b.String()
		want = "0\n1\n2\n3\n4\n"
	)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestWriterConcurrent(t *testing.T) {
	var (
		g errgroup.Group
		b bytes.Buffer
		w = NewWriter(&b, 5)
	)
	for i := 0; i < 5; i++ {
		i := i // TODO: remove after Go 1.22.
		w := w.AddSection()
		g.Go(func() error {
			fmt.Fprintln(w, i)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		t.Error(err)
	}
	if err := w.Close(); err != nil {
		t.Error(err)
	}
	var (
		got  = b.String()
		want = "0\n1\n2\n3\n4\n"
	)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

type errWriter struct {
	err error
}

func (w errWriter) Write([]byte) (int, error) {
	return 0, w.err
}

func TestWriterError(t *testing.T) {
	var (
		want = errors.New("error")
		w    = NewWriter(errWriter{want}, 1)
	)
	fmt.Fprintln(w.AddSection(), "ok")
	if got := w.Close(); got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
