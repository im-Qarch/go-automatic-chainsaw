package api

import (
	"log"
	"os"
	"testing"
	"time"

	db "db.sqlc.dev/app/db/sqlc"
	"db.sqlc.dev/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load the config")
	}
	gin.SetMode(config.GinTestMode)
	os.Exit(m.Run())
}
