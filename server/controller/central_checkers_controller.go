package controller

import (
	"context"
)

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
	Board []Tile
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
	Message     string
}

// BoardUpdate message
type BoardUpdate struct {
	BoardState  BoardState
	PrevMoveSet []Move
}

// GameRole int
type GameRole int

// GameRole enumeration
const (
	PlayerWhite GameRole = iota
	PlayerBlack
	Spectator
)

// GameStartRequest message
type GameStartRequest struct {
	GameID string
}

// GameStartRole message
type GameStartRole struct {
	PlayerRole GameRole
	GameID     string
}

// GameStartResponse message
type GameStartResponse struct {
	GameData GameStartRole
	Message  string
}

// HandleMakeMoves TODO
func HandleMakeMoves(ctx context.Context, req *MoveRequest) *MoveResponse {
	return nil
}

// HandleBoardUpdateSubscription TODO
func HandleBoardUpdateSubscription(req *BoardSubscriptionRequest) *BoardUpdate {
	return nil
}

// HandleStartGame TODO
func HandleStartGame(req *GameStartRequest) *GameStartResponse {
	return nil
}
