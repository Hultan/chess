package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluator_ValueEmptyBoard(t *testing.T) {
	b := NewBoard(false)
	assert.Equal(t, 0, b.Value())
}

func TestEvaluator_ValueBoard(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, 0, b.Value())
}

func TestEvaluator_ValueSymmetry(t *testing.T) {
	b := NewBoard(false)

	// Make sure that a symmetrically placed pieces is
	// worth the same
	for piece := PieceWhitePawn; piece <= PieceWhiteKing; piece++ {
		for i := 0; i < 64; i++ {
			pos := Pos(i)
			sym := Sym(i)
			b.setPiece(piece, pos)
			value := b.Value()
			b.removePiece(pos)
			b.setPiece(piece+8, sym)
			assert.Equal(t, 0, b.Value()+value, "%s : %s failed", getPieceName(piece), pos.ToAlg())
			b.removePiece(sym)
		}
	}
}

func TestEvaluator_ValuePawn(t *testing.T) {
	b := NewBoard(false)
	b.setPiece(PieceWhitePawn, Alg("a3"))
	v := b.Value()
	assert.Equal(t, 105, v)
}

func TestEvaluator_ValueBishop(t *testing.T) {
	b := NewBoard(false)
	b.setPiece(PieceWhiteBishop, Alg("a3"))
	v := b.Value()
	assert.Equal(t, 320, v)
}

func TestEvaluator_ValueKnight(t *testing.T) {
	b := NewBoard(false)
	b.setPiece(PieceWhiteKnight, Alg("e5"))
	v := b.Value()
	assert.Equal(t, 340, v)
}

func TestEvaluator_ValueRook(t *testing.T) {
	b := NewBoard(false)
	b.setPiece(PieceWhiteRook, Alg("a5"))
	v := b.Value()
	assert.Equal(t, 495, v)
}

func TestEvaluator_ValueQueen(t *testing.T) {
	b := NewBoard(false)
	b.setPiece(PieceWhiteQueen, Alg("e5"))
	v := b.Value()
	assert.Equal(t, 905, v)
}

func TestEvaluator_ValueKing(t *testing.T) {
	b := NewBoard(false)
	b.setPiece(PieceWhiteKing, Alg("e5"))
	v := b.Value()
	assert.Equal(t, 20040, v)
}
