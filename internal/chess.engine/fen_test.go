package chess_engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_FENAndPrint(t *testing.T) {
	tests := []struct {
		name  string
		fen   string
		print string
	}{
		{
			"FEN 1",
			"8/6P1/pn1RPKR1/8/p1k2b2/5B1p/2N3bp/3q4 w - - 0 1",
			"        \n      P \npn RPKR \n        \np k  b  \n     B p\n  N   bp\n   q    \nWhite to move\nHalf move count : 0\nMove count : 1\n",
		},
		{
			"FEN 2",
			"3r2Nr/1pp3p1/1P1Ppq2/p3p2B/1n1QP2p/pN2PRPK/PPknP2B/4b2b w - - 0 1",
			"   r  Nr\n pp   p \n P Ppq  \np   p  B\n n QP  p\npN  PRPK\nPPknP  B\n    b  b\nWhite to move\nHalf move count : 0\nMove count : 1\n",
		},
		{
			"FEN 3",
			"r3k2r/pp1n2pp/2p2q2/b2p1n2/BP1Pp3/P1N2P2/2PB2PP/R2Q1RK1 w kq b3 0 13",
			"r   k  r\npp n  pp\n  p  q  \nb  p n  \nBP Pp   \nP N  P  \n  PB  PP\nR  Q RK \nWhite to move\nCastling : kq\nEn passant : b3\nHalf move count : 0\nMove count : 13\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := FromFEN(tt.fen)
			assert.Nil(t, err, "", "Failed (FromFEN()): "+tt.name)
			assert.Equalf(t, tt.print, b.print(), "Failed (print()): "+tt.name)
			got := b.ToFEN()
			assert.Equalf(t, tt.fen, got, "Failed (ToFEN()): "+tt.name)
		})
	}
}
