package cbutil

import "io"

type MultipleWriter struct {
	writers []io.Writer
}

func (m MultipleWriter) Write(p []byte) (n int, err error) {
	for _, writer := range m.writers {
		n, e := writer.Write(p)
		if e != nil {
			return n, e
		}
	}

	return n, nil
}

func NewMultipleWriter(writers ...io.Writer) MultipleWriter {
	return MultipleWriter{writers:writers}
}

