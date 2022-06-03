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
		b.setPiece(pieceBlackRook, 56)
		b.setPiece(pieceBlackKnight, 57)
		b.setPiece(pieceBlackBishop, 58)
		b.setPiece(pieceBlackQueen, 59)
		b.setPiece(pieceBlackKing, 60)
		b.setPiece(pieceBlackBishop, 61)
		b.setPiece(pieceBlackKnight, 62)
		b.setPiece(pieceBlackRook, 63)

		for i := 0; i < 8; i++ {
			b.setPiece(pieceBlackPawn, 48+i)
			b.setPiece(pieceWhitePawn, 8+i)
		}

		b.setPiece(pieceWhiteRook, 0)
		b.setPiece(pieceWhiteKnight, 1)
		b.setPiece(pieceWhiteBishop, 2)
		b.setPiece(pieceWhiteQueen, 3)
		b.setPiece(pieceWhiteKing, 4)
		b.setPiece(pieceWhiteBishop, 5)
		b.setPiece(pieceWhiteKnight, 6)
		b.setPiece(pieceWhiteRook, 7)
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
	if t != colorNone && f == t {
		panic("can't capture own piece")
	}

	if b.Piece(from) == pieceNone {
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
		nb.removeCastlingRights(castlingWhiteQueen)
	}
	if from == 7 {
		nb.removeCastlingRights(castlingWhiteKing)
	}
	if from == 4 {
		nb.removeCastlingRights(castlingWhiteKing)
		nb.removeCastlingRights(castlingWhiteQueen)
	}
	if from == 56 {
		nb.removeCastlingRights(castlingBlackQueen)
	}
	if from == 63 {
		nb.removeCastlingRights(castlingBlackKing)
	}
	if from == 60 {
		nb.removeCastlingRights(castlingBlackKing)
		nb.removeCastlingRights(castlingBlackQueen)
	}

	// Next player to move
	nb.toggleToMove()

	// Adjust move count
	if nb.ToMove() == colorWhite {
		nb.setMoveCount(nb.MoveCount() + 1)
	}

	// Adjust half move count
	if b.Piece(from) == pieceWhitePawn || b.Piece(from) == pieceBlackPawn {
		nb.setHalfMoveCount(0)
	} else if b.Color(to) != colorNone && b.Color(from) != b.Color(to) {
		nb.setHalfMoveCount(0)
	} else {
		nb.setHalfMoveCount(nb.HalfMoveCount() + 1)
	}

	return nb
}

//
// Manipulate board
//

func (b *Board) setPiece(piece piece, index int) {
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

func (b *Board) Piece(index int) piece {
	i := index / 16
	m := index % 16

	p := uint64(0b1111 << (m * 4))
	p2 := b.board[i] & p >> (m * 4)

	return piece(p2)
}

func (b *Board) Color(index int) color {
	p := b.Piece(index)

	if p == pieceNone {
		return colorNone
	}
	if p >= pieceWhitePawn && p <= pieceWhiteKing {
		return colorWhite
	}
	if p >= pieceBlackPawn && p <= pieceBlackKing {
		return colorBlack
	}

	panic("invalid color")
}

//
// Move
//

func (b *Board) ToMove() color {
	if b.extra&1 == 0 {
		return colorWhite
	}
	return colorBlack
}

func (b *Board) toggleToMove() color {
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
	case castlingWhiteKing:
		return b.extra&0b00000001_00000000 > 1
	case castlingWhiteQueen:
		return b.extra&0b00000010_00000000 > 1
	case castlingBlackKing:
		return b.extra&0b00000100_00000000 > 1
	case castlingBlackQueen:
		return b.extra&0b00001000_00000000 > 1
	default:
		return false
	}
}

func (b *Board) removeCastlingRights(c castlingRight) {
	switch c {
	case castlingWhiteKing:
		b.extra &= 0b11111111_11111111_11111110_11111111
	case castlingWhiteQueen:
		b.extra &= 0b11111111_11111111_11111101_11111111
	case castlingBlackKing:
		b.extra &= 0b11111111_11111111_11111011_11111111
	case castlingBlackQueen:
		b.extra &= 0b11111111_11111111_11110111_11111111
	default:
	}
}

func (b *Board) resetCastlingRights() {
	b.extra |= 0b1111_00000000
}
