package telnet

import (
	"log"
)

type fnType func() (int, fnType, error)

// TODO: use something like this to transfer control between processing functions. function may return a nil function to close the connection
func Loop(fn fnType) error {
	result, f, err := fn()
	for {
		if f == nil {
			break
		}
		result, f, err = f()
	}

	log.Println(result)

	return err
}
