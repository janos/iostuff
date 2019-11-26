// Copyright (c) 2017, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iostuff

import (
	"bytes"
	"io"
)

// LineWriter ensures that only whole lines are sent to
// encapsulated io.Writer. It is useful to simplify
// writers that work with whole lines, like log lines
// or textual data processing.
type LineWriter struct {
	w    io.Writer
	buff []byte
}

// NewLineWriter wraps io.Writer and sends to its
// Write method only full lines. To flush the possible
// remaining bytes in buffer, make sure to call the Close
// method as soon as the writing is done.
func NewLineWriter(w io.Writer) *LineWriter {
	return &LineWriter{
		w:    w,
		buff: []byte{},
	}
}

// Write sends to the encapsulated Writer line by line
// and the remaining bytes keeps in buffer.
func (w *LineWriter) Write(p []byte) (n int, err error) {
	l := len(p)
	for {
		i := bytes.IndexByte(p, '\n')
		if i < 0 {
			w.buff = append(w.buff, p...)
			break
		}
		c, err := w.w.Write(append(w.buff, p[:i+1]...))
		if err != nil {
			return c, err
		}
		w.buff = w.buff[:0]
		p = p[i+1:]
	}
	return l, nil
}

// Flush writes any remaining data as a single line.
func (w *LineWriter) Flush() error {
	_, err := w.w.Write(w.buff)
	return err
}

// Close flushes the buffer and closes the encapsulated
// writer if it satisfies io.Closer interface.
func (w *LineWriter) Close() error {
	if err := w.Flush(); err != nil {
		return err
	}

	if c, ok := w.w.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
