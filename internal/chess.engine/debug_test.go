package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_printEmptyBoard(t *testing.T) {
	b := NewBoard(false)
	assert.Equal(t, "        \n        \n        \n        \n        \n        \n        \n        \nWhite to move\nCastling : KQkq\nHalf move count : 0\nMove count : 1\n", b.print())
}

func TestBoard_print(t *testing.T) {
	b := NewBoard(true)
	assert.Equal(t, "rnbqkbnr\npppppppp\n        \n        \n        \n        \nPPPPPPPP\nRNBQKBNR\nWhite to move\nCastling : KQkq\nHalf move count : 0\nMove count : 1\n", b.print())

	b = b.MovePiece(Alg("b1"), Alg("c3"))
	assert.Equal(t, "rnbqkbnr\npppppppp\n        \n        \n        \n  N     \nPPPPPPPP\nR BQKBNR\nBlack to move\nCastling : KQkq\nHalf move count : 1\nMove count : 1\n", b.print())

	b = b.MovePiece(Alg("b8"), Alg("c6"))
	assert.Equal(t, "r bqkbnr\npppppppp\n  n     \n        \n        \n  N     \nPPPPPPPP\nR BQKBNR\nWhite to move\nCastling : KQkq\nHalf move count : 2\nMove count : 2\n", b.print())

	b = b.MovePiece(Alg("d2"), Alg("d4"))
	assert.Equal(t, "r bqkbnr\npppppppp\n  n     \n        \n   P    \n  N     \nPPP PPPP\nR BQKBNR\nBlack to move\nCastling : KQkq\nEn passant : c3\nHalf move count : 0\nMove count : 2\n", b.print())
}
