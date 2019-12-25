package shell

import (
	"io"
)

// an interface same as in os/exec
type Handler interface {
	Run() error
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
}

type actionHandler struct {
	////	handler    Handler
	fn         HandlerFunc
	err        error
	stdin      io.ReadCloser
	stdout     io.WriteCloser
	stderr     io.WriteCloser
	stdinPipe  io.WriteCloser
	stdoutPipe io.ReadCloser
	stderrPipe io.ReadCloser
	args       []string
}

// convenient type to use as a handler function
type HandlerFunc func(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error

// turns a HandlerFunc into a Handler interface
func PromoteHandlerFunc(fn HandlerFunc, args ...string) actionHandler {
	// create pipe pairs
	stdin, stdinPipe := io.Pipe()
	stdoutPipe, stdout := io.Pipe()
	stderrPipe, stderr := io.Pipe()

	handler := actionHandler{
		fn:         fn,
		err:        nil,
		stdin:      stdin,
		stdout:     stdout,
		stderr:     stderr,
		stdinPipe:  stdinPipe,
		stdoutPipe: stdoutPipe,
		stderrPipe: stderrPipe,
	}

	return handler
}

func (handler *actionHandler) Run() error {
	if handler.err != nil {
		return handler.err
	}

	handler.err = handler.fn(handler.stdin, handler.stdout, handler.stderr, handler.args...)

	return handler.err
}

// for the Handler interface
func (handler *actionHandler) StdinPipe() (io.WriteCloser, error) {
	if handler.err != nil {
		return nil, handler.err
	}

	return handler.stdinPipe, nil
}

// for the Handler interface
func (handler *actionHandler) StdoutPipe() (io.ReadCloser, error) {
	if handler.err != nil {
		return nil, handler.err
	}

	return handler.stdoutPipe, nil
}

// for the Handler interface
func (handler *actionHandler) StderrPipe() (io.ReadCloser, error) {
	if handler.err != nil {
		return nil, handler.err
	}

	return handler.stderrPipe, nil
}
