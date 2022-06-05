package chess_engine

import (
	"errors"
)

// Board represents a chess board
//  - board - each uint64 contains the piece and color of two ranks
//  - extra - bit 0 - color to move
//          - bit 1-7 - half move count, since last capture or pawn advance
//          - bit 8-11 - castling rights : WK,WQ,BK,BQ
//          - bit 12-16 - en passant target, can only occur on 16 squares (0-16)
//                  0    no en passant target
//                  1-8  rank 3
//                  9-16 rank 5
//          - bit 17-31 - move count
type Board struct {
	board [4]uint64
	extra uint32
}

func NewBoard(setup bool) *Board {
	b := &Board{}

	if setup {
		b.setPiece(PieceBlackRook, 56)
		b.setPiece(PieceBlackKnight, 57)
		b.setPiece(PieceBlackBishop, 58)
		b.setPiece(PieceBlackQueen, 59)
		b.setPiece(PieceBlackKing, 60)
		b.setPiece(PieceBlackBishop, 61)
		b.setPiece(PieceBlackKnight, 62)
		b.setPiece(PieceBlackRook, 63)

		for i := 0; i < 8; i++ {
			b.setPiece(PieceBlackPawn, 48+i)
			b.setPiece(PieceWhitePawn, 8+i)
		}

		b.setPiece(PieceWhiteRook, 0)
		b.setPiece(PieceWhiteKnight, 1)
		b.setPiece(PieceWhiteBishop, 2)
		b.setPiece(PieceWhiteQueen, 3)
		b.setPiece(PieceWhiteKing, 4)
		b.setPiece(PieceWhiteBishop, 5)
		b.setPiece(PieceWhiteKnight, 6)
		b.setPiece(PieceWhiteRook, 7)
	}

	// Extra
	b.resetCastlingRights()
	b.clearEnPassantTarget()
	b.setMoveCount(1)

	return b
}

func (b *Board) Equals(o *Board) bool {
	for i := 0; i < 4; i++ {
		if b.board[i] != o.board[i] {
			return false
		}
	}
	if b.extra != o.extra {
		return false
	}

	return true
}

func (b *Board) Copy() *Board {
	// Create a new board
	nb := &Board{}

	// Copy the old board
	for i := 0; i < 4; i++ {
		nb.board[i] = b.board[i]
	}

	// Copy extra information
	nb.extra = b.extra

	return nb
}

func (b *Board) MovePiece(from, to int) *Board {
	if e := b.checkValidMoveBasic(from, to); e != nil {
		panic(e)
	}

	// Create a new board
	nb := b.Copy()

	// Move piece
	p := nb.Piece(from)
	nb.setPiece(p, to)
	nb.removePiece(from)

	// En passant
	nb.checkEnPassant(b, from, to)

	// Castling
	nb.checkCastlingRights(from)

	// Next player to move
	nb.toggleToMove()

	// Adjust move count
	nb.increaseMoveCount()

	// Adjust half move count
	nb.increaseHalfMoveCount(b, from, to)

	return nb
}

func (b *Board) Piece(index int) Piece {
	i, m := index/16, index%16

	p := uint64(0b1111 << (m * 4))
	p2 := b.board[i] & p >> (m * 4)

	return Piece(p2)
}

func (b *Board) Color(index int) Color {
	p := b.Piece(index)
	return b.ColorFromPiece(p)

	panic("invalid color")
}

func (b *Board) ColorFromPiece(p Piece) Color {
	if p == PieceNone {
		return ColorNone
	}
	if p >= PieceWhitePawn && p <= PieceWhiteKing {
		return ColorWhite
	}
	if p >= PieceBlackPawn && p <= PieceBlackKing {
		return ColorBlack
	}

	panic("invalid color")
}

func (b *Board) ToMove() Color {
	if b.extra&1 == 0 {
		return ColorWhite
	}
	return ColorBlack
}

func (b *Board) MoveCount() int {
	return int((b.extra & 0b11111111_11111110_00000000_00000000) >> 17)
}

func (b *Board) HalfMoveCount() int {
	return int((b.extra & 0b00000000_00000000_00000000_11111110) >> 1)
}

func (b *Board) CastlingRights(c castlingRight) bool {
	switch c {
	case CastlingWhiteKing:
		return b.extra&0b00000001_00000000 > 1
	case CastlingWhiteQueen:
		return b.extra&0b00000010_00000000 > 1
	case CastlingBlackKing:
		return b.extra&0b00000100_00000000 > 1
	case CastlingBlackQueen:
		return b.extra&0b00001000_00000000 > 1
	default:
		return false
	}
}

//
// Private functions
//

func (b *Board) setPiece(piece Piece, index int) {
	i, m := index/16, index%16

	b.board[i] = b.board[i] | uint64(piece<<(m*4))
}

func (b *Board) removePiece(index int) {
	i, m := index/16, index%16

	b.board[i] &= ^uint64(0b1111 << (m * 4))
}

func (b *Board) checkValidMoveBasic(from, to int) error {
	f, t := b.Color(from), b.Color(to)
	if f != b.ToMove() {
		return errors.New("moving out of turn")
	}
	if t != ColorNone && f == t {
		return errors.New("can't capture own piece")
	}
	if b.Piece(from) == PieceNone {
		return errors.New("can't move a none piece")
	}

	return nil
}

//
// To move
//

func (b *Board) setToMove(whiteToMove bool) {
	if whiteToMove {
		b.extra &= 0b11111111_11111111_11111111_11111110
	} else {
		b.extra |= 1
	}
}

func (b *Board) toggleToMove() Color {
	b.extra ^= 1 // Switch color to move

	return b.ToMove()
}

//
// Move count
//

func (b *Board) setMoveCount(c int) {
	b.extra &= 0b00000000_00000001_11111111_11111111
	b.extra |= uint32(c) << 17
}

func (b *Board) increaseMoveCount() {
	if b.ToMove() == ColorWhite {
		b.setMoveCount(b.MoveCount() + 1)
	}
}

//
// Half move count
//

func (b *Board) setHalfMoveCount(c int) {
	b.extra &= 0b11111111_11111111_11111111_00000001
	b.extra |= uint32(c) << 1
}

func (b *Board) increaseHalfMoveCount(oldBoard *Board, from, to int) {
	if oldBoard.Piece(from) == PieceWhitePawn || oldBoard.Piece(from) == PieceBlackPawn {
		b.setHalfMoveCount(0)
	} else if oldBoard.Color(to) != ColorNone && oldBoard.Color(from) != oldBoard.Color(to) {
		b.setHalfMoveCount(0)
	} else {
		b.setHalfMoveCount(b.HalfMoveCount() + 1)
	}
}

//
// Castling
//

func (b *Board) setCastlingRights(c castlingRight) {
	switch c {
	case CastlingWhiteKing:
		b.extra |= 0b00000000_00000000_00000001_00000000
	case CastlingWhiteQueen:
		b.extra |= 0b00000000_00000000_00000010_00000000
	case CastlingBlackKing:
		b.extra |= 0b00000000_00000000_00000100_00000000
	case CastlingBlackQueen:
		b.extra |= 0b00000000_00000000_0000100000000000
	default:
	}
}

func (b *Board) removeCastlingRights(c castlingRight) {
	switch c {
	case CastlingWhiteKing:
		b.extra &= 0b11111111_11111111_11111110_11111111
	case CastlingWhiteQueen:
		b.extra &= 0b11111111_11111111_11111101_11111111
	case CastlingBlackKing:
		b.extra &= 0b11111111_11111111_11111011_11111111
	case CastlingBlackQueen:
		b.extra &= 0b11111111_11111111_11110111_11111111
	default:
	}
}

func (b *Board) resetCastlingRights() {
	b.extra |= 0b1111_00000000
}

func (b *Board) checkCastlingRights(from int) {
	if from == alg("a1") {
		b.removeCastlingRights(CastlingWhiteQueen)
	}
	if from == alg("h1") {
		b.removeCastlingRights(CastlingWhiteKing)
	}
	if from == alg("e1") {
		b.removeCastlingRights(CastlingWhiteKing)
		b.removeCastlingRights(CastlingWhiteQueen)
	}
	if from == alg("a8") {
		b.removeCastlingRights(CastlingBlackQueen)
	}
	if from == alg("h8") {
		b.removeCastlingRights(CastlingBlackKing)
	}
	if from == alg("e8") {
		b.removeCastlingRights(CastlingBlackKing)
		b.removeCastlingRights(CastlingBlackQueen)
	}
}

//
// En passant
//

func (b *Board) checkEnPassant(oldBoard *Board, from, to int) {
	if oldBoard.Piece(from) == PieceWhitePawn && from >= 8 && from <= 15 && to-from == 16 {
		b.setEnPassantTarget(from - 8)
		return
	}
	if oldBoard.Piece(from) == PieceBlackPawn && from >= 48 && from <= 55 && from-to == 16 {
		b.setEnPassantTarget(from - 48 + 8)
		return
	}
	b.clearEnPassantTarget()
}

func (b *Board) setEnPassantTarget(i int) {
	b.extra &= 0b11111111_11111110_00001111_11111111
	b.extra |= uint32(i) << 12
}

func (b *Board) clearEnPassantTarget() {
	b.extra &= 0b11111111_11111110_00001111_11111111
}

func (b *Board) getEnPassantTarget() int {
	return int((b.extra & 0b00000000_00000001_11110000_00000000) >> 12)
}
