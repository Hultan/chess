package chess_engine

func (b *Board) getFileLetter(f int) string {
	switch f {
	case 1:
		return "a"
	case 2:
		return "b"
	case 3:
		return "c"
	case 4:
		return "d"
	case 5:
		return "e"
	case 6:
		return "f"
	case 7:
		return "g"
	case 8:
		return "h"
	}
	return ""
}

func (b *Board) getLetterFromPiece(p Piece) string {
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
		panic("invalid piece in getLetterFromPiece()")
	}
}

func (b *Board) getPieceFromLetter(s string) Piece {
	switch s {
	case " ":
		return PieceNone
	case "P":
		return PieceWhitePawn
	case "B":
		return PieceWhiteBishop
	case "N":
		return PieceWhiteKnight
	case "R":
		return PieceWhiteRook
	case "Q":
		return PieceWhiteQueen
	case "K":
		return PieceWhiteKing
	case "p":
		return PieceBlackPawn
	case "b":
		return PieceBlackBishop
	case "n":
		return PieceBlackKnight
	case "r":
		return PieceBlackRook
	case "q":
		return PieceBlackQueen
	case "k":
		return PieceBlackKing
	default:
		panic("invalid piece in getPieceFromLetter()")
	}
}
