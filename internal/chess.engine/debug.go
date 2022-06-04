package chess_engine

import (
	"fmt"
)

func (b *Board) print() string {
	rank := ""
	board := ""
	for i := 0; i < 64; i++ {
		p := b.Piece(i)
		l := b.getLetterFromPiece(p)
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
	if castling != "" {
		board += fmt.Sprintf("Castling : %s\n", castling)
	}
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
