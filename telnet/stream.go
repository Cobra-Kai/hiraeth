package telnet

import (
	//	"io"
	"bufio"
	"fmt"
	"net"
)

type TelnetStream struct {
	conn      *net.TCPConn
	readError error
	// TODO: a negotiation handler
	// negotiator *TelnetNegotiator
	// TODO: table of WILL/WONT/DO/DONT negotiation data
	// statemap map[byte]telnetState

	reader *bufio.Reader
}

// TODO: accept parameters to control how WILL/WONT/DO/DONT is handled
func NewTelnetStream(conn *net.TCPConn) *TelnetStream {
	return &TelnetStream{conn, nil, bufio.NewReader(conn)}
}

func (stream *TelnetStream) Read(data []byte) (n int, err error) {
	if stream.readError != nil {
		return 0, stream.readError
	}

	// TODO: decode IAC, and expand escaped IAC into buffer
	/* This can be done as a loop that calls any pending IAC handlers, then
	 * returns available data up to the next IAC.
	 * Another option to implement this might be to use channels or to
	 * contruct an io.Reader interface that buffers more deeply
	 */

	count, err := stream.reader.Read(data)
	if err != nil {
		stream.readError = err
		// We'll assume the caller will log the error: log.Println(err)
	}

	return count, err

}

func (stream *TelnetStream) Write(data []byte) (n int, err error) {
	// TODO: decode and escape IAC
	return stream.conn.Write(data)
}

func (stream *TelnetStream) Close() error {
	return stream.conn.Close()
}

func (stream *TelnetStream) CloseWithError(err error) error {
	stream.readError = err
	return stream.conn.Close()
}

func (stream *TelnetStream) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(stream, a...)
}

func (stream *TelnetStream) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(stream, format, a...)
}

func (stream *TelnetStream) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(stream, a...)
}

// TODO: implement io.ByteReader
// TODO: implement io.ByteScanner
// TODO: implement io.RuneReader
// TODO: implement io.RuneScanner
// TODO: use bufio to support bufio.Reader.ReadBytes() and ReadLine()
