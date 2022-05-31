package chess_engine

import (
	"fmt"
)

type Board struct {
	board [4]int
	extra uint32
}

func NewBoard(setup bool) *Board {
	b := &Board{}

	if setup {
		b.SetPiece(pieceBlackRook, 56)
		b.SetPiece(pieceBlackKnight, 57)
		b.SetPiece(pieceBlackBishop, 58)
		b.SetPiece(pieceBlackQueen, 59)
		b.SetPiece(pieceBlackKing, 60)
		b.SetPiece(pieceBlackBishop, 61)
		b.SetPiece(pieceBlackKnight, 62)
		b.SetPiece(pieceBlackRook, 63)

		for i := 0; i < 8; i++ {
			b.SetPiece(pieceBlackPawn, 48+i)
			b.SetPiece(pieceWhitePawn, 8+i)
		}

		b.SetPiece(pieceWhiteRook, 0)
		b.SetPiece(pieceWhiteKnight, 1)
		b.SetPiece(pieceWhiteBishop, 2)
		b.SetPiece(pieceWhiteQueen, 3)
		b.SetPiece(pieceWhiteKing, 4)
		b.SetPiece(pieceWhiteBishop, 5)
		b.SetPiece(pieceWhiteKnight, 6)
		b.SetPiece(pieceWhiteRook, 7)
	}

	return b
}

func (b *Board) Copy() *Board {
	// Create a new board
	nb := &Board{}

	// Copy the old board
	for i, d := range b.board {
		nb.board[i] = d
	}
	nb.extra = b.extra

	return nb
}

func (b *Board) MovePiece(from, to int) *Board {
	// Create a new board
	nb := b.Copy()

	// Move piece
	p := nb.Piece(from)
	nb.SetPiece(p, to)
	nb.RemovePiece(from)

	// Next player to move
	fmt.Println(nb.extra)
	nb.extra |= 1
	fmt.Println(nb.extra)

	return nb
}

func (b *Board) SetPiece(piece int, index int) {
	i := index / 16
	m := index % 16

	p := piece << (m * 4)
	b.board[i] = b.board[i] | p

	// fmt.Println(i)
	// fmt.Println(m)
	// fmt.Printf("%b\n", piece)
	// fmt.Printf("%b\n", p)
	// fmt.Printf("%b\n", b.board[i])
}

func (b *Board) RemovePiece(index int) {
	i := index / 16
	m := index % 16

	p := int(0b000 << (m - 1))
	b.board[i] = b.board[i] & p
}

func (b *Board) Piece(index int) int {
	i := index / 16
	m := index % 16

	p := int(0b1111 << (m * 4))
	piece := b.board[i] & p >> (m * 4)

	fmt.Println(i)
	fmt.Println(m)
	fmt.Printf("%b\n", p)
	fmt.Printf("%b\n", piece)

	return piece
}

func (b *Board) Color(index int) int {
	i := index / 16
	m := index % 16

	p := int(0b1000 << (m * 4))
	c := b.board[i] & p >> (m*4 + 3)

	// fmt.Println(i)
	// fmt.Println(m)
	// fmt.Printf("%b\n", p)
	// fmt.Printf("%b\n", c)

	return c
}

func (b *Board) ToMove() int {
	return int(b.extra & 1)
}
