package mudserver

import (
	"io"
	"log"

	"github.com/Cobra-Kai/hiraeth/telnet"
)

type mudHandler struct {
}

func Init(addr string) error {

	// TODO: initialize some useful handler

	handler := NewMudHandler()
	return telnet.ListenAndServe(addr, handler)
}

func NewMudHandler() mudHandler {
	return mudHandler{}
}

// implements Handler interface
func (h mudHandler) ServeTELNET(ts *telnet.TelnetStream) {
	buf := make([]byte, 32)

	ts.Printf("Hello Mud!\r\n")

	// TODO: implement something more useful
	for {
		_, err := ts.Read(buf)
		if err == io.EOF {
			ts.Close()
			break
		} else if err != nil {
			ts.CloseWithError(err)
			log.Println(err)
			break
		}
		ts.Write(buf)
	}
}
