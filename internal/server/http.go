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
) *gin.Engine {
	g := gin.Default()
	gin.SetMode(gin.DebugMode)

	v1 := g.Group("/v1")
	userRouter := v1.Group("/user")
	{
		noAuthRouter := userRouter.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
		}
	}

	return g
}
