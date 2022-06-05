package chess_engine

type Evaluator struct {
	board *Board
}

func NewEvaluator(b *Board) *Evaluator {
	return &Evaluator{b}
}

// Value returns the board value
func (e *Evaluator) Value() int {
	value := 0
	for y := 8; y >= 1; y-- {
		for x := 1; x <= 8; x++ {
			pos := XY(x, y)
			p := e.board.Piece(pos)
			c := e.board.ColorFromPiece(p)

			if p == PieceNone {
				continue
			}
			pieceValue := 0
			posValue := 0
			if c == ColorWhite {
				pieceValue += e.getBasePieceValue(pos)
				posValue += e.getPiecePositionBonus(pos)
			} else {
				pieceValue += e.getBasePieceValue(pos)
				posValue += e.getPiecePositionBonus(pos)
			}
			// fmt.Printf("%s : val(%d) pos(%d)\n", getPieceName(p), pieceValue, posValue)
			value += posValue + pieceValue
		}
	}

	return value
}

// getBasePieceValue returns the base value for the piece at position pos
func (e *Evaluator) getBasePieceValue(pos Position) int {
	piece := e.board.Piece(pos)
	color := e.board.Color(pos)

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
	if color == ColorBlack {
		value *= -1
	}
	return value
}

// getPiecePositionBonus returns the position bonus for the piece
// at position pos.
func (e *Evaluator) getPiecePositionBonus(pos Position) int {
	var bonusTable [8][8]int
	color := e.board.Color(pos)

	switch e.board.Piece(pos) {
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
		if e.isEndGame() {
			bonusTable = kingPositionFactorEnd
		} else {
			bonusTable = kingPositionFactorEarlyMid
		}
	default:
		return 0
	}

	x, y := pos.ToXY()
	if color == ColorBlack {
		return -bonusTable[y-1][x-1]
	}
	return bonusTable[8-y][x-1]
}

// TODO : Fix is end game (https://www.chessprogramming.org/Simplified_Evaluation_Function)
//    1. Both sides have no queens or
//    2. Every side which has a queen has additionally no other pieces or one minor piece maximum.
// isEndGame currently returns true
func (e *Evaluator) isEndGame() bool {
	return false
}
