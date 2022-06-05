package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluator_ValueEmptyBoard(t *testing.T) {
	b := NewBoard(false)
	e := NewEvaluator(b)
	assert.Equal(t, 0, e.Value())
}

func TestEvaluator_ValueBoard(t *testing.T) {
	b := NewBoard(true)
	e := NewEvaluator(b)
	assert.Equal(t, 0, e.Value())
}

func TestEvaluator_ValuePawn(t *testing.T) {
	b := NewBoard(false)
	b.setPiece(PieceWhitePawn, 10)
	e := NewEvaluator(b)
	assert.Equal(t, 110, e.Value())
}

func TestEvaluator_ValueRook(t *testing.T) {
	b := NewBoard(false)
	b.setPiece(PieceBlackRook, 32)
	e := NewEvaluator(b)
	assert.Equal(t, -495, e.Value())
}
