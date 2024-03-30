package api

import (
	"log"
	"os"
	"testing"

	"db.sqlc.dev/app/utils"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load the config")
	}
	gin.SetMode(config.GinTestMode)
	os.Exit(m.Run())
}
