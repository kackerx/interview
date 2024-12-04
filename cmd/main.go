package main

import (
	"context"
	"fmt"

	"github.com/kackerx/interview/cmd/wire"
	"github.com/kackerx/interview/common/log"
	"github.com/kackerx/interview/internal/conf"
	"github.com/kackerx/interview/pkg/http"
	"github.com/kackerx/interview/pkg/validate"
)

func main() {
	validate.InitTrans()
	cfg := conf.NewConfig()
	app, clearUp, err := wire.NewWire(cfg)
	if err != nil {
		panic(err)
	}
	defer clearUp()

	log.New(context.Background()).Info(
		"服务启动",
		"serverName",
		cfg.Server.Name,
		"host: ",
		cfg.Server.Host,
		"Port: ",
		cfg.Server.Port,
	)

	http.Run(app, fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
}
