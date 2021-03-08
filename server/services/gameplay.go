package services

import (
	"context"

	"github.com/MettyS/checkers/server/controller"
	pb "github.com/MettyS/checkers/server/generated"
	d "github.com/MettyS/checkers/server/shared"
)

// GameplayServer server for gameplay service
type GameplayServer struct {
	pb.UnimplementedGameplayServiceServer
}

func convertMoveRequestInward(req *pb.MoveRequest) d.MoveRequest {
	domMoveRequest := d.MoveRequest{}
	domMoveRequest.MoveSet = make([]d.Move, len(req.GetMoveSet()))
	domMoveRequest.GameID = req.GetGameId()
	domMoveRequest.PlayerID = req.GetPlayerId()

	for _, m := range req.GetMoveSet() {
		tempMove := d.Move{
			IndexFrom: m.GetIndexFrom(),
			IndexTo:   m.GetIndexTo(),
		}
		domMoveRequest.MoveSet = append(domMoveRequest.MoveSet, tempMove)
	}
	return domMoveRequest
}

func convertMoveResponseOutward(res d.MoveResponse) *pb.MoveResponse {
	pbMoveResponse := pb.MoveResponse{}
	pbMoveResponse.MoveSuccess = res.MoveSuccess
	pbMoveResponse.Message = &res.Message

	return &pbMoveResponse
}

func convertBoardSubscriptionRequestInward(req *pb.BoardSubscriptionRequest) d.BoardSubscriptionRequest {
	domBoardSubscriptionRequest := d.BoardSubscriptionRequest{
		GameID: req.GetGameId(),
	}

	return domBoardSubscriptionRequest
}

func convertBoardUpdateOutward(res d.BoardUpdate) *pb.BoardUpdate {
	pbBoardUpdate := pb.BoardUpdate{}

	pbBoard := pb.BoardState{
		Board: make([]pb.Tile, d.BoardSize),
	}
	for _, t := range res.BoardState.Board {
		pbBoard.Board = append(pbBoard.Board, pb.Tile(int32(t)))
	}
	pbBoardUpdate.BoardState = &pbBoard

	pbMoves := make([]*pb.Move, len(res.PrevMoveSet))
	for _, m := range res.PrevMoveSet {
		tempMove := pb.Move{
			IndexFrom: m.IndexFrom,
			IndexTo:   m.IndexTo,
		}
		pbMoves = append(pbMoves, &tempMove)
	}
	pbBoardUpdate.PrevMoveSet = pbMoves
	return &pbBoardUpdate
}

// MakeMoves (context.Context, *MoveRequest) (*MoveResponse, error)
func (s *GameplayServer) MakeMoves(ctx context.Context, req *pb.MoveRequest) (*pb.MoveResponse, error) {
	domMoveRequest := convertMoveRequestInward(req)
	domMoveResponse, err := controller.HandleMakeMoves(ctx, domMoveRequest)

	if err != nil {
		return nil, err
	}

	pbMoveResponse := convertMoveResponseOutward(domMoveResponse)
	return pbMoveResponse, nil
}

// BoardUpdateSubscription (*BoardSubscriptionRequest, GameplayService_BoardUpdateSubscriptionServer) error
func (s *GameplayServer) BoardUpdateSubscription(req *pb.BoardSubscriptionRequest, updateStream pb.GameplayService_BoardUpdateSubscriptionServer) error {
	domBoardSubscriptionRequest := convertBoardSubscriptionRequestInward(req)
	domBoardUpdate, err := controller.HandleBoardUpdateSubscription(domBoardSubscriptionRequest)

	if err != nil {
		return err
	}

	pbBoardUpdate := convertBoardUpdateOutward(domBoardUpdate)
	return updateStream.Send(pbBoardUpdate) // returns error
}
