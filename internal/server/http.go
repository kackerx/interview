package server

import (
	"github.com/gin-gonic/gin"

	"github.com/kackerx/interview/common/middleware"
	"github.com/kackerx/interview/internal/conf"
	"github.com/kackerx/interview/internal/handler"
)

func NewHTTPServer(
	cfg *conf.Conf,
	jwt *middleware.JWT,
	userHandler *handler.UserHandler,
	taskHandler *handler.TaskHandler,
) *gin.Engine {
	g := gin.Default()
	// gin.SetMode(gin.DebugMode)
	gin.SetMode(cfg.Server.Mode)

	// 注册路由
	v1 := g.Group("/v1")
	registerUserRouter(v1, userHandler)
	registerTaskRouter(v1, taskHandler, jwt)

	return g
}

// registerUserRouter 用户路由
func registerUserRouter(g *gin.RouterGroup, userHandler *handler.UserHandler) {
	userRouter := g.Group("/auth")
	noAuthRouter := userRouter.Group("/")
	{
		noAuthRouter.POST("/users", userHandler.RegisterUser)
		noAuthRouter.POST("/login", userHandler.LoginUser)
	}
}

// registerTaskRouter 任务路由
func registerTaskRouter(g *gin.RouterGroup, taskHandler *handler.TaskHandler, jwt *middleware.JWT) {
	taskRouter := g.Group("/tasks") // Use(middleware.StrictAuth(jwt))
	AuthRouter := taskRouter.Group("/").Use(middleware.StrictAuth(jwt))
	{
		AuthRouter.POST("", taskHandler.CreateTask)
		AuthRouter.POST("/register", taskHandler.CreateTask)
	}
}
