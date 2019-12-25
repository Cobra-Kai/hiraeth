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
package telnet

import (
	"log"

	"github.com/Cobra-Kai/hiraeth/shell"

	"github.com/reiver/go-telnet"
)

func Init(addr string) {
	handler := shell.NewShellHandler()

	log.Println("TELNET server listening on", addr)
	err := telnet.ListenAndServe(addr, handler)
	if nil != err {
		panic(err)
	}
}

func InitSsl(addr string) {
	handler := shell.NewShellHandler()

	err := telnet.ListenAndServeTLS(addr, "cert.pem", "key.pem", handler)
	if nil != err {
		panic(err)
	}
}
