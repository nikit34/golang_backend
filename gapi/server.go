package gapi

import (
	"fmt"
	
	db "github.com/nikit34/template_backend/db/sqlc"
	"github.com/nikit34/template_backend/token"
	"github.com/nikit34/template_backend/util"
)


type Server struct {
	store db.Store
	config util.Config
	tokenMaker token.Maker
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}