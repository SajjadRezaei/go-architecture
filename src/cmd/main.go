package main

import (
	"github.com/SajjadRezaei/go-clean-architecture/api"
	"github.com/SajjadRezaei/go-clean-architecture/config"
	"github.com/SajjadRezaei/go-clean-architecture/infra/cache"
	database "github.com/SajjadRezaei/go-clean-architecture/infra/persistence/database"
	"github.com/SajjadRezaei/go-clean-architecture/infra/persistence/migration"
	"github.com/SajjadRezaei/go-clean-architecture/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {

	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migration.Up1()

	api.InitServer(cfg)
}
