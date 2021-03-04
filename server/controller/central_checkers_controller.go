package controller

// Tile int
type Tile int

// Tile enumeration
const (
	NoPiece Tile = iota
	White
	Black
	WhitePromoted
	BlackPromoted
)

// Move message
type Move struct {
	IndexFrom uint32
	IndexTo   uint32
}

// BoardState message
type BoardState struct {
	Board *[]*Tile
}

// BoardSubscriptionRequest message
type BoardSubscriptionRequest struct {
	GameID string
}

// MoveRequest message
type MoveRequest struct {
	MoveSet []Move
}

// MoveResponse message
type MoveResponse struct {
	MoveSuccess bool
	Message     []string
}

// BoardUpdate message
type BoardUpdate struct {
	BoardState  BoardState
	PrevMoveSet *[]*Move
}
