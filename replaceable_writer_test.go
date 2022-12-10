// Copyright (c) 2022, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iostuff_test

import (
	"io"
	"testing"

	"resenje.org/iostuff"
)

func TestReplaceableWriter(t *testing.T) {
	var currentFlag int

	w := iostuff.NewReplaceableWriter(func(flag int) (io.Writer, int, error) {
		flag++
		currentFlag = flag
		return io.Discard, flag, nil
	})
	defer w.Close()

	if currentFlag != 0 {
		t.Errorf("got current flag %v, want %v", currentFlag, 0)
	}

	for i := 1; i < 10; i++ {
		n, err := w.Write([]byte("test"))
		if err != nil {
			t.Fatal(err)
		}
		if n != 4 {
			t.Errorf("got %v bytes written, want %v", n, 4)
		}
		if currentFlag != i {
			t.Errorf("got current flag %v, want %v", currentFlag, i)
		}
	}
}

func BenchmarkReplaceableWriter_Write_discard(b *testing.B) {
	w := iostuff.NewReplaceableWriter(func(flag struct{}) (io.Writer, struct{}, error) {
		return io.Discard, struct{}{}, nil
	})
	defer w.Close()

	b.ResetTimer()

	var n int
	for i := 0; i < b.N; i++ {
		n, _ = w.Write(nil)
	}
	_ = n
}
