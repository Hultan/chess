package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_SetPiece(t *testing.T) {
	b := NewBoard(false)

	b.SetPiece(pieceWhiteKing, 1)
	assert.Equal(t, pieceWhiteKing, b.Piece(1))
	assert.Equal(t, colorWhite, b.Color(1))

	b.SetPiece(pieceBlackKing, 9)
	assert.Equal(t, pieceBlackKing, b.Piece(9))
	assert.Equal(t, colorBlack, b.Color(9))

	b.SetPiece(pieceBlackRook, 16)
	assert.Equal(t, pieceBlackRook, b.Piece(16))
	assert.Equal(t, colorBlack, b.Color(16))

	b.SetPiece(pieceWhiteBishop, 63)
	assert.Equal(t, pieceWhiteBishop, b.Piece(63))
	assert.Equal(t, colorWhite, b.Color(63))
}

func TestBoard_Copy(t *testing.T) {
	b := NewBoard(false)
	b.SetPiece(pieceWhiteBishop, 34)
	b.SetPiece(pieceWhiteKing, 36)
	b.SetPiece(pieceBlackKing, 43)

	nb := b.Copy()

	assert.Equal(t, nb.board[0], b.board[0])
	assert.Equal(t, nb.board[1], b.board[1])
	assert.Equal(t, nb.board[2], b.board[2])
	assert.Equal(t, nb.board[3], b.board[3])
	assert.Equal(t, nb.extra, b.extra)
}

func TestBoard_MovePiece(t *testing.T) {
	b := NewBoard(false)
	b.SetPiece(pieceWhiteBishop, 34)
	b.SetPiece(pieceWhiteKing, 36)
	b.SetPiece(pieceBlackKing, 43)

	nb := b.MovePiece(36, 52)

	assert.NotEqual(t, nb.ToMove(), b.ToMove())
	assert.Equal(t, nb.board[0], b.board[0])
	assert.Equal(t, nb.board[1], b.board[1])
	assert.NotEqual(t, nb.board[2], b.board[2])
	assert.NotEqual(t, nb.board[3], b.board[3])
}

func TestNewBoard(t *testing.T) {
	b := NewBoard(true)

	assert.Equal(t, pieceBlackRook, b.Piece(56))
	assert.Equal(t, pieceBlackKnight, b.Piece(57))
	assert.Equal(t, pieceBlackBishop, b.Piece(58))
	assert.Equal(t, pieceBlackQueen, b.Piece(59))
	assert.Equal(t, pieceBlackKing, b.Piece(60))
	assert.Equal(t, pieceBlackBishop, b.Piece(61))
	assert.Equal(t, pieceBlackKnight, b.Piece(62))
	assert.Equal(t, pieceBlackRook, b.Piece(63))
}
