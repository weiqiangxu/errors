package errors

import (
	"fmt"
	"io"
)

// StackTrace is stack of Frames from innermost (newest) to outermost (oldest).
type StackTrace []Frame

// Format formats the stack of Frames according to the fmt.Formatter interface.
//
//	%s	lists source files for each Frame in the stack
//	%v	lists the source file and line number for each Frame in the stack
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//	%+v   Prints filename, function, and line number for each Frame in the stack.
func (st StackTrace) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			for _, f := range st {
				_, err := io.WriteString(s, "\n")
				if err != nil {
					return
				}
				f.Format(s, verb)
			}
		case s.Flag('#'):
			_, err := fmt.Fprintf(s, "%#v", []Frame(st))
			if err != nil {
				return
			}
		default:
			st.formatSlice(s, verb)
		}
	case 's':
		st.formatSlice(s, verb)
	}
}

// formatSlice will format this StackTrace into the given buffer as a slice of
// Frame, only valid when called with '%s' or '%v'.
func (st StackTrace) formatSlice(s fmt.State, verb rune) {
	_, err := io.WriteString(s, "[")
	if err != nil {
		return
	}
	for i, f := range st {
		if i > 0 {
			_, err := io.WriteString(s, " ")
			if err != nil {
				return
			}
		}
		f.Format(s, verb)
	}
	_, err = io.WriteString(s, "]")
	if err != nil {
		return
	}
}
