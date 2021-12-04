package routers

import (
	"Navigation-Web/routers/Navigation"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	r := gin.Default()
	Navigation.NavigationRouterApi.InitNavigationRouter(r)
	err := r.Run(":8888")
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
