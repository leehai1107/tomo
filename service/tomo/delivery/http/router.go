package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leehai1107/tomo/pkg/apiwrapper"
)

type Router interface {
	Register(routerGroup gin.IRouter)
}

type routerImpl struct {
	handler IHandler
}

func NewRouter(
	handler IHandler,
) Router {
	return &routerImpl{
		handler: handler,
	}
}

func (p *routerImpl) Register(r gin.IRouter) {

	//routes for apis
	api := r.Group("api/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			apiwrapper.SendSuccess(c, time.Now())
		})
	}

	adminApi := api.Group("admin")
	{
		adminApi.POST("/create-account", func(c *gin.Context) {
			p.handler.Register(c)
		})
	}

	userApi := api.Group("user")
	{
		userApi.POST("/login", func(c *gin.Context) {
			p.handler.Login(c)
		})

		userApi.POST("/register", func(c *gin.Context) {
			p.handler.Register(c)
		})
	}

	// Add WebSocket route
	chatApi := api.Group("chat")
	{
		// Wrap the ServeWS handler in a safe handler
		chatApi.GET("/ws/:roomId", func(c *gin.Context) {
			defer func() {
				if r := recover(); r != nil {
					c.AbortWithStatusJSON(500, gin.H{"error": "Internal Server Error"})
				}
			}()
			p.handler.ServeWS(c)
		})
	}

}
