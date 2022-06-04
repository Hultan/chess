package chess_engine

import (
	"errors"
	"strconv"
	"strings"
)

var InvalidFEN = errors.New("invalid FEN")

// ToFEN generates a FEN string : https://www.chess.com/terms/fen-chess
func (b *Board) ToFEN() string {
	m := make(map[int]func() string, 6)
	m[0] = b.toFenBoard
	m[1] = b.toFenToMove
	m[2] = b.toFenCastling
	m[3] = b.toFenEnPassant
	m[4] = b.toFenHalfMoveCount
	m[5] = b.toFenMoveCount

	// Call each step function
	s := ""
	for i := 0; i < 6; i++ {
		s += m[i]() + " "
	}

	return s[:len(s)-1]
}

// FromFEN parses a fen string and sets up the board accordingly : https://www.chess.com/terms/fen-chess
func FromFEN(fen string) (*Board, error) {
	nb := NewBoard(false)
	fens := strings.Fields(fen)

	m := make(map[int]func(string) error, 6)
	m[0] = nb.fromFenBoard
	m[1] = nb.fromFenToMove
	m[2] = nb.fromFenCastling
	m[3] = nb.fromFenEnPassant
	m[4] = nb.fromFenHalfMoveCount
	m[5] = nb.fromFenMoveCount

	// Call each step function
	for i := 0; i < 6; i++ {
		err := m[i](fens[i])
		if err != nil {
			return nil, err
		}
	}

	return nb, nil
}

//
// Private methods
//

func (b *Board) toFenBoard() string {
	result := ""
	for y := 8; y >= 1; y-- {
		result += b.toFenBoardRow(y) + "/"
	}

	return result[:len(result)-1]
}

func (b *Board) toFenBoardRow(y int) string {
	result := ""
	spaces := 0
	for x := 1; x <= 8; x++ {
		l := b.getLetterFromPiece(b.Piece(xy(x, y)))

		if l == " " {
			// Empty square
			spaces++
		} else {
			// Piece square
			if spaces > 0 {
				result += strconv.Itoa(spaces)
				spaces = 0
			}
			result += l
		}
	}
	if spaces > 0 {
		result += strconv.Itoa(spaces)
		spaces = 0
	}

	return result
}

func (b *Board) toFenToMove() string {
	if b.ToMove() == ColorWhite {
		return "w"
	}
	return "b"
}

func (b *Board) toFenCastling() string {
	result := ""
	if b.CastlingRights(CastlingWhiteKing) {
		result += "K"
	}
	if b.CastlingRights(CastlingWhiteQueen) {
		result += "Q"
	}
	if b.CastlingRights(CastlingBlackKing) {
		result += "k"
	}
	if b.CastlingRights(CastlingBlackQueen) {
		result += "q"
	}
	if result == "" {
		return "-"
	}
	return result
}

func (b *Board) toFenEnPassant() string {
	t := b.getEnPassantTarget()
	if t == 0 {
		return "-"
	}

	result := ""
	if t >= 1 && t <= 8 {
		result = b.getFileLetter(t) + "3"
	} else {
		result = b.getFileLetter(t) + "6"
	}
	return result
}

func (b *Board) toFenHalfMoveCount() string {
	return strconv.Itoa(b.HalfMoveCount())
}

func (b *Board) toFenMoveCount() string {
	return strconv.Itoa(b.MoveCount())
}

func (b *Board) fromFenBoard(s string) error {
	rows := strings.Split(s, "/")
	if len(rows) != 8 {
		return InvalidFEN
	}
	for i, row := range rows {
		err := b.parseFen1Row(8-i, row)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Board) parseFen1Row(y int, row string) error {
	if len(row) < 1 || len(row) > 8 {
		return InvalidFEN
	}

	col := 0
	pos := 0
	valid := "PBNRQKpbnrqk12345678"
	for {
		if col >= 8 {
			break
		}

		letter := row[pos]
		index := strings.Index(valid, string(letter))
		if index < 0 {
			return InvalidFEN
		}

		var typ Piece

		if index >= 0 && index <= 11 {
			col += 1
		} else {
			skip := index - 11
			for i := 0; i < skip; i++ {
				b.setPiece(PieceNone, xy(col+1, y))
			}
			col += skip
			pos++
			continue
		}

		typ = b.getPieceFromLetter(string(letter))
		b.setPiece(typ, xy(col, y))
		pos++
	}
	return nil
}

func (b *Board) fromFenToMove(s string) error {
	if s == "w" {
		b.setToMove(true)
		return nil
	} else if s == "b" {
		b.setToMove(false)
		return nil
	}
	return InvalidFEN
}

func (b *Board) fromFenCastling(s string) error {
	if len(s) > 4 {
		return InvalidFEN
	}

	b.removeCastlingRights(CastlingWhiteKing)
	b.removeCastlingRights(CastlingWhiteQueen)
	b.removeCastlingRights(CastlingBlackKing)
	b.removeCastlingRights(CastlingBlackQueen)

	if s == "-" {
		return nil
	}
	if strings.Contains(s, "K") {
		b.setCastlingRights(CastlingWhiteKing)
	}
	if strings.Contains(s, "Q") {
		b.setCastlingRights(CastlingWhiteQueen)
	}
	if strings.Contains(s, "k") {
		b.setCastlingRights(CastlingBlackKing)
	}
	if strings.Contains(s, "q") {
		b.setCastlingRights(CastlingBlackQueen)
	}
	return nil
}

func (b *Board) fromFenEnPassant(s string) error {
	if s == "-" {
		b.clearEnPassantTarget()
		return nil
	}
	i := alg(s)
	if i >= 16 && i <= 23 {
		b.setEnPassantTarget(i - 15)
	} else {
		b.setEnPassantTarget(i - 47)
	}

	return nil
}

func (b *Board) fromFenHalfMoveCount(s string) error {
	count, err := strconv.Atoi(s)
	if err != nil {
		return InvalidFEN
	}
	b.setHalfMoveCount(count)
	return nil
}

func (b *Board) fromFenMoveCount(s string) error {
	count, err := strconv.Atoi(s)
	if err != nil {
		return InvalidFEN
	}
	b.setMoveCount(count)
	return nil
}
