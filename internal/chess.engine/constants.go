package chess_engine

const (
	pieceNone        uint64 = 0b0000
	pieceWhitePawn   uint64 = 0b0001
	pieceWhiteBishop uint64 = 0b0010
	pieceWhiteKnight uint64 = 0b0011
	pieceWhiteRook   uint64 = 0b0100
	pieceWhiteQueen  uint64 = 0b0101
	pieceWhiteKing   uint64 = 0b0110

	pieceBlackPawn   uint64 = 0b1001
	pieceBlackBishop uint64 = 0b1010
	pieceBlackKnight uint64 = 0b1011
	pieceBlackRook   uint64 = 0b1100
	pieceBlackQueen  uint64 = 0b1101
	pieceBlackKing   uint64 = 0b1110
)

const (
	colorWhite int = 0b0
	colorBlack int = 0b1
)
