package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
