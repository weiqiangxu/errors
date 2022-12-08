package errors

// pp create for implement fmt.State using for test formatter
type pp struct {
	buf []byte
}

// Write is the function to call to emit formatted output to be printed.
func (p *pp) Write(b []byte) (n int, err error) {
	p.buf = append(p.buf, b...)
	return 0, nil
}

// Width returns the value of the width option and whether it has been set.
func (p *pp) Width() (wid int, ok bool) {
	return 0, false
}

// Precision returns the value of the precision option and whether it has been set.
func (p *pp) Precision() (prec int, ok bool) {
	return 0, false
}

// Flag reports whether the flag c, a character, has been set.
func (p *pp) Flag(c int) bool {
	switch c {
	case '-':
		return true
	case '+':
		return true
	case '#':
		return true
	case ' ':
		return true
	case '0':
		return true
	}
	return true
}

func (p *pp) String() string {
	return string(p.buf)
}
