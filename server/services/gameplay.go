package services

import (
	"context"

	controller "github.com/MettyS/checkers/server/controller"
	pb "github.com/MettyS/checkers/server/generated"
)

// GameplayServer server for gameplay service
type GameplayServer struct {
	pb.UnimplementedGameplayServiceServer
}

func convertMoveRequestInward(req *pb.MoveRequest) *controller.MoveRequest {
	domMoveRequest := &controller.MoveRequest{}
	domMoveRequest.MoveSet = make([]controller.Move, len(req.GetMoveSet()))

	for _, m := range req.GetMoveSet() {
		tempMove := controller.Move{
			IndexFrom: m.GetIndexFrom(),
			IndexTo:   m.GetIndexTo(),
		}
		domMoveRequest.MoveSet = append(domMoveRequest.MoveSet, tempMove)
	}
	return domMoveRequest
}

func convertMoveResponseOutward(res *controller.MoveResponse) *pb.MoveResponse {
	pbMoveResponse := pb.MoveResponse{}
	pbMoveResponse.MoveSuccess = res.MoveSuccess
	pbMoveResponse.Message = &res.Message

	return &pbMoveResponse
}

// MakeMoves (context.Context, *MoveRequest) (*MoveResponse, error)
func (s *GameplayServer) MakeMoves(ctx context.Context, req *pb.MoveRequest) (*pb.MoveResponse, error) {
	domMoveRequest := convertMoveRequestInward(req)
	// channels may not be best here since only expecting 1 output
	// I don't think channels are necessary here
	domMoveResponse := controller.HandleMakeMoves(ctx, domMoveRequest)

	return convertMoveResponseOutward(domMoveResponse), nil
}

// BoardUpdateSubscription (*BoardSubscriptionRequest, GameplayService_BoardUpdateSubscriptionServer) error
func (s *GameplayServer) BoardUpdateSubscription(req *pb.BoardSubscriptionRequest, updateStream pb.GameplayService_BoardUpdateSubscriptionServer) error {
	// updateStream.Send(*BoardUpdate) // returns error
	return nil
}
