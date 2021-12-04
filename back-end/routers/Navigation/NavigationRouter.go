package Navigation

import (
	"Navigation-Web/controller"
	"github.com/gin-gonic/gin"
)

type NavigationRouter struct {
}

var NavigationRouterApi = new(NavigationRouter)

func (n *NavigationRouter) InitNavigationRouter(r *gin.Engine) {
	navigationRouter := r.Group("navigation")
	{
		navigationRouter.GET("/navigation", controller.GetAll)
		navigationRouter.POST("/navigation", controller.Create)
		navigationRouter.PUT("/navigation/:id", controller.Update)
		navigationRouter.DELETE("/navigation", controller.Delete)
	}
}
