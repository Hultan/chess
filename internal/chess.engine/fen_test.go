package chess_engine

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_ToFEN(t *testing.T) {
	b := NewBoard(false)
	startingFen := "8/6P1/pn1RPKR1/8/p1k2b2/5B1p/2N3bp/3q4 w - - 0 1"
	b.FromFEN(startingFen)
	p := b.print()
	assert.Equal(t, "        \n      P \npn RPKR \n        \np k  b  \n     B p\n  N   bp\n   q    \nWhite to move\nHalf move count : 0\nMove count : 1\n", p)
	newFen := b.ToFEN()
	assert.Equal(t, startingFen, newFen)
}

func TestBoard_ToFEN2(t *testing.T) {
	b := NewBoard(false)
	startingFen := "3r2Nr/1pp3p1/1P1Ppq2/p3p2B/1n1QP2p/pN2PRPK/PPknP2B/4b2b w - - 0 1"
	b.FromFEN(startingFen)
	p := b.print()
	assert.Equal(t, "   r  Nr\n pp   p \n P Ppq  \np   p  B\n n QP  p\npN  PRPK\nPPknP  B\n    b  b\nWhite to move\nHalf move count : 0\nMove count : 1\n", p)
	newFen := b.ToFEN()
	assert.Equal(t, startingFen, newFen)
}

func TestBoard_ToFEN3(t *testing.T) {
	b := NewBoard(false)
	startingFen := "r3k2r/pp1n2pp/2p2q2/b2p1n2/BP1Pp3/P1N2P2/2PB2PP/R2Q1RK1 w kq b3 0 13"
	b.FromFEN(startingFen)
	p := b.print()
	fmt.Println(p)
	assert.Equal(t, "r   k  r\npp n  pp\n  p  q  \nb  p n  \nBP Pp   \nP N  P  \n  PB  PP\nR  Q RK \nWhite to move\nCastling : kq\nEn passant : b3\nHalf move count : 0\nMove count : 13\n", p)
	newFen := b.ToFEN()
	assert.Equal(t, startingFen, newFen)
}
