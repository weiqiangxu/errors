package errors

import (
	"fmt"
	"io"
)

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }

func (w *withMessage) Cause() error { return w.cause }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *withMessage) Unwrap() error { return w.cause }

// Format cover fmt Formatter
func (w *withMessage) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, err := fmt.Fprintf(s, "%+v\n", w.Cause())
			if err != nil {
				return
			}
			_, err = io.WriteString(s, w.msg)
			if err != nil {
				return
			}
			return
		}
		fallthrough
	case 's', 'q':
		_, err := io.WriteString(s, w.Error())
		if err != nil {
			return
		}
	}
}
