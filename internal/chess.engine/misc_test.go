package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_Alg(t *testing.T) {
	type args struct {
		alg string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"A1", args{"A1"}, 0},
		{"A3", args{"A3"}, 16},
		{"F4", args{"F4"}, 29},
		{"H8", args{"H8"}, 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, alg(tt.args.alg), "alg(%v)", tt.args.alg)
		})
	}
}

func Test_xy(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1,1", args{1, 1}, 0},
		{"8,8", args{8, 8}, 63},
		{"5,4", args{5, 4}, 28},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, xy(tt.args.x, tt.args.y), "xy(%v, %v)", tt.args.x, tt.args.y)
		})
	}
}
