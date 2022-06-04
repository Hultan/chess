package chess_engine

import (
	"errors"
	"strconv"
	"strings"
)

var InvalidFEN = errors.New("invalid FEN")

// ToFEN generates a FEN string : https://www.chess.com/terms/fen-chess
func (b *Board) ToFEN() string {
	s := b.fenSection1() + " "
	s += b.fenSection2() + " "
	s += b.fenSection3() + " "
	s += b.fenSection4() + " "
	s += b.fenSection5() + " "
	s += b.fenSection6()

	return s
}

// FromFEN parses a fen string and sets up the board accordingly : https://www.chess.com/terms/fen-chess
func (b *Board) FromFEN(fen string) error {
	fens := strings.Fields(fen)

	err := b.parseFen1(fens[0])
	if err != nil {
		return err
	}
	err = b.parseFen2(fens[1])
	if err != nil {
		return err
	}
	err = b.parseFen3(fens[2])
	if err != nil {
		return err
	}
	err = b.parseFen4(fens[3])
	if err != nil {
		return err
	}
	err = b.parseFen5(fens[4])
	if err != nil {
		return err
	}
	err = b.parseFen6(fens[5])
	if err != nil {
		return err
	}

	return nil
}

//
// Private methods
//

func (b *Board) fenSection1() string {
	result := ""
	spaces := 0
	for y := 8; y >= 1; y-- {
		for x := 1; x <= 8; x++ {
			l := b.getLetterFromPiece(b.Piece(xy(x, y)))

			if l == " " {
				spaces++
			} else {
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
		if y != 1 {
			result += "/"
		}
	}

	return result
}

func (b *Board) fenSection2() string {
	if b.ToMove() == ColorWhite {
		return "w"
	}
	return "b"
}

func (b *Board) fenSection3() string {
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

func (b *Board) fenSection4() string {
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

func (b *Board) fenSection5() string {
	return strconv.Itoa(b.HalfMoveCount())
}

func (b *Board) fenSection6() string {
	return strconv.Itoa(b.MoveCount())
}

func (b *Board) parseFen1(s string) error {
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

func (b *Board) parseFen2(s string) error {
	if s == "w" {
		b.setToMove(true)
		return nil
	} else if s == "b" {
		b.setToMove(false)
		return nil
	}
	return InvalidFEN
}

func (b *Board) parseFen3(s string) error {
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

func (b *Board) parseFen4(s string) error {
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

func (b *Board) parseFen5(s string) error {
	count, err := strconv.Atoi(s)
	if err != nil {
		return InvalidFEN
	}
	b.setHalfMoveCount(count)
	return nil
}

func (b *Board) parseFen6(s string) error {
	count, err := strconv.Atoi(s)
	if err != nil {
		return InvalidFEN
	}
	b.setMoveCount(count)
	return nil
}
