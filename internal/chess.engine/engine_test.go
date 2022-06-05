package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmptyBoard(t *testing.T) {
	b := NewBoard(false)

	assert.NotNil(t, b)
	for i := 0; i < 63; i++ {
		assert.Equal(t, PieceNone, b.Piece(Pos(i)))
	}

	// Extra
	assert.Equal(t, ColorWhite, b.ToMove())
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))
	assert.Equal(t, 0, b.getEnPassantTarget())
	assert.Equal(t, 1, b.MoveCount())
	assert.Equal(t, 0, b.HalfMoveCount())
}

func TestNewBoard(t *testing.T) {
	b := NewBoard(true)

	assert.NotNil(t, b)
	assert.Equal(t, PieceBlackRook, b.Piece(Alg("A8")))
	assert.Equal(t, PieceBlackKnight, b.Piece(Alg("B8")))
	assert.Equal(t, PieceBlackBishop, b.Piece(Alg("C8")))
	assert.Equal(t, PieceBlackQueen, b.Piece(Alg("D8")))
	assert.Equal(t, PieceBlackKing, b.Piece(Alg("E8")))
	assert.Equal(t, PieceBlackBishop, b.Piece(Alg("F8")))
	assert.Equal(t, PieceBlackKnight, b.Piece(Alg("G8")))
	assert.Equal(t, PieceBlackRook, b.Piece(Alg("H8")))

	for i := 0; i < 8; i++ {
		assert.Equal(t, PieceBlackPawn, b.Piece(48+Pos(i)))
		assert.Equal(t, PieceNone, b.Piece(40+Pos(i)))
		assert.Equal(t, PieceNone, b.Piece(32+Pos(i)))
		assert.Equal(t, PieceNone, b.Piece(24+Pos(i)))
		assert.Equal(t, PieceNone, b.Piece(16+Pos(i)))
		assert.Equal(t, PieceWhitePawn, b.Piece(8+Pos(i)))
	}

	assert.Equal(t, PieceWhiteRook, b.Piece(Alg("A1")))
	assert.Equal(t, PieceWhiteKnight, b.Piece(Alg("B1")))
	assert.Equal(t, PieceWhiteBishop, b.Piece(Alg("C1")))
	assert.Equal(t, PieceWhiteQueen, b.Piece(Alg("D1")))
	assert.Equal(t, PieceWhiteKing, b.Piece(Alg("E1")))
	assert.Equal(t, PieceWhiteBishop, b.Piece(Alg("F1")))
	assert.Equal(t, PieceWhiteKnight, b.Piece(Alg("G1")))
	assert.Equal(t, PieceWhiteRook, b.Piece(Alg("H1")))

	// Extra
	assert.Equal(t, ColorWhite, b.ToMove())
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))
	assert.Equal(t, 0, b.getEnPassantTarget())
	assert.Equal(t, 1, b.MoveCount())
	assert.Equal(t, 0, b.HalfMoveCount())
}

func TestBoard_Copy(t *testing.T) {
	b := NewBoard(false)

	// Manipulate the board
	b.setPiece(PieceWhiteBishop, Alg("B4"))
	b.setPiece(PieceWhiteKing, Alg("C5"))
	b.setPiece(PieceBlackKing, Alg("C8"))

	// Manipulate extra data
	b.removeCastlingRights(CastlingBlackQueen)
	b.increaseMoveCount()
	b.setEnPassantTarget(4)

	nb := b.Copy()

	assert.NotNil(t, nb)
	assert.True(t, b.Equals(nb))
	assert.Equal(t, nb.board[0], b.board[0])
	assert.Equal(t, nb.board[1], b.board[1])
	assert.Equal(t, nb.board[2], b.board[2])
	assert.Equal(t, nb.board[3], b.board[3])
	assert.Equal(t, nb.extra, b.extra)
}

func TestBoard_Equals(t *testing.T) {
	b := NewBoard(false)

	// Manipulate the board
	b.setPiece(PieceWhiteBishop, Alg("B4"))
	b.setPiece(PieceWhiteKing, Alg("C5"))
	b.setPiece(PieceBlackKing, Alg("C8"))

	// Manipulate extra data
	b.removeCastlingRights(CastlingBlackQueen)
	b.increaseMoveCount()
	b.setEnPassantTarget(4)

	nb := b.Copy()

	assert.NotNil(t, nb)
	assert.True(t, b.Equals(nb))

	b.clearEnPassantTarget()
	assert.False(t, b.Equals(nb))
}

func TestBoard_MovePiece(t *testing.T) {
	b := NewBoard(true)

	nb := b.MovePiece(Alg("b1"), Alg("c3"))

	assert.NotEqual(t, nb.ToMove(), b.ToMove())
	assert.NotEqual(t, nb.board[0], b.board[0])
	assert.NotEqual(t, nb.board[1], b.board[1])
	assert.Equal(t, nb.board[2], b.board[2])
	assert.Equal(t, nb.board[3], b.board[3])
	assert.NotEqual(t, nb.extra, b.extra)
}

func TestBoard_setPiece(t *testing.T) {
	b := NewBoard(false)

	b.setPiece(PieceWhiteKing, Alg("B1"))
	assert.Equal(t, PieceWhiteKing, b.Piece(Alg("B1")))
	assert.Equal(t, ColorWhite, b.Color(Alg("B1")))

	b.setPiece(PieceBlackKing, Alg("B2"))
	assert.Equal(t, PieceBlackKing, b.Piece(Alg("B2")))
	assert.Equal(t, ColorBlack, b.Color(Alg("B2")))

	b.setPiece(PieceBlackRook, Alg("D6"))
	assert.Equal(t, PieceBlackRook, b.Piece(Alg("D6")))
	assert.Equal(t, ColorBlack, b.Color(Alg("D6")))

	b.setPiece(PieceWhiteBishop, Alg("H8"))
	assert.Equal(t, PieceWhiteBishop, b.Piece(Alg("H8")))
	assert.Equal(t, ColorWhite, b.Color(Alg("H8")))
}

func TestBoard_removePiece(t *testing.T) {
	b := NewBoard(true)

	b.removePiece(Alg("B1"))
	assert.Equal(t, PieceNone, b.Piece(Alg("B1")))
	assert.Equal(t, ColorNone, b.Color(Alg("B1")))

	b.removePiece(Alg("D6"))
	assert.Equal(t, PieceNone, b.Piece(Alg("D6")))
	assert.Equal(t, ColorNone, b.Color(Alg("D6")))
}

func TestBoard_ToMove(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, ColorWhite, b.ToMove())

	nb := b.MovePiece(Alg("A2"), Alg("A4"))
	assert.Equal(t, ColorBlack, nb.ToMove())

	b = nb.MovePiece(Alg("H7"), Alg("H5"))
	assert.Equal(t, ColorWhite, b.ToMove())
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

	b = b.MovePiece(Alg("h2"), Alg("h4"))
	b = b.MovePiece(Alg("a7"), Alg("a5"))
	b = b.MovePiece(Alg("h1"), Alg("h2"))

	assert.Equal(t, false, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))

	b = NewBoard(true)
	b = b.MovePiece(Alg("a2"), Alg("a4"))
	b = b.MovePiece(Alg("h7"), Alg("h5"))
	b = b.MovePiece(Alg("a1"), Alg("a2"))

	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, false, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))

	b = NewBoard(true)
	b = b.MovePiece(Alg("e2"), Alg("e4"))
	b = b.MovePiece(Alg("g7"), Alg("g5"))
	b = b.MovePiece(Alg("e1"), Alg("e2"))

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

	b = b.MovePiece(Alg("a2"), Alg("a4"))
	b = b.MovePiece(Alg("h7"), Alg("h5"))
	b = b.MovePiece(Alg("a4"), Alg("a5"))
	b = b.MovePiece(Alg("h8"), Alg("h7"))

	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, false, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, true, b.CastlingRights(CastlingBlackQueen))

	b = b.MovePiece(Alg("b2"), Alg("b4"))
	b = b.MovePiece(Alg("a7"), Alg("a5"))
	b = b.MovePiece(Alg("b4"), Alg("b5"))
	b = b.MovePiece(Alg("a8"), Alg("a7"))

	assert.Equal(t, true, b.CastlingRights(CastlingWhiteKing))
	assert.Equal(t, true, b.CastlingRights(CastlingWhiteQueen))
	assert.Equal(t, false, b.CastlingRights(CastlingBlackKing))
	assert.Equal(t, false, b.CastlingRights(CastlingBlackQueen))

	b = NewBoard(true)
	b = b.MovePiece(Alg("e2"), Alg("e4"))
	b = b.MovePiece(Alg("e7"), Alg("e5"))
	b = b.MovePiece(Alg("c2"), Alg("c4"))
	b = b.MovePiece(Alg("e8"), Alg("e7"))

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
	b = b.MovePiece(Alg("e2"), Alg("e4"))
	assert.Equal(t, 1, b.MoveCount())
	b = b.MovePiece(Alg("a7"), Alg("a5"))
	assert.Equal(t, 2, b.MoveCount())
	b = b.MovePiece(Alg("e4"), Alg("e5"))
	assert.Equal(t, 2, b.MoveCount())
	b = b.MovePiece(Alg("a5"), Alg("a4"))
	assert.Equal(t, 3, b.MoveCount())
}

func TestBoard_HalfMoveCount_PawnReset(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, 0, b.HalfMoveCount())
	b = b.MovePiece(Alg("b1"), Alg("c3"))
	assert.Equal(t, 1, b.HalfMoveCount())
	b = b.MovePiece(Alg("b8"), Alg("c6"))
	assert.Equal(t, 2, b.HalfMoveCount())
	b = b.MovePiece(Alg("e2"), Alg("e4"))
	assert.Equal(t, 0, b.HalfMoveCount())

}

func TestBoard_HalfMoveCount_CaptureReset(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, 0, b.HalfMoveCount())
	b = b.MovePiece(Alg("b1"), Alg("c3"))
	assert.Equal(t, 1, b.HalfMoveCount())
	b = b.MovePiece(Alg("b8"), Alg("c6"))
	assert.Equal(t, 2, b.HalfMoveCount())
	b = b.MovePiece(Alg("c3"), Alg("d5"))
	assert.Equal(t, 3, b.HalfMoveCount())
	b = b.MovePiece(Alg("c6"), Alg("b4"))
	assert.Equal(t, 4, b.HalfMoveCount())
	b = b.MovePiece(Alg("d5"), Alg("b4"))
	assert.Equal(t, 0, b.HalfMoveCount())
}

func TestBoard_EnPassant(t *testing.T) {
	b := NewBoard(true)
	b = b.MovePiece(Alg("b2"), Alg("b4"))
	assert.Equal(t, 1, b.getEnPassantTarget())
	b = b.MovePiece(Alg("b7"), Alg("b6"))
	assert.Equal(t, 0, b.getEnPassantTarget())

	b = NewBoard(true)
	b = b.MovePiece(Alg("b2"), Alg("b3"))
	assert.Equal(t, 0, b.getEnPassantTarget())
	b = b.MovePiece(Alg("b7"), Alg("b5"))
	assert.Equal(t, 9, b.getEnPassantTarget())
	b = b.MovePiece(Alg("c2"), Alg("c3"))
	assert.Equal(t, 0, b.getEnPassantTarget())
}
