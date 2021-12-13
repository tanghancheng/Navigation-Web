package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

/**
处理自定义中间件
*/
func JWTAuth() gin.HandlerFunc {

	return func(context *gin.Context) {
		log.Println("进入到jwt中间件")
		//if true {
		//	context.JSON(http.StatusInternalServerError,"fail")
		//	context.Abort()
		//	return
		//}
		context.Next()
	}
}
