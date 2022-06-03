package chess_engine

import (
	"fmt"
)

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

	// Extra
	if b.ToMove() == ColorWhite {
		board += "White to move\n"
	} else {
		board += "Black to move\n"
	}
	castling := ""
	if b.CastlingRights(CastlingWhiteKing) {
		castling += "K"
	}
	if b.CastlingRights(CastlingWhiteQueen) {
		castling += "Q"
	}
	if b.CastlingRights(CastlingBlackKing) {
		castling += "k"
	}
	if b.CastlingRights(CastlingBlackQueen) {
		castling += "q"
	}
	board += fmt.Sprintf("Castling : %s\n", castling)

	if e := b.getEnPassantTarget(); e != 0 {
		if e <= 8 {
			board += fmt.Sprintf("En passant : %s3\n", b.getFileLetter(e))
		} else {
			board += fmt.Sprintf("En passant : %s6\n", b.getFileLetter(e-8))
		}
	}
	board += fmt.Sprintf("Half move count : %d\n", b.HalfMoveCount())
	board += fmt.Sprintf("Move count : %d\n", b.MoveCount())

	return board
}

func (b *Board) getFileLetter(e int) string {
	switch e {
	case 1:
		return "A"
	case 2:
		return "B"
	case 3:
		return "C"
	case 4:
		return "D"
	case 5:
		return "E"
	case 6:
		return "F"
	case 7:
		return "G"
	case 8:
		return "H"
	}
	return ""
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
