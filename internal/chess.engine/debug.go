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

func (b *Board) getPieceLetter(p piece) string {
	switch p {
	case pieceNone:
		return " "
	case pieceWhitePawn:
		return "P"
	case pieceWhiteBishop:
		return "B"
	case pieceWhiteKnight:
		return "N"
	case pieceWhiteRook:
		return "R"
	case pieceWhiteQueen:
		return "Q"
	case pieceWhiteKing:
		return "K"
	case pieceBlackPawn:
		return "p"
	case pieceBlackBishop:
		return "b"
	case pieceBlackKnight:
		return "n"
	case pieceBlackRook:
		return "r"
	case pieceBlackQueen:
		return "q"
	case pieceBlackKing:
		return "k"
	default:
		panic("invalid piece in getPieceLetter()")
	}
}
