package services

import (
	"context"

	"github.com/MettyS/checkers/server/controller"
	pb "github.com/MettyS/checkers/server/generated"
	d "github.com/MettyS/checkers/server/shared"
)

// GameControlServer server for game control service
type GameControlServer struct {
	pb.UnimplementedGameControlServiceServer
}

func convertGameStartRequestInward(req *pb.GameStartRequest) d.GameStartRequest {
	domGameStartRequest := d.GameStartRequest{}
	domGameStartRequest.GameID = req.GetGameId()
	return domGameStartRequest
}

func convertGameStartResponseOutward(req d.GameStartResponse) *pb.GameStartResponse {
	pbGameStartResponse := pb.GameStartResponse{}
	pbGameStartResponse.Message = &req.Message

	pbGameStartResponse.GameData = &pb.GameStartRole{
		PlayerRole: pb.GameRole_SPECTATOR, // TODO
		GameId:     req.GameData.GameID,
		PlayerId:   req.GameData.PlayerID,
	}
	return &pbGameStartResponse
}

// StartGame (context.Context, *GameStartRequest) (*GameStartResponse, error)
func (s *GameControlServer) StartGame(ctx context.Context, req *pb.GameStartRequest) (*pb.GameStartResponse, error) {
	domGameStartRequest := convertGameStartRequestInward(req)
	domGameStartResponse, err := controller.HandleStartGame(domGameStartRequest)
	return convertGameStartResponseOutward(domGameStartResponse), err
}
