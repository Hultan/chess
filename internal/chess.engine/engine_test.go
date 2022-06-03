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

	// Extra
	assert.Equal(t, colorWhite, b.ToMove())
	assert.Equal(t, true, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(castlingBlackQueen))
	assert.Equal(t, 1, b.MoveCount())
	assert.Equal(t, 0, b.HalfMoveCount())
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

	// Extra
	assert.Equal(t, colorWhite, b.ToMove())
	assert.Equal(t, true, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(castlingBlackQueen))
	assert.Equal(t, 1, b.MoveCount())
}

func TestBoard_Copy(t *testing.T) {
	b := NewBoard(false)
	b.setPiece(pieceWhiteBishop, alg("B4"))
	b.setPiece(pieceWhiteKing, alg("C5"))
	b.setPiece(pieceBlackKing, alg("C8"))

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
	b.setPiece(pieceWhiteBishop, alg("D6"))
	b.setPiece(pieceWhiteKing, alg("D5"))
	b.setPiece(pieceBlackKing, alg("E7"))

	nb := b.MovePiece(alg("E7"), alg("B2"))

	assert.NotEqual(t, nb.ToMove(), b.ToMove())
	assert.NotEqual(t, nb.board[0], b.board[0])
	assert.Equal(t, nb.board[1], b.board[1])
	assert.Equal(t, nb.board[2], b.board[2])
	assert.NotEqual(t, nb.board[3], b.board[3])
}

func TestBoard_setPiece(t *testing.T) {
	b := NewBoard(false)

	b.setPiece(pieceWhiteKing, alg("B1"))
	assert.Equal(t, pieceWhiteKing, b.Piece(alg("B1")))
	assert.Equal(t, colorWhite, b.Color(alg("B1")))

	b.setPiece(pieceBlackKing, alg("B2"))
	assert.Equal(t, pieceBlackKing, b.Piece(alg("B2")))
	assert.Equal(t, colorBlack, b.Color(alg("B2")))

	b.setPiece(pieceBlackRook, alg("D6"))
	assert.Equal(t, pieceBlackRook, b.Piece(alg("D6")))
	assert.Equal(t, colorBlack, b.Color(alg("D6")))

	b.setPiece(pieceWhiteBishop, alg("H8"))
	assert.Equal(t, pieceWhiteBishop, b.Piece(alg("H8")))
	assert.Equal(t, colorWhite, b.Color(alg("H8")))
}

func TestBoard_removePiece(t *testing.T) {
	b := NewBoard(true)

	b.removePiece(alg("B1"))
	assert.Equal(t, pieceNone, b.Piece(alg("B1")))
	assert.Equal(t, colorNone, b.Color(alg("B1")))

	b.removePiece(alg("D6"))
	assert.Equal(t, pieceNone, b.Piece(alg("D6")))
	assert.Equal(t, colorNone, b.Color(alg("D6")))
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

	b = b.MovePiece(7, 23)

	assert.Equal(t, false, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(castlingBlackQueen))

	b = b.MovePiece(0, 16)

	assert.Equal(t, false, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, false, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(castlingBlackQueen))

	b.resetCastlingRights()
	b = b.MovePiece(4, 20)

	assert.Equal(t, false, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, false, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(castlingBlackQueen))
}

func TestBoard_BlackCastlingRights(t *testing.T) {
	b := NewBoard(true)

	assert.Equal(t, true, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(castlingBlackQueen))

	b = b.MovePiece(63, 55)

	assert.Equal(t, true, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, false, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(castlingBlackQueen))

	b = b.MovePiece(56, 48)

	assert.Equal(t, true, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, false, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, false, b.CastlingRights(castlingBlackQueen))

	b.resetCastlingRights()
	b = b.MovePiece(60, 44)

	assert.Equal(t, true, b.CastlingRights(castlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(castlingWhiteQueen))
	assert.Equal(t, false, b.CastlingRights(castlingBlackKing))
	assert.Equal(t, false, b.CastlingRights(castlingBlackQueen))
}

func TestBoard_MoveCount(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, 1, b.MoveCount())
	b = b.MovePiece(0, 8)
	assert.Equal(t, 1, b.MoveCount())
	b = b.MovePiece(63, 55)
	assert.Equal(t, 2, b.MoveCount())
	b = b.MovePiece(1, 9)
	assert.Equal(t, 2, b.MoveCount())
	b = b.MovePiece(62, 54)
	assert.Equal(t, 3, b.MoveCount())
}

func TestBoard_HalfMoveCount_PawnReset(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, 0, b.HalfMoveCount())
	b = b.MovePiece(1, 10)
	assert.Equal(t, 1, b.HalfMoveCount())
	b = b.MovePiece(62, 55)
	assert.Equal(t, 2, b.HalfMoveCount())
	b = b.MovePiece(8, 16)
	assert.Equal(t, 0, b.HalfMoveCount())

}

func TestBoard_HalfMoveCount_CaptureReset(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, 0, b.HalfMoveCount())
	b = b.MovePiece(alg("b1"), alg("c3"))
	assert.Equal(t, 1, b.HalfMoveCount())
	b = b.MovePiece(alg("b8"), alg("c6"))
	assert.Equal(t, 2, b.HalfMoveCount())
	b = b.MovePiece(alg("c3"), alg("d5"))
	assert.Equal(t, 3, b.HalfMoveCount())
	b = b.MovePiece(alg("c6"), alg("b4"))
	assert.Equal(t, 4, b.HalfMoveCount())
	b = b.MovePiece(alg("d5"), alg("b4"))
	assert.Equal(t, 0, b.HalfMoveCount())
}
