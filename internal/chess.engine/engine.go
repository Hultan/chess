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

func NewBoard(setup bool) *Board {
	b := &Board{}

	if setup {
		b.SetPiece(pieceBlackRook, 56)
		b.SetPiece(pieceBlackKnight, 57)
		b.SetPiece(pieceBlackBishop, 58)
		b.SetPiece(pieceBlackQueen, 59)
		b.SetPiece(pieceBlackKing, 60)
		b.SetPiece(pieceBlackBishop, 61)
		b.SetPiece(pieceBlackKnight, 62)
		b.SetPiece(pieceBlackRook, 63)

		for i := 0; i < 8; i++ {
			b.SetPiece(pieceBlackPawn, 48+i)
			b.SetPiece(pieceWhitePawn, 8+i)
		}

		b.SetPiece(pieceWhiteRook, 0)
		b.SetPiece(pieceWhiteKnight, 1)
		b.SetPiece(pieceWhiteBishop, 2)
		b.SetPiece(pieceWhiteQueen, 3)
		b.SetPiece(pieceWhiteKing, 4)
		b.SetPiece(pieceWhiteBishop, 5)
		b.SetPiece(pieceWhiteKnight, 6)
		b.SetPiece(pieceWhiteRook, 7)

		b.resetCastlingRights()
		b.setMoveCount(1)
	}

	return b
}

func (b *Board) Copy() *Board {
	// Create a new board
	nb := &Board{}

	// Copy the old board
	for i, d := range b.board {
		nb.board[i] = d
	}
	nb.extra = b.extra

	return nb
}

func (b *Board) MovePiece(from, to int) *Board {
	// Create a new board
	nb := b.Copy()

	// Move piece
	p := nb.Piece(from)
	nb.SetPiece(p, to)
	nb.RemovePiece(from)

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

	return nb
}

func (b *Board) SetPiece(piece piece, index int) {
	i := index / 16
	m := index % 16

	p := piece << (m * 4)
	b.board[i] = b.board[i] | uint64(p)
}

func (b *Board) RemovePiece(index int) {
	i := index / 16
	m := index % 16

	p := uint64(0b000 << m)
	b.board[i] = b.board[i] & p
}

func (b *Board) Piece(index int) piece {
	i := index / 16
	m := index % 16

	p := uint64(0b1111 << (m * 4))
	p2 := b.board[i] & p >> (m * 4)

	return piece(p2)
}

func (b *Board) Color(index int) color {
	i := index / 16
	m := index % 16

	p := uint64(0b1000 << (m * 4))
	c := b.board[i] & p >> (m*4 + 3)

	return color(c)
}

func (b *Board) ToMove() color {
	return color(b.extra & 1)
}

func (b *Board) toggleToMove() color {
	b.extra ^= 1 // Switch color to move

	if b.ToMove() == colorWhite {
		b.setMoveCount(b.MoveCount() + 1)
	}

	return b.ToMove()
}

func (b *Board) MoveCount() int {
	return int((b.extra & (0b11111111_11111111 << 16)) >> 16)
}

func (b *Board) setMoveCount(c int) {
	b.extra &= 0b00000000_00000000_11111111_11111111
	b.extra |= uint32(c) << 16
}

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
