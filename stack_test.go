package errors

import (
	"fmt"
	"testing"
)

func Test_stack_StackTrace(t *testing.T) {
	s := callers()
	tests := []struct {
		name string
		s    stack
		want StackTrace
	}{
		{
			name: "test stack trace",
			s:    *s,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.StackTrace()
			for k, v := range got {
				t.Logf("k=%d and v=%s and `%+v` and `%#v`", k, v, v, v)
			}
		})
	}
}

func Test_stack_Format(t *testing.T) {
	s := callers()
	p := &pp{}
	type args struct {
		st   fmt.State
		verb rune
	}
	tests := []struct {
		name string
		s    stack
		args args
	}{
		{
			name: "test format",
			s:    *s,
			args: args{
				st:   p,
				verb: 'v',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Format(tt.args.st, tt.args.verb)
			if o, ok := tt.args.st.(interface {
				String() string
			}); ok {
				t.Logf("%s", o.String())
				return
			}
		})
	}
}
