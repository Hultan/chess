package chess_engine

type Piece uint64

const (
	pieceNone Piece = 0b0000

	pieceWhitePawn   Piece = 0b0001
	pieceWhiteBishop Piece = 0b0010
	pieceWhiteKnight Piece = 0b0011
	pieceWhiteRook   Piece = 0b0100
	pieceWhiteQueen  Piece = 0b0101
	pieceWhiteKing   Piece = 0b0110

	pieceBlackPawn   Piece = 0b1001
	pieceBlackBishop Piece = 0b1010
	pieceBlackKnight Piece = 0b1011
	pieceBlackRook   Piece = 0b1100
	pieceBlackQueen  Piece = 0b1101
	pieceBlackKing   Piece = 0b1110
)

type Color int

const (
	colorNone Color = iota
	colorWhite
	colorBlack
)

type castlingRight int

const (
	castlingWhiteKing castlingRight = iota
	castlingWhiteQueen
	castlingBlackKing
	castlingBlackQueen
)
