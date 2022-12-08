package errors

import (
	stdError "errors"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test for new stack",
			args: args{
				message: "catch error",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := New(tt.args.message)
			t.Logf("%+v", err)
			t.Logf("%s", err)
		})
	}
}

func TestErrorf(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test string",
			args: args{
				format: "error=%s",
				args:   []interface{}{New("hello")},
			},
			wantErr: false,
		},
		{
			name: "test stack",
			args: args{
				format: "error=%+v",
				args:   []interface{}{New("hello")},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Errorf(tt.args.format, tt.args.args...)
			t.Logf("%s", err)
		})
	}
}

func TestWithStack(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "with stack",
			args: args{
				err: New("i am jack"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("no with stack=%+v", tt.args.err)
			err := WithStack(tt.args.err)
			t.Logf("with stack=%+v", err)
			t.Logf("%s", err)
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		err     error
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "wrap",
			args: args{
				err:     New("i am jack"),
				message: "hello",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.args.err, tt.args.message)
			t.Logf("%s", err)
		})
	}
}

func TestWrapf(t *testing.T) {
	type args struct {
		err    error
		format string
		args   []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test wrap and format",
			args: args{
				err: New("jack"), format: "%s min=%d", args: []interface{}{"hello", 20},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrapf(tt.args.err, tt.args.format, tt.args.args...)
			t.Logf("%s", err)
			// hello min=20: jack
		})
	}
}

func TestWithMessage(t *testing.T) {
	type args struct {
		err     error
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WithMessage(tt.args.err, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("WithMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWithMessagef(t *testing.T) {
	type args struct {
		err    error
		format string
		args   []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test message and format",
			args: args{
				err:    New("jack"),
				format: "min=%d name=%s",
				args: []interface{}{
					18,
					"m",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WithMessagef(tt.args.err, tt.args.format, tt.args.args...)
			t.Logf("%+v", err)
			t.Logf("%s", err)
		})
	}
}

func TestCause(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test cause",
			args: args{
				err: Wrap(New("hello"), "i am jack"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Cause(tt.args.err)
			t.Logf("%s", err)
		})
	}
}

func TestIs(t *testing.T) {
	type args struct {
		err    error
		target error
	}
	err := New("hello")
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "错误链中的任何错误是否与目标匹配",
			args: args{
				err:    Wrap(err, "i am jack"),
				target: err,
			},
			want: false,
		},
		{
			name: "错误链中的任何错误是否与目标匹配",
			args: args{
				err:    Cause(Wrap(err, "i am jack")),
				target: err,
			},
			want: false,
		},
		{
			name: "错误链中的任何错误是否与目标匹配",
			args: args{
				err:    New("hello"),
				target: err,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Is(tt.args.err, tt.args.target)
			t.Logf("%+v", got)
		})
	}
}

func TestAs(t *testing.T) {
	type args struct {
		err    error
		target interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "查找错误链中与目标匹配的第一个错误，如果是，则设置目标设置为该错误值并返回true",
			args: args{
				err:    Wrap(New("hello"), "i am jack"),
				target: New("hello"),
			},
			want: false,
		},
		{
			name: "查找错误链中与目标匹配的第一个错误，如果是，则设置目标设置为该错误值并返回true",
			args: args{
				err:    Cause(Wrap(New("hello"), "i am jack")),
				target: New("hello"),
			},
			want: false,
		},
		{
			name: "查找错误链中与目标匹配的第一个错误，如果是，则设置目标设置为该错误值并返回true",
			args: args{
				err:    New("jack"),
				target: New("hello"),
			},
			want: false,
		},
		{
			name: "查找错误链中与目标匹配的第一个错误，如果是，则设置目标设置为该错误值并返回true",
			args: args{
				err:    stdError.New("jack"),
				target: New("hello"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := As(tt.args.err, tt.args.target)
			t.Logf("%+v", e)
		})
	}
}

func TestUnwrap(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "unwrap会擦除一层stack",
			args: args{
				err: WithStack(New("jack")),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%+v", tt.args.err)
			err := Unwrap(tt.args.err)
			t.Logf("%+v", err)
		})
	}
}
