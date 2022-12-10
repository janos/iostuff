// Copyright (c) 2017, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iostuff_test

import (
	"bytes"
	"strings"
	"testing"

	"resenje.org/iostuff"
)

func TestPrefixWriter(t *testing.T) {
	for _, c := range []struct {
		name   string
		prefix string
		in     string
		out    string
	}{
		{
			name:   "blank",
			prefix: "",
			in:     "",
			out:    "",
		},
		{
			name:   "empty input",
			prefix: "prefix",
			in:     "",
			out:    "",
		},
		{
			name:   "single char input",
			prefix: "prefix",
			in:     "1",
			out:    "prefix1",
		},
		{
			name:   "single char in second line input",
			prefix: "prefix",
			in:     "\n1",
			out:    "prefix\nprefix1",
		},
		{
			name:   "multiline with single char input",
			prefix: "> ",
			in: `1
2
3
4
5`,
			out: `> 1
> 2
> 3
> 4
> 5`,
		},
		{
			name:   "multiline input",
			prefix: "> ",
			in: `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Etiam imperdiet orci nec sem scelerisque, vehicula laoreet neque bibendum.
Aliquam elementum dui eu egestas fermentum.
Nam elementum nisi in lacinia scelerisque.
Nullam nec eros gravida dui blandit consectetur vel et tortor.
Maecenas sit amet nunc elementum dui lobortis ornare ut a ante.
Phasellus rhoncus tortor nec metus ullamcorper, vel maximus nisl posuere.`,
			out: `> Lorem ipsum dolor sit amet, consectetur adipiscing elit.
> Etiam imperdiet orci nec sem scelerisque, vehicula laoreet neque bibendum.
> Aliquam elementum dui eu egestas fermentum.
> Nam elementum nisi in lacinia scelerisque.
> Nullam nec eros gravida dui blandit consectetur vel et tortor.
> Maecenas sit amet nunc elementum dui lobortis ornare ut a ante.
> Phasellus rhoncus tortor nec metus ullamcorper, vel maximus nisl posuere.`,
		},
		{
			name:   "multiline input with blank last line",
			prefix: "> ",
			in: `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Etiam imperdiet orci nec sem scelerisque, vehicula laoreet neque bibendum.
Aliquam elementum dui eu egestas fermentum.
Nam elementum nisi in lacinia scelerisque.
Nullam nec eros gravida dui blandit consectetur vel et tortor.
Maecenas sit amet nunc elementum dui lobortis ornare ut a ante.
Phasellus rhoncus tortor nec metus ullamcorper, vel maximus nisl posuere.
`,
			out: `> Lorem ipsum dolor sit amet, consectetur adipiscing elit.
> Etiam imperdiet orci nec sem scelerisque, vehicula laoreet neque bibendum.
> Aliquam elementum dui eu egestas fermentum.
> Nam elementum nisi in lacinia scelerisque.
> Nullam nec eros gravida dui blandit consectetur vel et tortor.
> Maecenas sit amet nunc elementum dui lobortis ornare ut a ante.
> Phasellus rhoncus tortor nec metus ullamcorper, vel maximus nisl posuere.
> `,
		},
		{
			name:   "multiline prefix with multiline with single char input",
			prefix: ">\n> ",
			in: `1
2
3
4
5`,
			out: `>
> 1
>
> 2
>
> 3
>
> 4
>
> 5`,
		},
		{
			name:   "multiline prefix with multiline input",
			prefix: ">\n> ",
			in: `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Etiam imperdiet orci nec sem scelerisque, vehicula laoreet neque bibendum.
Aliquam elementum dui eu egestas fermentum.
Nam elementum nisi in lacinia scelerisque.
Nullam nec eros gravida dui blandit consectetur vel et tortor.
Maecenas sit amet nunc elementum dui lobortis ornare ut a ante.
Phasellus rhoncus tortor nec metus ullamcorper, vel maximus nisl posuere.`,
			out: `>
> Lorem ipsum dolor sit amet, consectetur adipiscing elit.
>
> Etiam imperdiet orci nec sem scelerisque, vehicula laoreet neque bibendum.
>
> Aliquam elementum dui eu egestas fermentum.
>
> Nam elementum nisi in lacinia scelerisque.
>
> Nullam nec eros gravida dui blandit consectetur vel et tortor.
>
> Maecenas sit amet nunc elementum dui lobortis ornare ut a ante.
>
> Phasellus rhoncus tortor nec metus ullamcorper, vel maximus nisl posuere.`,
		},
		{
			name:   "multiline prefix with multiline input with blank last line",
			prefix: ">\n> ",
			in: `Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Etiam imperdiet orci nec sem scelerisque, vehicula laoreet neque bibendum.
Aliquam elementum dui eu egestas fermentum.
Nam elementum nisi in lacinia scelerisque.
Nullam nec eros gravida dui blandit consectetur vel et tortor.
Maecenas sit amet nunc elementum dui lobortis ornare ut a ante.
Phasellus rhoncus tortor nec metus ullamcorper, vel maximus nisl posuere.
`,
			out: `>
> Lorem ipsum dolor sit amet, consectetur adipiscing elit.
>
> Etiam imperdiet orci nec sem scelerisque, vehicula laoreet neque bibendum.
>
> Aliquam elementum dui eu egestas fermentum.
>
> Nam elementum nisi in lacinia scelerisque.
>
> Nullam nec eros gravida dui blandit consectetur vel et tortor.
>
> Maecenas sit amet nunc elementum dui lobortis ornare ut a ante.
>
> Phasellus rhoncus tortor nec metus ullamcorper, vel maximus nisl posuere.
>
> `,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			r := strings.NewReader(c.in)
			buf := &bytes.Buffer{}
			if _, err := r.WriteTo(iostuff.NewPrefixWriter(c.prefix, buf)); err != nil {
				t.Fatal(err)
			}
			if buf.String() != c.out {
				t.Errorf("expected %q, got %q", c.out, buf.String())
			}
		})
	}
}
