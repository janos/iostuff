// Copyright (c) 2022, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iostuff

import "io"

type NopWriteCloser struct {
	io.Writer
}

func (NopWriteCloser) Close() error {
	return nil
}

func NewNopWriteCloser(w io.Writer) io.WriteCloser {
	return &NopWriteCloser{
		Writer: w,
	}
}
