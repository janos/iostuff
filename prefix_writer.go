// Copyright (c) 2017, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iostuff

import (
	"bytes"
	"io"
)

// PrefixWriter appends a fixed prefix to each line that is
// written to encapsulated io.Writer.
type PrefixWriter struct {
	prefix     []byte
	w          io.Writer
	firstWrite bool
}

// NewPrefixWriter creates a new PrefixWriter.
func NewPrefixWriter(prefix string, w io.Writer) *PrefixWriter {
	return &PrefixWriter{
		prefix:     []byte(prefix),
		w:          w,
		firstWrite: true,
	}
}

func (w *PrefixWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	s := bytes.Split(p, []byte{'\n'})
	for i := 0; i < len(s)-1; i++ {
		s[i] = append(s[i], '\n')
	}
	p = bytes.Join(s, w.prefix)
	if w.firstWrite {
		p = append(w.prefix, p...)
		w.firstWrite = false
	}
	if _, err = w.w.Write(p); err != nil {
		return 0, err
	}
	return n, err
}
