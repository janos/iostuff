// Copyright (c) 2017, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iostuff

import (
	"testing"
)

type sliceWriter struct {
	s []string
}

func (w *sliceWriter) Write(p []byte) (n int, err error) {
	w.s = append(w.s, string(p))
	return
}

func TestLineWriter(t *testing.T) {
	for _, c := range []struct {
		name string
		in   []string
		out  []string
	}{
		{
			name: "blank",
			in:   []string{""},
			out:  []string{""},
		},
		{
			name: "new line",
			in:   []string{"\n"},
			out:  []string{"\n"},
		},
		{
			name: "simple",
			in:   []string{"one"},
			out:  []string{"one"},
		},
		{
			name: "simple with trailing newline",
			in:   []string{"one\n"},
			out:  []string{"one\n"},
		},
		{
			name: "simple with leading newline",
			in:   []string{"\none"},
			out:  []string{"\n", "one"},
		},
		{
			name: "two lines",
			in:   []string{"one\ntest"},
			out:  []string{"one\n", "test"},
		},
		{
			name: "two lines with trailing newline",
			in:   []string{"one\ntest\n"},
			out:  []string{"one\n", "test\n"},
		},
		{
			name: "two lines with leading and trailing newline",
			in:   []string{"\none\ntest\n"},
			out:  []string{"\n", "one\n", "test\n"},
		},
		{
			name: "simple two writes",
			in:   []string{"on", "e"},
			out:  []string{"one"},
		},
		{
			name: "two words two writes split",
			in:   []string{"one\nte", "st"},
			out:  []string{"one\n", "test"},
		},
		{
			name: "simple three writes",
			in:   []string{"one", "te", "st"},
			out:  []string{"onetest"},
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			w := &sliceWriter{}
			lw := NewLineWriter(w)
			for _, s := range c.in {
				n, err := lw.Write([]byte(s))
				if err != nil {
					t.Fatal(err)
				}
				if n != len(s) {
					t.Fatalf("expected to write %v, written %v bytes", len(s), n)
				}
			}
			if err := lw.Close(); err != nil {
				t.Fatal(err)
			}
			for i := range c.out {
				if c.out[i] != w.s[i] {
					t.Errorf("expected line %v %q, got %q", i+1, c.out[i], w.s[i])
				}
			}
		})
	}
}
