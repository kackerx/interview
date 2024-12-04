package main

import (
	"github.com/kackerx/interview/internal/conf"
	"github.com/kackerx/interview/internal/data"
	"github.com/kackerx/interview/internal/server"
)

func main() {
	cfg := conf.NewConfig()
	db := data.NewDb(cfg)
	migrate := server.NewMigrate(db)
	migrate.Start()
}
