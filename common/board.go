package common

type Piece int

const (
	WhitePawn Piece = iota
	WhiteRook
	WhiteKnight
	WhiteBishop
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackRook
	BlackKnight
	BlackBishop
	BlackQueen
	BlackKing
)

type Tile struct {
	piece Piece
}

type ChessBoard [8][8]Tile

func (b ChessBoard) NewBoard() ChessBoard {
	return NewBoard()
}

func NewBoard() ChessBoard {
	var board [8][8]Tile
	return board
}
