package errors

import (
	"fmt"
	"io"
	"path"
	"runtime"
	"strconv"
)

// Frame represents a program counter inside a stack frame
type Frame uintptr

// pc returns the program counter for this frame;
// multiple frames may have the same PC value.
func (f Frame) pc() uintptr { return uintptr(f) - 1 }

// file returns the full path to the file that contains the
// function for this Frame's pc.
func (f Frame) file() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	file, _ := fn.FileLine(f.pc())
	return file
}

// line returns the line number of source code of the
// function for this Frame's pc.
func (f Frame) line() int {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return 0
	}
	_, line := fn.FileLine(f.pc())
	return line
}

// name returns the name of this function, if known.
func (f Frame) name() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

// Format formats the frame according to the fmt.Formatter interface.
//
//	%s    source file
//	%d    source line
//	%n    function name
//	%v    equivalent to %s:%d
//
// Format accepts flags that alter the printing of some verbs, as follows:
//
//	%+s   function name and path of source file relative to the compile time
//	      GOPATH separated by \n\t (<funBaseName>\n\t<path>)
//	%+v   equivalent to %+s:%d
func (f Frame) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		switch {
		case s.Flag('+'):
			_, err := io.WriteString(s, f.name())
			if err != nil {
				return
			}
			_, err = io.WriteString(s, "\n\t")
			if err != nil {
				return
			}
			_, err = io.WriteString(s, f.file())
			if err != nil {
				return
			}
		default:
			_, err := io.WriteString(s, path.Base(f.file()))
			if err != nil {
				return
			}
		}
	case 'd':
		_, err := io.WriteString(s, strconv.Itoa(f.line()))
		if err != nil {
			return
		}
	case 'n':
		_, err := io.WriteString(s, funBaseName(f.name()))
		if err != nil {
			return
		}
	case 'v':
		f.Format(s, 's')
		_, err := io.WriteString(s, ":")
		if err != nil {
			return
		}
		f.Format(s, 'd')
	}
}

// MarshalText formats a stacktrace Frame as a text string. The output is the
// same as that of fmt.Sprintf("%+v", f), but without newlines or tabs.
func (f Frame) MarshalText() ([]byte, error) {
	name := f.name()
	if name == "unknown" {
		return []byte(name), nil
	}
	return []byte(fmt.Sprintf("%s %s:%d", name, f.file(), f.line())), nil
}
