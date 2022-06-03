package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_print(t *testing.T) {
	b := NewBoard(false)
	assert.Equal(t, "        \n        \n        \n        \n        \n        \n        \n        \n", b.print())
	b = NewBoard(true)
	assert.Equal(t, "rnbqkbnr\npppppppp\n        \n        \n        \n        \nPPPPPPPP\nRNBQKBNR\n", b.print())

	b = b.MovePiece(1, 18)
	assert.Equal(t, "rnbqkbnr\npppppppp\n        \n        \n        \n  N     \nPPPPPPPP\nR BQKBNR\n", b.print())
}
