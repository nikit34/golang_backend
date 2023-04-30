package gapi

import (
	"testing"
	"time"

	db "github.com/nikit34/template_backend/db/sqlc"
	"github.com/nikit34/template_backend/util"
	"github.com/nikit34/template_backend/worker"
	"github.com/stretchr/testify/require"
)


func newTestServer(t *testing.T, store db.Store, taskDistributor worker.TaskDistributor) *Server {
	config := util.Config{
		TokenSymmetricKey: util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistributor)
	require.NoError(t, err)

	return server
}
