package chess_engine

type Piece uint64

const (
	PieceNone Piece = 0b0000

	PieceWhitePawn   Piece = 0b0001
	PieceWhiteBishop Piece = 0b0010
	PieceWhiteKnight Piece = 0b0011
	PieceWhiteRook   Piece = 0b0100
	PieceWhiteQueen  Piece = 0b0101
	PieceWhiteKing   Piece = 0b0110

	PieceBlackPawn   Piece = 0b1001
	PieceBlackBishop Piece = 0b1010
	PieceBlackKnight Piece = 0b1011
	PieceBlackRook   Piece = 0b1100
	PieceBlackQueen  Piece = 0b1101
	PieceBlackKing   Piece = 0b1110
)

type Color int

const (
	ColorNone Color = iota
	ColorWhite
	ColorBlack
)

type castlingRight int

const (
	CastlingWhiteKing castlingRight = iota
	CastlingWhiteQueen
	CastlingBlackKing
	CastlingBlackQueen
)
