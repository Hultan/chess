package chess_engine

// Value returns the board value
func (b *Board) Value() int {
	value := 0

	for p := 0; p < 64; p++ {
		pos := Pos(p)

		piece := b.Piece(pos)
		if piece == PieceNone {
			continue
		}

		pieceValue := b.getPieceValue(piece)
		posValue := b.getPiecePositionBonus(pos, piece)
		// fmt.Printf("%s : val(%d) pos(%d)\n", getPieceName(p), pieceValue, posValue)
		value += posValue + pieceValue
	}

	return value
}

// getPieceValue returns the base value for the piece at position pos
func (b *Board) getPieceValue(piece Piece) int {
	value := 0
	switch piece {
	case PieceWhitePawn, PieceBlackPawn:
		value = 100
	case PieceWhiteBishop, PieceBlackBishop:
		value = 330
	case PieceWhiteKnight, PieceBlackKnight:
		value = 320
	case PieceWhiteRook, PieceBlackRook:
		value = 500
	case PieceWhiteQueen, PieceBlackQueen:
		value = 900
	case PieceWhiteKing, PieceBlackKing:
		value = 20000.0
	}

	color := b.ColorFromPiece(piece)
	if color == ColorBlack {
		value *= -1
	}

	return value
}

// getPiecePositionBonus returns the position bonus for the piece
// at position pos.
func (b *Board) getPiecePositionBonus(pos Position, piece Piece) int {
	var bonusTable [8][8]int

	switch piece {
	case PieceWhitePawn, PieceBlackPawn:
		bonusTable = pawnPositionFactor
	case PieceWhiteKnight, PieceBlackKnight:
		bonusTable = knightPositionFactor
	case PieceWhiteBishop, PieceBlackBishop:
		bonusTable = bishopPositionFactor
	case PieceWhiteRook, PieceBlackRook:
		bonusTable = rookPositionFactor
	case PieceWhiteQueen, PieceBlackQueen:
		bonusTable = queenPositionFactor
	case PieceWhiteKing, PieceBlackKing:
		if b.isEndGame() {
			bonusTable = kingPositionFactorEnd
		} else {
			bonusTable = kingPositionFactorEarlyMid
		}
	default:
		return 0
	}

	x, y := pos.ToXY()
	factor := 1
	color := b.ColorFromPiece(piece)
	if color == ColorBlack {
		// Get the symmetric black position
		sym := Sym(int(pos))
		x, y = sym.ToXY()
		factor = -1
	}

	return factor * bonusTable[8-y][x-1]
}
