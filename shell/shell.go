package shell

import (
	"fmt"
	"io"
	"strings"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

type ShellHandler struct {
	Prompt string
}

func (handler ShellHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
	/// wish I could do something like this ...
	/// for {
	///	line, err := r.ReadString('\n')
	///	if err != nil {
	///		break
	///	}
	///	oi.LongWrite(w, []byte(line))
	/// }

	// TODO: track connection state

	var buffer [1]byte
	p := buffer[:]
	var lineBuf strings.Builder

	for {
		n, err := r.Read(p)
		if err != nil {
			// TODO: post error somewhere
			break
		} else if n <= 0 {
			// I'm not sure why this happens
			continue
		}

		if p[0] == '\n' {
			line := lineBuf.String()
			lineBuf.Reset()

			// TODO: there needs to be a way to preserve the line spacing for messages, emotes, etc.
			args := strings.Fields(line)
			commandName := args[0]

			// TODO: look up the command, else ...

			act, ok := getCommandAction(commandName)

			if !ok {
				var response strings.Builder
				fmt.Fprintf(&response, "I don't know `%s'\r\n", args[0])

				_, err := oi.LongWrite(w, []byte(response.String()))
				if err != nil {
					// TODO: post error somewhere
					break
				}
			} else {
				// TODO: link up inputs and outputs
				// TODO... redirect(ctx, r, w)
				stdinPipe, err := act.StdinPipe()
				if err != nil {
					break
				}
				stdoutPipe, err := act.StdoutPipe()
				if err != nil {
					break
				}
				stderrPipe, err := act.StderrPipe()
				if err != nil {
					break
				}
				connectPipe(ctx, stdinPipe, r)
				connectPipe(ctx, w, stdoutPipe)
				connectPipe(ctx, w, stderrPipe)
				act.Run()
			}
		} else {
			lineBuf.WriteByte(p[0])
		}
	}
}

func NewShellHandler() *ShellHandler {

	handler := ShellHandler{}

	return &handler
}

func connectPipe(ctx telnet.Context, writer io.Writer, reader io.Reader) {
	if writer == nil || reader == nil {
		return
	}

	go func() {
		var buffer [1]byte
		p := buffer[:]

		// TODO: read in multiple bytes, not just 1 byte

		for {
			n, err := reader.Read(p)
			if err != nil {
				// TODO: post error somewhere
				break
			} else if n <= 0 {
				// I'm not sure why this happens
				continue
			}

			oi.LongWrite(writer, p)
		}
	}()
}
