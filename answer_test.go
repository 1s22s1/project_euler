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

func TestQ0003(t *testing.T) {
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
			args: args{n: 13195},
			want: 29,
		},
		{
			name: "normal2",
			args: args{n: 600851475143},
			want: 6857,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, q0003(tt.args.n)); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		n string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal1",
			args: args{n: "13431"},
			want: true,
		},
		{
			name: "normal2",
			args: args{n: "13432"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, isPalindrome(tt.args.n)); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestQ0004(t *testing.T) {
	type args struct {
		first int
		last  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{first: 10, last: 99},
			want: 9009,
		},
		{
			name: "normal2",
			args: args{first: 100, last: 999},
			want: 906609,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, q0004(tt.args.first, tt.args.last)); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestGcd(t *testing.T) {
	type args struct {
		n int
		m int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{n: 36, m: 24},
			want: 12,
		},
		{
			name: "normal2",
			args: args{n: 18, m: 30},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, gcd(tt.args.n, tt.args.m)); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestLcm(t *testing.T) {
	type args struct {
		n int
		m int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{n: 6, m: 8},
			want: 24,
		},
		{
			name: "normal2",
			args: args{n: 32, m: 24},
			want: 96,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, lcm(tt.args.n, tt.args.m)); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestQ0005(t *testing.T) {
	type args struct {
		first int
		last  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{first: 1, last: 10},
			want: 2520,
		},
		{
			name: "normal2",
			args: args{first: 1, last: 20},
			want: 232792560,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, q0005(tt.args.first, tt.args.last)); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestQ0006(t *testing.T) {
	type args struct {
		first int
		last  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{first: 1, last: 10},
			want: 2640,
		},
		{
			name: "normal2",
			args: args{first: 1, last: 100},
			want: 25164150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, q0006(tt.args.first, tt.args.last)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
