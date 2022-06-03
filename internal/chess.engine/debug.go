package chess_engine

func (b *Board) print() string {
	rank := ""
	board := ""
	for i := 0; i < 64; i++ {
		p := b.Piece(i)
		l := b.getPieceLetter(p)
		rank += l
		if (i+1)%8 == 0 {
			board = rank + "\n" + board
			rank = ""
		}
	}
	return board
}

func (b *Board) getPieceLetter(p Piece) string {
	switch p {
	case PieceNone:
		return " "
	case PieceWhitePawn:
		return "P"
	case PieceWhiteBishop:
		return "B"
	case PieceWhiteKnight:
		return "N"
	case PieceWhiteRook:
		return "R"
	case PieceWhiteQueen:
		return "Q"
	case PieceWhiteKing:
		return "K"
	case PieceBlackPawn:
		return "p"
	case PieceBlackBishop:
		return "b"
	case PieceBlackKnight:
		return "n"
	case PieceBlackRook:
		return "r"
	case PieceBlackQueen:
		return "q"
	case PieceBlackKing:
		return "k"
	default:
		panic("invalid piece in getPieceLetter()")
	}
}
