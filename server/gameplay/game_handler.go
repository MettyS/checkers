package gameplay

import (
	"fmt"

	d "github.com/MettyS/checkers/server/shared"
)

// Game TODO
type Game struct {
	Board []d.Tile

	CurrentPlayer string // playerID
	Participants  map[string]d.GameRole
}

// AddParticipant add a participant to the game
func (g *Game) AddParticipant(playerID string) (d.GameRole, string) {
	if player, exists := g.Participants[playerID]; exists {
		return player, "Player already exists (shouldn't happen with random tokens)"
	} else if numPlayers := len(g.Participants); numPlayers < 2 {
		if numPlayers == 0 {
			player = d.PlayerWhite
		} else {
			player = d.PlayerBlack
		}
		g.Participants[playerID] = player
		return player, ""
	} else {
		player = d.Spectator
		g.Participants[playerID] = player
		return player, ""
	}
}

// CreateGameHandler create a new instance of a game handler
func CreateGameHandler() Game {
	game := Game{
		Board:         createDefaultBoard(),
		CurrentPlayer: "",
		Participants:  make(map[string]d.GameRole),
	}
	return game
}

func createDefaultBoard() []d.Tile {
	board := make([]d.Tile, d.BoardSize)

	for i := 0; i < 32; i++ {
		if i < 12 {
			board[i] = d.Tile(d.Black)
		} else if i < 20 {
			board[i] = d.Tile(d.NoPiece)
		} else {
			board[i] = d.Tile(d.White)
		}
	}
	return board
}

// AttemptMoves validate moves and player and update board state
func (g *Game) AttemptMoves(playerID string, moves []d.Move) (bool, string) {
	playerRole, exists := g.Participants[playerID]
	if !exists {
		return false, "Player not registered in this match."
	}
	if playerRole == d.Spectator {
		return false, "Participant is a spectator."
	}
	if playerID != g.CurrentPlayer {
		return false, "Player must wait for turn."
	}

	for _, move := range moves {
		success, message := g.checkMove(playerRole, move)
		fmt.Printf("%v, %v", success, message)
	}
	return false, ""
}

// moveIsValid(fromIdx: number, toIdx: number): boolean {
//     const fromTile = this.tiles[fromIdx];
//     const toTile = this.tiles[toIdx];
//     if (toTile !== TileStatus.Empty) {
//       return false;
//     }
//     const fromRowIdx = Math.floor((fromIdx - 1) / 4);
//     const fromColIdx = (fromIdx - 1) % 4;
//     const fromEdge = (fromRowIdx % 2 === 0) ? (fromColIdx === 3) : (fromColIdx === 0);
//     const validTargetOffsets = [];
//     if (fromTile === TileStatus.White
//       || fromTile === TileStatus.WhiteKing
//       || fromTile === TileStatus.BlackKing) {
//       const rowShift = (fromRowIdx % 2 === 0) ? -1 : 1;
//       validTargetOffsets.push(-4);
//       if (!fromEdge) {
//         validTargetOffsets.push(-4 + rowShift);
//       }
//     }
//     if (fromTile === TileStatus.Black
//       || fromTile === TileStatus.BlackKing
//       || fromTile === TileStatus.WhiteKing) {
//       const rowShift = (fromRowIdx % 2 === 0) ? -1 : 1;
//       validTargetOffsets.push(4);
//       if (!fromEdge) {
//         validTargetOffsets.push(4 + rowShift);
//       }
//     }

//     // TODO jump moves

//     return validTargetOffsets.map((offset) => offset + fromIdx).includes(toIdx);
//   }

func (g *Game) checkMove(playerRole d.GameRole, move d.Move) (bool, string) {
	fromTile := g.Board[move.IndexFrom]
	toTile := g.Board[move.IndexTo]

	if playerRole == d.PlayerWhite && !(fromTile == d.White || fromTile == d.WhitePromoted) {
		return false, "Player can only move their pieces."
	}

	if toTile != d.NoPiece {
		return false, "A piece can only be moved onto an empty tile."
	}

	// single step row range:
	// white, : fromIndex / 4 ) * 4 = top range exclusive
	// 			^ - 1 ) * 4 = bottom range inclusive

	// white  : specific squares
	// spot - 4 || spot - 5  :  evenOffset = 0, oddOffset = +1

	// black  : fromIndex /4 ) + 1 ) * 4 = bottom range inclusive
	//			^ + 1 ) * 4 = top range exclusive

	// black  : specfic squares
	// spot + 4 || spot + 5 : evenOffset = -1, oddOffset = 0

	return false, ""
}
