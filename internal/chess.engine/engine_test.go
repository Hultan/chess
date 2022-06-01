package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmptyBoard(t *testing.T) {
	b := NewBoard(false)

	assert.NotNil(t, b)
	for i := 0; i < 63; i++ {
		assert.Equal(t, pieceNone, b.Piece(i))
	}
}

func TestNewBoard(t *testing.T) {
	b := NewBoard(true)

	assert.NotNil(t, b)
	assert.Equal(t, pieceBlackRook, b.Piece(alg("A8")))
	assert.Equal(t, pieceBlackKnight, b.Piece(alg("B8")))
	assert.Equal(t, pieceBlackBishop, b.Piece(alg("C8")))
	assert.Equal(t, pieceBlackQueen, b.Piece(alg("D8")))
	assert.Equal(t, pieceBlackKing, b.Piece(alg("E8")))
	assert.Equal(t, pieceBlackBishop, b.Piece(alg("F8")))
	assert.Equal(t, pieceBlackKnight, b.Piece(alg("G8")))
	assert.Equal(t, pieceBlackRook, b.Piece(alg("H8")))

	for i := 0; i < 8; i++ {
		assert.Equal(t, pieceBlackPawn, b.Piece(48+i))
		assert.Equal(t, pieceNone, b.Piece(40+i))
		assert.Equal(t, pieceNone, b.Piece(32+i))
		assert.Equal(t, pieceNone, b.Piece(24+i))
		assert.Equal(t, pieceNone, b.Piece(16+i))
		assert.Equal(t, pieceWhitePawn, b.Piece(8+i))
	}

	assert.Equal(t, pieceWhiteRook, b.Piece(alg("A1")))
	assert.Equal(t, pieceWhiteKnight, b.Piece(alg("B1")))
	assert.Equal(t, pieceWhiteBishop, b.Piece(alg("C1")))
	assert.Equal(t, pieceWhiteQueen, b.Piece(alg("D1")))
	assert.Equal(t, pieceWhiteKing, b.Piece(alg("E1")))
	assert.Equal(t, pieceWhiteBishop, b.Piece(alg("F1")))
	assert.Equal(t, pieceWhiteKnight, b.Piece(alg("G1")))
	assert.Equal(t, pieceWhiteRook, b.Piece(alg("H1")))
}

func TestBoard_Copy(t *testing.T) {
	b := NewBoard(false)
	b.SetPiece(pieceWhiteBishop, alg("B4"))
	b.SetPiece(pieceWhiteKing, alg("C5"))
	b.SetPiece(pieceBlackKing, alg("C8"))

	nb := b.Copy()

	assert.NotNil(t, nb)
	assert.Equal(t, nb.board[0], b.board[0])
	assert.Equal(t, nb.board[1], b.board[1])
	assert.Equal(t, nb.board[2], b.board[2])
	assert.Equal(t, nb.board[3], b.board[3])
	assert.Equal(t, nb.extra, b.extra)
}

func TestBoard_MovePiece(t *testing.T) {
	b := NewBoard(false)
	b.SetPiece(pieceWhiteBishop, alg("D6"))
	b.SetPiece(pieceWhiteKing, alg("D5"))
	b.SetPiece(pieceBlackKing, alg("E7"))

	nb := b.MovePiece(alg("E7"), alg("B2"))

	assert.NotEqual(t, nb.ToMove(), b.ToMove())
	assert.NotEqual(t, nb.board[0], b.board[0])
	assert.Equal(t, nb.board[1], b.board[1])
	assert.Equal(t, nb.board[2], b.board[2])
	assert.NotEqual(t, nb.board[3], b.board[3])
}

func TestBoard_SetPiece(t *testing.T) {
	b := NewBoard(false)

	b.SetPiece(pieceWhiteKing, alg("B1"))
	assert.Equal(t, pieceWhiteKing, b.Piece(alg("B1")))
	assert.Equal(t, colorWhite, b.Color(alg("B1")))

	b.SetPiece(pieceBlackKing, alg("B2"))
	assert.Equal(t, pieceBlackKing, b.Piece(alg("B2")))
	assert.Equal(t, colorBlack, b.Color(alg("B2")))

	b.SetPiece(pieceBlackRook, alg("D6"))
	assert.Equal(t, pieceBlackRook, b.Piece(alg("D6")))
	assert.Equal(t, colorBlack, b.Color(alg("D6")))

	b.SetPiece(pieceWhiteBishop, alg("H8"))
	assert.Equal(t, pieceWhiteBishop, b.Piece(alg("H8")))
	assert.Equal(t, colorWhite, b.Color(alg("H8")))
}

func TestBoard_ToMove(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, colorWhite, b.ToMove())

	nb := b.MovePiece(alg("A2"), alg("A4"))
	assert.Equal(t, colorBlack, nb.ToMove())

	b = nb.MovePiece(alg("H7"), alg("H5"))
	assert.Equal(t, colorWhite, b.ToMove())
}

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

func TestBoard_ToggleToMove(t *testing.T) {
	b := NewBoard(true)

	assert.Equal(t, colorWhite, b.ToMove())
	b.toggleToMove()
	assert.Equal(t, colorBlack, b.ToMove())
	b.toggleToMove()
	assert.Equal(t, colorWhite, b.ToMove())
}

func TestBoard_WhiteCastlingRights(t *testing.T) {
	b := NewBoard(true)

	assert.Equal(t, true, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(castlingBlackQueen))

	nb := b.MovePiece(7, 23)

	assert.Equal(t, false, nb.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, nb.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, nb.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, nb.CastlingRights(castlingBlackQueen))

	nb = nb.MovePiece(0, 16)

	assert.Equal(t, false, nb.CastlingRights(castlingWhiteKing))
	assert.Equal(t, false, nb.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, nb.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, nb.CastlingRights(castlingBlackQueen))

	nb.resetCastlingRights()
	nb = nb.MovePiece(4, 20)

	assert.Equal(t, false, nb.CastlingRights(castlingWhiteKing))
	assert.Equal(t, false, nb.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, nb.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, nb.CastlingRights(castlingBlackQueen))
}

func TestBoard_BlackCastlingRights(t *testing.T) {
	b := NewBoard(true)

	assert.Equal(t, true, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(castlingBlackQueen))

	nb := b.MovePiece(63, 55)

	assert.Equal(t, true, nb.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, nb.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, false, nb.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, nb.CastlingRights(castlingBlackQueen))

	nb = nb.MovePiece(56, 48)

	assert.Equal(t, true, nb.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, nb.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, false, nb.CastlingRights(castlingBlackKing))
	assert.Equal(t, false, nb.CastlingRights(castlingBlackQueen))

	nb.resetCastlingRights()
	nb = nb.MovePiece(60, 44)

	assert.Equal(t, true, nb.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, nb.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, false, nb.CastlingRights(castlingBlackKing))
	assert.Equal(t, false, nb.CastlingRights(castlingBlackQueen))
}
