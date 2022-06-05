package chess_engine

import (
	"strings"
)

func alg(alg string) int {
	if len(alg) != 2 {
		panic("invalid alg")
	}
	a := strings.ToLower(alg)
	x := int(a[0]) - 97
	y := int(a[1]) - 49

	if x < 0 || x > 7 || y < 0 || y > 7 {
		panic("invalid alg")
	}

	return x + y*8
}

func xy(x, y int) int {
	if x < 1 || x > 8 || y < 1 || y > 8 {
		panic("invalid xy")
	}

	return (x - 1) + (y-1)*8
}

func toXY(i int) (int, int) {
	return i%8 + 1, i/8 + 1
}

func getFileLetter(f int) string {
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

func getPieceName(p Piece) string {
	switch p {
	case PieceNone:
		return " "
	case PieceWhitePawn:
		return "White Pawn"
	case PieceWhiteBishop:
		return "White Bishop"
	case PieceWhiteKnight:
		return "White Knight"
	case PieceWhiteRook:
		return "White Rook"
	case PieceWhiteQueen:
		return "White Queen"
	case PieceWhiteKing:
		return "White King"
	case PieceBlackPawn:
		return "Black Pawn"
	case PieceBlackBishop:
		return "Black Bishop"
	case PieceBlackKnight:
		return "Black Knight"
	case PieceBlackRook:
		return "Black Rook"
	case PieceBlackQueen:
		return "Black Queen"
	case PieceBlackKing:
		return "Black King"
	default:
		panic("invalid piece in getLetterFromPiece()")
	}
}

func getLetterFromPiece(p Piece) string {
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

func getPieceFromLetter(s string) Piece {
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
