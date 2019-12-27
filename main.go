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
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"

	"github.com/Cobra-Kai/hiraeth/database"
	"github.com/Cobra-Kai/hiraeth/mudserver"
	"github.com/Cobra-Kai/hiraeth/webserver"
)

var (
	telnetAddr = flag.String("telnet", "localhost:4000", "TELNET listen address")
	//	telnetSslAddr = flag.String("telnet-ssl", "localhost:4992", "TELNETS listen address")
	httpAddr   = flag.String("http", "localhost:8080", "HTTP listen address")
	dbFilename = flag.String("db", "mud.db", "database filename")
	maxProcs   = flag.Int("maxproc", 0, "MAXPROCS value")
)

func main() {
	flag.Parse()

	if maxProcs != nil && *maxProcs == 0 {
		runtime.GOMAXPROCS(*maxProcs)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	log.Println("Starting ...")

	go database.Init(*dbFilename)

	go webserver.Init(*httpAddr)
	go mudserver.Init(*telnetAddr)

	<-quit

	log.Println("Done!")
}
