package chess_engine

const (
	pieceNone uint64 = 0b0000

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

type color int

const (
	colorWhite color = 0b0
	colorBlack color = 0b1
)

type castlingRight int

const (
	castlingWhiteKing castlingRight = iota
	castlingWhiteQueen
	castlingBlackKing
	castlingBlackQueen
)
