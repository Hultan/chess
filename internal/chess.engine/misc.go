package chess_engine

import (
	"strings"
)

func alg(alg string) int {
	if len(alg) != 2 {
		panic("invalid alg")
	}
	a := strings.ToLower(alg)
	x := int(a[0]) - 97
	y := int(a[1]) - 49

	if x < 0 || x > 7 || y < 0 || y > 7 {
		panic("invalid alg")
	}

	return x + y*8
}

func xy(x, y int) int {
	if x < 1 || x > 8 || y < 1 || y > 8 {
		panic("invalid xy")
	}

	return (x - 1) + (y-1)*8
}
