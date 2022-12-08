package errors

import (
	"fmt"
	"io"
)

type withStack struct {
	error
	*stack
}

func (w *withStack) Cause() error { return w.error }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withStack) Unwrap() error { return w.error }

func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, err := fmt.Fprintf(s, "%+v", w.Cause())
			if err != nil {
				return
			}
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		_, err := io.WriteString(s, w.Error())
		if err != nil {
			return
		}
	case 'q':
		_, err := fmt.Fprintf(s, "%q", w.Error())
		if err != nil {
			return
		}
	}
}
