// Copyright (c) 2019 Jon Mayo <jon@rm-f.net>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
package main

import (
	"log"

	"github.com/nsf/termbox-go"
)

func mvaddstr(x, y int, s string) {
	fg := termbox.ColorYellow
	bg := termbox.ColorBlue
	for _, r := range s {
		termbox.SetCell(x, y, r, fg, bg)
		// TODO: get runewidth
		x += 1
	}
}

func center(y int, s string) {
	w, _ := termbox.Size()
	x := 0

	// TODO: use runewidth to determine display length
	l := len(s)
	if l < w {
		x = (w - l) / 2
	}
	mvaddstr(x, y, s)
}

func paint() {
	termbox.Clear(termbox.ColorYellow, termbox.ColorBlue)
	center(0, "Press ESC to exit!")

	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	paint()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				return
			}
		case termbox.EventResize:
			paint()
		}
	}
}
