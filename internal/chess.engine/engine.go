package chess_engine

// Board represents a chess board
//  - board - each uint64 contains the piece and color of two ranks
//  - extra - bit 0 - color to move
//          - bit 1-7 - half move count, since last capture or pawn advance
//          - bit 8-11 - castling rights : WK,WQ,BK,BQ
//          - bit 12-15 - en passant target, can only occur on 16 squares
//          - bit 16-31 - move count
type Board struct {
	board [4]uint64
	extra uint32
}

//
// NewBoard
//

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
	// TODO : En passant target
	b.setMoveCount(1)

	return b
}

//
// Generate a new board
//

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
	f, t := b.Color(from), b.Color(to)
	if f != b.ToMove() {
		panic("moving out of turn")
	}
	if t != ColorNone && f == t {
		panic("can't capture own piece")
	}

	if b.Piece(from) == PieceNone {
		panic("can't move a none piece")
	}

	// Create a new board
	nb := b.Copy()

	// Move piece
	p := nb.Piece(from)
	nb.setPiece(p, to)
	nb.removePiece(from)

	// Castling
	if from == 0 {
		nb.removeCastlingRights(CastlingWhiteQueen)
	}
	if from == 7 {
		nb.removeCastlingRights(CastlingWhiteKing)
	}
	if from == 4 {
		nb.removeCastlingRights(CastlingWhiteKing)
		nb.removeCastlingRights(CastlingWhiteQueen)
	}
	if from == 56 {
		nb.removeCastlingRights(CastlingBlackQueen)
	}
	if from == 63 {
		nb.removeCastlingRights(CastlingBlackKing)
	}
	if from == 60 {
		nb.removeCastlingRights(CastlingBlackKing)
		nb.removeCastlingRights(CastlingBlackQueen)
	}

	// Next player to move
	nb.toggleToMove()

	// Adjust move count
	if nb.ToMove() == ColorWhite {
		nb.setMoveCount(nb.MoveCount() + 1)
	}

	// Adjust half move count
	if b.Piece(from) == PieceWhitePawn || b.Piece(from) == PieceBlackPawn {
		nb.setHalfMoveCount(0)
	} else if b.Color(to) != ColorNone && b.Color(from) != b.Color(to) {
		nb.setHalfMoveCount(0)
	} else {
		nb.setHalfMoveCount(nb.HalfMoveCount() + 1)
	}

	return nb
}

//
// Manipulate board
//

func (b *Board) setPiece(piece Piece, index int) {
	i, m := index/16, index%16

	b.board[i] = b.board[i] | uint64(piece<<(m*4))
}

func (b *Board) removePiece(index int) {
	i, m := index/16, index%16

	b.board[i] &= ^uint64(0b1111 << (m * 4))
}

//
// Get board info
//

func (b *Board) Piece(index int) Piece {
	i := index / 16
	m := index % 16

	p := uint64(0b1111 << (m * 4))
	p2 := b.board[i] & p >> (m * 4)

	return Piece(p2)
}

func (b *Board) Color(index int) Color {
	p := b.Piece(index)

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

//
// Move
//

func (b *Board) ToMove() Color {
	if b.extra&1 == 0 {
		return ColorWhite
	}
	return ColorBlack
}

func (b *Board) toggleToMove() Color {
	b.extra ^= 1 // Switch color to move

	return b.ToMove()
}

//
// Move count
//

func (b *Board) MoveCount() int {
	return int((b.extra & 0b11111111_11111111_00000000_00000000) >> 16)
}

func (b *Board) setMoveCount(c int) {
	b.extra &= 0b00000000_00000000_11111111_11111111
	b.extra |= uint32(c) << 16
}

//
// Half move count
//

func (b *Board) HalfMoveCount() int {
	return int((b.extra & 0b00000000_00000000_00000000_11111110) >> 1)
}

func (b *Board) setHalfMoveCount(c int) {
	b.extra &= 0b11111111_11111111_11111111_00000001
	b.extra |= uint32(c) << 1
}

//
// Castling
//

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
