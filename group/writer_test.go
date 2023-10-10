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
		b   bytes.Buffer
		w   = NewWriter(&b, 5)
		err error
	)
	for i := 0; i < 5; i++ {
		_, e := fmt.Fprintln(w.Section(i), i)
		err = errors.Join(err, e)
	}
	if err = errors.Join(err, w.Close()); err != nil {
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
		g.Go(func() error {
			w := w.Section(i)
			defer w.Close()
			fmt.Fprintln(w, i)
			return nil
		})
	}
	if err := errors.Join(g.Wait(), w.Close()); err != nil {
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
		want   = errors.New("error")
		w      = NewWriter(errWriter{want}, 1)
		_, got = fmt.Fprintln(w.Section(0), "ok")
	)
	if got = errors.Join(got, w.Close()); got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}
