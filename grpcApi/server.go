package grpcApi

import (
	"fmt"

	"github.com/saalikmubeen/go-grpc-implementation/authToken"
	"github.com/saalikmubeen/go-grpc-implementation/pb"
	"github.com/saalikmubeen/go-grpc-implementation/utils"

	generated_db "github.com/saalikmubeen/go-grpc-implementation/db/sqlc"
)

// Server serves gRPC requests
type server struct {
	config    utils.Config
	store     generated_db.Store
	authToken authToken.Maker
	pb.UnimplementedSimpleBankServiceServer
}

// NewServer creates a new gRPC server.
func NewServer(config utils.Config, store generated_db.Store) (*server, error) {
	tokenMaker, err := authToken.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &server{
		config:    config,
		store:     store,
		authToken: tokenMaker,
	}

	return server, nil
}
