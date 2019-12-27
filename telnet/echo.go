package telnet

import (
	"io"
	"log"
)

type echoHandler struct {
}

func NewEchoHandler() echoHandler {
	return echoHandler{}
}

// implements Handler interface
func (h echoHandler) ServeTELNET(ts *TelnetStream) {
	buf := make([]byte, 32)

	ts.Printf("Hello Echo!\r\n")
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
