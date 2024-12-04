//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/kackerx/interview/common/middleware"
	"github.com/kackerx/interview/internal/appservice"
	"github.com/kackerx/interview/internal/conf"
	"github.com/kackerx/interview/internal/data"
	"github.com/kackerx/interview/internal/domain/service"
	"github.com/kackerx/interview/internal/handler"
	"github.com/kackerx/interview/internal/server"
)

var repositorySet = wire.NewSet(
	data.NewDb,
	data.NewData,
	data.NewUserRepo,
	data.NewTaskRepo,
	data.NewTransaction,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserDomainService,
	service.NewTaskDomainService,
)

var appServiceSet = wire.NewSet(
	appservice.NewAppService,
	appservice.NewUserAppService,
	appservice.NewtaskAppService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewtaskHandler,
)

var serverSet = wire.NewSet(server.NewHTTPServer)

var commonSet = wire.NewSet(middleware.NewJwt)

func NewWire(cfg *conf.Conf) (*gin.Engine, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		appServiceSet,
		handlerSet,
		serverSet,
		commonSet,
	))
}
