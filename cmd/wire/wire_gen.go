// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func NewWire(cfg *conf.Conf) (*gin.Engine, func(), error) {
	jwt := middleware.NewJwt(cfg)
	handlerHandler := handler.NewHandler()
	appService := appservice.NewAppService()
	serviceService := service.NewService(jwt)
	db := data.NewDb(cfg)
	dataData := data.NewData(db)
	userRepo := data.NewUserRepo(dataData)
	userDomainService := service.NewUserDomainService(serviceService, userRepo, jwt)
	userAppService := appservice.NewUserAppService(appService, userDomainService)
	userHandler := handler.NewUserHandler(handlerHandler, userAppService)
	taskRepo := data.NewTaskRepo(dataData)
	taskDomainService := service.NewTaskDomainService(serviceService, taskRepo)
	documentRepo := data.NewDocumentRepo(dataData)
	documentDomainService := service.NewDocumentDomainService(serviceService, documentRepo)
	transaction := data.NewTransaction(dataData)
	taskAppService := appservice.NewtaskAppService(appService, taskDomainService, documentDomainService, transaction)
	taskHandler := handler.NewtaskHandler(handlerHandler, taskAppService)
	engine := server.NewHTTPServer(cfg, jwt, userHandler, taskHandler)
	return engine, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(data.NewDb, data.NewData, data.NewUserRepo, data.NewTaskRepo, data.NewTransaction, data.NewDocumentRepo)

var serviceSet = wire.NewSet(service.NewService, service.NewUserDomainService, service.NewTaskDomainService, service.NewDocumentDomainService)

var appServiceSet = wire.NewSet(appservice.NewAppService, appservice.NewUserAppService, appservice.NewtaskAppService)

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler, handler.NewtaskHandler)

var serverSet = wire.NewSet(server.NewHTTPServer)

var commonSet = wire.NewSet(middleware.NewJwt)
