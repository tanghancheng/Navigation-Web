package routers

import (
	"Navigation-Web/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.Default()
	navigationRouter := r.Group("navigation")
	{
		navigationRouter.GET("/navigation", controller.GetAll)
	}
	err := r.Run(":8888")
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
