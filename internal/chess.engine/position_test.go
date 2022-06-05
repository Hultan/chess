package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPos(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{"Pos(0)", args{0}, 0},
		{"Pos(63)", args{63}, 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Pos(tt.args.i), "Pos(%v)", tt.args.i)
		})
	}
}

func TestPos_Invalid(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Pos(-1)", args{-1}},
		{"Pos(64)", args{64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() { Pos(tt.args.i) }, "Pos(%v)", tt.args.i)
		})
	}
}

func TestAlg(t *testing.T) {
	type args struct {
		alg string
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{"a1", args{"a1"}, 0},
		{"a8", args{"a8"}, 56},
		{"h1", args{"h1"}, 7},
		{"h8", args{"h8"}, 63},
		{"e4", args{"e4"}, 28},
		{"c7", args{"c7"}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Alg(tt.args.alg), "Alg(%v)", tt.args.alg)
		})
	}
}

func TestAlg_Invalid(t *testing.T) {
	type args struct {
		alg string
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{""}},
		{"a9", args{"a9"}},
		{"x8", args{"x8"}},
		{"h11", args{"h11"}},
		{"xh8!", args{"xh8!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() { Alg(tt.args.alg) }, "Alg(%v)", tt.args.alg)
		})
	}
}

func TestXY(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want Position
	}{
		{"Arg(1,1)", args{1, 1}, 0},
		{"Arg(1,8)", args{1, 8}, 56},
		{"Arg(8,1)", args{8, 1}, 7},
		{"Arg(8,8)", args{8, 8}, 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, XY(tt.args.x, tt.args.y), "XY(%v, %v)", tt.args.x, tt.args.y)
		})
	}
}

func TestXY_Invalid(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Arg(0,1)", args{0, 1}},
		{"Arg(1,0)", args{1, 0}},
		{"Arg(1,9)", args{1, 9}},
		{"Arg(9,1)", args{9, 1}},
		{"Arg(77,-12)", args{77, -12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Panics(t, func() { XY(tt.args.x, tt.args.y) }, "XY(%v, %v)", tt.args.x, tt.args.y)
		})
	}
}

func TestPosition_ToAlg(t *testing.T) {
	tests := []struct {
		name string
		p    Position
		want string
	}{
		{"Pos(0)", Pos(0), "a1"},
		{"Pos(56)", Pos(56), "a8"},
		{"Pos(7)", Pos(7), "h1"},
		{"Pos(63)", Pos(63), "h8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.p.ToAlg(), "ToAlg()")
		})
	}
}

func TestPosition_ToXY(t *testing.T) {
	tests := []struct {
		name  string
		p     Position
		want  int
		want1 int
	}{
		{"Pos(0)", 0, 1, 1},
		{"Pos(7)", 7, 8, 1},
		{"Pos(56)", 56, 1, 8},
		{"Pos(63)", 63, 8, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.p.ToXY()
			assert.Equalf(t, tt.want, got, "ToXY()")
			assert.Equalf(t, tt.want1, got1, "ToXY()")
		})
	}
}
