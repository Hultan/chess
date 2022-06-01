package chess_engine

type piece uint64

const (
	pieceNone piece = 0b0000

	pieceWhitePawn   piece = 0b0001
	pieceWhiteBishop piece = 0b0010
	pieceWhiteKnight piece = 0b0011
	pieceWhiteRook   piece = 0b0100
	pieceWhiteQueen  piece = 0b0101
	pieceWhiteKing   piece = 0b0110

	pieceBlackPawn   piece = 0b1001
	pieceBlackBishop piece = 0b1010
	pieceBlackKnight piece = 0b1011
	pieceBlackRook   piece = 0b1100
	pieceBlackQueen  piece = 0b1101
	pieceBlackKing   piece = 0b1110
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
