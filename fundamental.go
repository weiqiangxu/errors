package errors

import (
	"fmt"
	"io"
)

// fundamental is an error that has a message and a stack, but no caller.
type fundamental struct {
	msg string
	*stack
}

func (f *fundamental) Error() string { return f.msg }

func (f *fundamental) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			_, err := io.WriteString(s, f.msg)
			if err != nil {
				return
			}
			f.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		_, err := io.WriteString(s, f.msg)
		if err != nil {
			return
		}
	case 'q':
		_, err := fmt.Fprintf(s, "%q", f.msg)
		if err != nil {
			return
		}
	}
}
