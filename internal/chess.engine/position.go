package chess_engine

import (
	"fmt"
	"strings"
)

type Position uint8

// Pos creates a new position from an integer
func Pos(i int) Position {
	if i < 0 || i > 63 {
		panic("invalid position in Pos()")
	}
	return Position(i)
}

// Sym creates a new position from an integer
// but in a symmetrical position to i
func Sym(i int) Position {
	if i < 0 || i > 63 {
		panic("invalid position in Pos()")
	}
	return Position(63 - i)
}

// Alg creates a new position based on a string (ie "b6" or "h3")
func Alg(alg string) Position {
	if len(alg) != 2 {
		panic("invalid alg")
	}
	a := strings.ToLower(alg)
	x := int(a[0]) - 97
	y := int(a[1]) - 49

	if x < 0 || x > 7 || y < 0 || y > 7 {
		panic("invalid alg")
	}

	return Pos(x + y*8)
}

// XY creates a new position based on coordinates (1-8, 1-8)
func XY(x, y int) Position {
	if x < 1 || x > 8 || y < 1 || y > 8 {
		panic("invalid xy")
	}

	return Pos((x - 1) + (y-1)*8)
}

// ToXY returns the coordinates (1-8, 1-8)
func (p Position) ToXY() (int, int) {
	return int(p%8 + 1), int(p/8 + 1)
}

// ToAlg returns the position as a string (i e "a1" or "g6")
func (p Position) ToAlg() string {
	return fmt.Sprintf("%c%v", byte(p%8+97), p/8+1)
}
