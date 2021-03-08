package controller

import (
	"context"

	"github.com/MettyS/checkers/server/gameplay"
	d "github.com/MettyS/checkers/server/shared"

	"fmt"
)

// GameList map list of all game instances, pairing a GameID with a game handler
var GameList map[string]gameplay.Game

// d.GameStartRequest expected fields:
// GameID string - optional
func validateGameStartRequest(req d.GameStartRequest) (bool, string) {
	return true, ""
}

// d.MoveRequest expected fields:
// MoveSet []Move - required, not nil
// PlayerID string - required, not empty
// GameID string - required, not empty
func validateMoveRequest(req d.MoveRequest) (bool, string) {
	if req.MoveSet == nil {
		return false, "Field MoveSet required."
	} else if req.PlayerID == "" {
		return false, "Field PlayerID must not be empty."
	} else if req.GameID == "" {
		return false, "Field GameID must not be empty."
	}
	return true, ""
}

func validateGameExists(gameID string) (bool, string) {
	if _, exists := GameList[gameID]; !exists {
		return false, "Game does not exist."
	}
	return true, ""
}

// HandleMakeMoves TODO
func HandleMakeMoves(ctx context.Context, req d.MoveRequest) (d.MoveResponse, error) {
	missingField, message := validateMoveRequest(req)

	if missingField {
		return d.MoveResponse{
			MoveSuccess: false,
			Message:     message,
		}, nil
	}

	validGame, message := validateGameExists(req.GameID)
	if validGame != true {
		return d.MoveResponse{
			MoveSuccess: false,
			Message:     message,
		}, nil
	}

	if gameInstance, exists := GameList[req.GameID]; exists {
		success, message := gameInstance.AttemptMoves(req.PlayerID, req.MoveSet)
		fmt.Printf("%v, %v", success, message)
	} else {

	}
	return d.MoveResponse{}, nil
}

// HandleBoardUpdateSubscription TODO
func HandleBoardUpdateSubscription(req d.BoardSubscriptionRequest) (d.BoardUpdate, error) {
	return d.BoardUpdate{}, nil
}

// HandleStartGame TODO
func HandleStartGame(req d.GameStartRequest) (d.GameStartResponse, error) {
	gameID := req.GameID
	playerID := "RandomKeyChangeThis"

	var playerRole d.GameRole
	var message string
	if gameInstance, exists := GameList[gameID]; exists {
		playerRole, message = gameInstance.AddParticipant(playerID)
	} else {
		gameID = "RandomGameIDOverwriteChangeThis"
		gameInstance = gameplay.CreateGameHandler()
		playerRole, message = gameInstance.AddParticipant(playerID)

		GameList[gameID] = gameInstance
	}
	return d.GameStartResponse{
		GameData: d.GameStartRole{
			PlayerRole: playerRole,
			GameID:     gameID,
			PlayerID:   playerID,
		},
		Message: message,
	}, nil
}
