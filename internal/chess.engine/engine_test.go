package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmptyBoard(t *testing.T) {
	b := NewBoard(false)

	assert.NotNil(t, b)
	for i := 0; i < 63; i++ {
		assert.Equal(t, PieceNone, b.Piece(i))
	}

	// Extra
	assert.Equal(t, ColorWhite, b.ToMove())
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))
	assert.Equal(t, 1, b.MoveCount())
	assert.Equal(t, 0, b.HalfMoveCount())
}

func TestNewBoard(t *testing.T) {
	b := NewBoard(true)

	assert.NotNil(t, b)
	assert.Equal(t, PieceBlackRook, b.Piece(alg("A8")))
	assert.Equal(t, PieceBlackKnight, b.Piece(alg("B8")))
	assert.Equal(t, PieceBlackBishop, b.Piece(alg("C8")))
	assert.Equal(t, PieceBlackQueen, b.Piece(alg("D8")))
	assert.Equal(t, PieceBlackKing, b.Piece(alg("E8")))
	assert.Equal(t, PieceBlackBishop, b.Piece(alg("F8")))
	assert.Equal(t, PieceBlackKnight, b.Piece(alg("G8")))
	assert.Equal(t, PieceBlackRook, b.Piece(alg("H8")))

	for i := 0; i < 8; i++ {
		assert.Equal(t, PieceBlackPawn, b.Piece(48+i))
		assert.Equal(t, PieceNone, b.Piece(40+i))
		assert.Equal(t, PieceNone, b.Piece(32+i))
		assert.Equal(t, PieceNone, b.Piece(24+i))
		assert.Equal(t, PieceNone, b.Piece(16+i))
		assert.Equal(t, PieceWhitePawn, b.Piece(8+i))
	}

	assert.Equal(t, PieceWhiteRook, b.Piece(alg("A1")))
	assert.Equal(t, PieceWhiteKnight, b.Piece(alg("B1")))
	assert.Equal(t, PieceWhiteBishop, b.Piece(alg("C1")))
	assert.Equal(t, PieceWhiteQueen, b.Piece(alg("D1")))
	assert.Equal(t, PieceWhiteKing, b.Piece(alg("E1")))
	assert.Equal(t, PieceWhiteBishop, b.Piece(alg("F1")))
	assert.Equal(t, PieceWhiteKnight, b.Piece(alg("G1")))
	assert.Equal(t, PieceWhiteRook, b.Piece(alg("H1")))

	// Extra
	assert.Equal(t, ColorWhite, b.ToMove())
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))
	assert.Equal(t, 1, b.MoveCount())
	assert.Equal(t, 0, b.HalfMoveCount())
}

func TestBoard_Copy(t *testing.T) {
	b := NewBoard(false)
	b.setPiece(PieceWhiteBishop, alg("B4"))
	b.setPiece(PieceWhiteKing, alg("C5"))
	b.setPiece(PieceBlackKing, alg("C8"))

	nb := b.Copy()

	assert.NotNil(t, nb)
	assert.Equal(t, nb.board[0], b.board[0])
	assert.Equal(t, nb.board[1], b.board[1])
	assert.Equal(t, nb.board[2], b.board[2])
	assert.Equal(t, nb.board[3], b.board[3])
	assert.Equal(t, nb.extra, b.extra)
}

func TestBoard_MovePiece(t *testing.T) {
	b := NewBoard(true)

	nb := b.MovePiece(alg("b1"), alg("c3"))

	assert.NotEqual(t, nb.ToMove(), b.ToMove())
	assert.NotEqual(t, nb.board[0], b.board[0])
	assert.NotEqual(t, nb.board[1], b.board[1])
	assert.Equal(t, nb.board[2], b.board[2])
	assert.Equal(t, nb.board[3], b.board[3])
	assert.NotEqual(t, nb.extra, b.extra)
}

func TestBoard_setPiece(t *testing.T) {
	b := NewBoard(false)

	b.setPiece(PieceWhiteKing, alg("B1"))
	assert.Equal(t, PieceWhiteKing, b.Piece(alg("B1")))
	assert.Equal(t, ColorWhite, b.Color(alg("B1")))

	b.setPiece(PieceBlackKing, alg("B2"))
	assert.Equal(t, PieceBlackKing, b.Piece(alg("B2")))
	assert.Equal(t, ColorBlack, b.Color(alg("B2")))

	b.setPiece(PieceBlackRook, alg("D6"))
	assert.Equal(t, PieceBlackRook, b.Piece(alg("D6")))
	assert.Equal(t, ColorBlack, b.Color(alg("D6")))

	b.setPiece(PieceWhiteBishop, alg("H8"))
	assert.Equal(t, PieceWhiteBishop, b.Piece(alg("H8")))
	assert.Equal(t, ColorWhite, b.Color(alg("H8")))
}

func TestBoard_removePiece(t *testing.T) {
	b := NewBoard(true)

	b.removePiece(alg("B1"))
	assert.Equal(t, PieceNone, b.Piece(alg("B1")))
	assert.Equal(t, ColorNone, b.Color(alg("B1")))

	b.removePiece(alg("D6"))
	assert.Equal(t, PieceNone, b.Piece(alg("D6")))
	assert.Equal(t, ColorNone, b.Color(alg("D6")))
}

func TestBoard_ToMove(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, ColorWhite, b.ToMove())

	nb := b.MovePiece(alg("A2"), alg("A4"))
	assert.Equal(t, ColorBlack, nb.ToMove())

	b = nb.MovePiece(alg("H7"), alg("H5"))
	assert.Equal(t, ColorWhite, b.ToMove())
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

	assert.Equal(t, ColorWhite, b.ToMove())
	b.toggleToMove()
	assert.Equal(t, ColorBlack, b.ToMove())
	b.toggleToMove()
	assert.Equal(t, ColorWhite, b.ToMove())
}

func TestBoard_WhiteCastlingRights(t *testing.T) {
	b := NewBoard(true)

	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))

	b = b.MovePiece(alg("h2"), alg("h4"))
	b = b.MovePiece(alg("a7"), alg("a5"))
	b = b.MovePiece(alg("h1"), alg("h2"))

	assert.Equal(t, false, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))

	b = NewBoard(true)
	b = b.MovePiece(alg("a2"), alg("a4"))
	b = b.MovePiece(alg("h7"), alg("h5"))
	b = b.MovePiece(alg("a1"), alg("a2"))

	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, false, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))

	b = NewBoard(true)
	b = b.MovePiece(alg("e2"), alg("e4"))
	b = b.MovePiece(alg("g7"), alg("g5"))
	b = b.MovePiece(alg("e1"), alg("e2"))

	assert.Equal(t, false, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, false, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))
}

func TestBoard_BlackCastlingRights(t *testing.T) {
	b := NewBoard(true)

	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))

	b = b.MovePiece(alg("a2"), alg("a4"))
	b = b.MovePiece(alg("h7"), alg("h5"))
	b = b.MovePiece(alg("a4"), alg("a5"))
	b = b.MovePiece(alg("h8"), alg("h7"))

	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, false, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))

	b = b.MovePiece(alg("b2"), alg("b4"))
	b = b.MovePiece(alg("a7"), alg("a5"))
	b = b.MovePiece(alg("b4"), alg("b5"))
	b = b.MovePiece(alg("a8"), alg("a7"))

	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, false, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, false, b.CastlingRights(CastlingBlackQueen))

	b = NewBoard(true)
	b = b.MovePiece(alg("e2"), alg("e4"))
	b = b.MovePiece(alg("e7"), alg("e5"))
	b = b.MovePiece(alg("c2"), alg("c4"))
	b = b.MovePiece(alg("e8"), alg("e7"))

	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, false, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, false, b.CastlingRights(CastlingBlackQueen))
}

func TestBoard_setMoveCount(t *testing.T) {
	b := NewBoard(true)
	b.setMoveCount(0)
	assert.Equal(t, 0, b.MoveCount())
	b.setMoveCount(1)
	assert.Equal(t, 1, b.MoveCount())
	b.setMoveCount(12)
	assert.Equal(t, 12, b.MoveCount())
	b.setMoveCount(53)
	assert.Equal(t, 53, b.MoveCount())
	b.setMoveCount(114)
	assert.Equal(t, 114, b.MoveCount())
}

func TestBoard_MoveCount(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, 1, b.MoveCount())
	b = b.MovePiece(alg("e2"), alg("e4"))
	assert.Equal(t, 1, b.MoveCount())
	b = b.MovePiece(alg("a7"), alg("a5"))
	assert.Equal(t, 2, b.MoveCount())
	b = b.MovePiece(alg("e4"), alg("e5"))
	assert.Equal(t, 2, b.MoveCount())
	b = b.MovePiece(alg("a5"), alg("a4"))
	assert.Equal(t, 3, b.MoveCount())
}

func TestBoard_HalfMoveCount_PawnReset(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, 0, b.HalfMoveCount())
	b = b.MovePiece(alg("b1"), alg("c3"))
	assert.Equal(t, 1, b.HalfMoveCount())
	b = b.MovePiece(alg("b8"), alg("c6"))
	assert.Equal(t, 2, b.HalfMoveCount())
	b = b.MovePiece(alg("e2"), alg("e4"))
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

func TestBoard_EnPassant(t *testing.T) {
	b := NewBoard(true)
	b = b.MovePiece(alg("b2"), alg("b4"))
	assert.Equal(t, 0, b.getEnPassantTarget())
	b = b.MovePiece(alg("b7"), alg("b6"))
	assert.Equal(t, 0, b.getEnPassantTarget())

	b = NewBoard(true)
	b = b.MovePiece(alg("b2"), alg("b3"))
	assert.Equal(t, 0, b.getEnPassantTarget())
	b = b.MovePiece(alg("b7"), alg("b5"))
	assert.Equal(t, 9, b.getEnPassantTarget())
	b = b.MovePiece(alg("c2"), alg("c3"))
	assert.Equal(t, 0, b.getEnPassantTarget())
}
