package handler

import (
	"Navigation-Web/models"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

/**
处理自定义中间件
*/
func JWTAuth() gin.HandlerFunc {

	return func(context *gin.Context) {
		log.Println("进入到jwt中间件")
		//todo 该中间间暂时不做鉴权 目前仅记录用户信息
		//获取用ip地址以及用使用的浏览器信息
		//获取ip地址如果是通过nginx 代理转发的请求则需要在nginx配置中 配置proxy_set_header X-Real-IP $remote_addr;
		//否则获取到的地址则会是nginx的内外ip
		ip := context.Request.Header.Get("X-Real-IP")
		if ip == "" {
			ip = context.Request.RemoteAddr
		}
		browserInfo := context.Request.Header.Get("User-Agent") //User-Agent
		fmt.Printf("ip %s/n,useDeviceInfo:%s/n", ip, browserInfo)
		log.Println(context.Request.RemoteAddr)
		userByip, err := models.UserFunc.GetUserByIp(ip)
		browserInfo = base64.StdEncoding.EncodeToString([]byte(browserInfo))
		if err != nil {
			log.Println(err, userByip)
		}
		if userByip.ID == 0 && err == nil {
			var user = models.User{
				Ip:          ip,
				BrowserInfo: browserInfo,
			}
			err = models.UserFunc.Created(&user)
			if err != nil {
				log.Println(err, user)
			}
		} else {
			log.Println(err, userByip)
		}

		//if true {
		//	context.JSON(http.StatusInternalServerError,"fail")
		//	context.Abort()
		//	return
		//}
		context.Next()
	}
}
