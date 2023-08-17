package answer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQ0001(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{n: 10},
			want: 23,
		},
		{
			name: "normal2",
			args: args{n: 1000},
			want: 233168,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, q0001(tt.args.n)); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestQ0002(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{n: 100},
			want: 44,
		},
		{
			name: "normal2",
			args: args{n: 4000000},
			want: 4613732,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, q0002(tt.args.n)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
