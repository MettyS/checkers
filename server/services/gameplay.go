package services

import (
	"context"

	controller "github.com/MettyS/checkers/server/controller"
	pb "github.com/MettyS/checkers/server/generated"
)

// Server ...
type Server struct {
	pb.UnimplementedGameplayServiceServer
}

// MakeMoves (context.Context, *MoveRequest) (*MoveResponse, error)
func (s *Server) MakeMoves(ctx context.Context, req *pb.MoveRequest) *pb.MoveResponse {
	domMoveRequest := &controller.MoveRequest{}
	domMoveRequest.MoveSet = make([]controller.Move, len(req.MoveSet))

	//domMoveRequest.MoveSet = &(req.MoveSet)

	// make factory patterns instead of in-line
	for i, m := range req.MoveSet {
		tempMove := controller.Move{
			IndexFrom: m.IndexFrom,
			IndexTo:   m.IndexTo,
		}
		domMoveRequest.MoveSet = append(domMoveRequest.MoveSet, tempMove)

	}

	return nil
}

// BoardUpdateSubscription (*BoardSubscriptionRequest, GameplayService_BoardUpdateSubscriptionServer) error
func (s *Server) BoardUpdateSubscription(req *pb.BoardSubscriptionRequest, updateStream pb.GameplayService_BoardUpdateSubscriptionServer) error {
	// updateStream.Send(*BoardUpdate) // returns error
	return nil
}
