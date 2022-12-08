package errors

import "testing"

func Test_funBaseName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test func base name 1",
			args: args{
				name: "/a/b/c",
			},
			want: "",
		},
		{
			name: "test func base name 2",
			args: args{
				name: "c\\d/e",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := funBaseName(tt.args.name)
			t.Logf("%s", got)
		})
	}
}
