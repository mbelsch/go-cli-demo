package exec

import "io"

type TeeWriter struct {
	writer1 io.Writer
	writer2 io.Writer
}

func NewTeeWriter(writer1 io.Writer, writer2 io.Writer) *TeeWriter {
	return &TeeWriter{
		writer1: writer1,
		writer2: writer2,
	}
}

func (w *TeeWriter) Write(p []byte) (n int, err error) {
	if w.writer1 != nil {
		n, err = w.writer1.Write(p)
		if err != nil {
			return n, err
		}
	}
	if w.writer2 != nil {
		n, err = w.writer2.Write(p)
	}

	return n, err
}