package gapi

import (
	"fmt"

	db "github.com/nikit34/template_backend/db/sqlc"
	"github.com/nikit34/template_backend/pb"
	"github.com/nikit34/template_backend/token"
	"github.com/nikit34/template_backend/util"
	"github.com/nikit34/template_backend/worker"
)


type Server struct {
	pb.UnimplementedTemplateBackendServer
	store db.Store
	config util.Config
	tokenMaker token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}