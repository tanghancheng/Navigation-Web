package routers

import (
	"Navigation-Web/handler"
	"Navigation-Web/routers/Navigation"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.Default()
	apiRouter := r.Group("api")
	apiRouter.Use(handler.JWTAuth())
	//init router
	{
		Navigation.NavigationRouter.InitNavigationRouter(apiRouter)
		Navigation.NoteRouter.InitNoteRouter(apiRouter)
	}
	err := r.Run(":9099")
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
