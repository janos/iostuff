// Copyright (c) 2022, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iostuff

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type ReplaceableWriter[F comparable] struct {
	constructor func(flag F) (io.Writer, F, error)
	w           io.Writer
	flag        F
	mu          sync.Mutex
}

func NewReplaceableWriter[F comparable](constructor func(flag F) (io.Writer, F, error)) *ReplaceableWriter[F] {
	return &ReplaceableWriter[F]{
		constructor: constructor,
	}
}

func (w *ReplaceableWriter[F]) Write(b []byte) (int, error) {
	writer, err := w.writer()
	if err != nil {
		return 0, err
	}
	return writer.Write(b)
}

func (w *ReplaceableWriter[F]) writer() (io.Writer, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	nw, flag, err := w.constructor(w.flag)
	if err != nil {
		return nil, err
	}
	if nw != nil {
		if closer, ok := w.w.(io.Closer); ok {
			if err := closer.Close(); err != nil {
				if closer, ok := nw.(io.Closer); ok {
					_ = closer.Close()
				}
				return nil, fmt.Errorf("close previous writer: %w", err)
			}
		}
		w.w = nw
		w.flag = flag
	}
	if w.w == nil {
		return nil, errors.New("replaceable writer not constructed")
	}
	return w.w, nil
}

func (w *ReplaceableWriter[F]) Close() error {
	if closer, ok := w.w.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
