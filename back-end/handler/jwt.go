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

		var userChan = make(chan models.User)

		go createUserInfo(userChan)
		var user = models.User{
			Ip:          ip,
			BrowserInfo: browserInfo,
		}
		userChan <- user

		//if true {
		//	context.JSON(http.StatusInternalServerError,"fail")
		//	context.Abort()
		//	return
		//}
		context.Next()
	}
}

func createUserInfo(ch chan models.User) {
	//todo 后续增加redis 记录用户信息 就不需要每次都操作數據庫了
	user := <-ch
	fmt.Printf("ip %s/n,useDeviceInfo:%s/n", user.Ip, user.BrowserInfo)
	userByip, err := models.UserFunc.GetUserByIp(user.Ip)
	user.BrowserInfo = base64.StdEncoding.EncodeToString([]byte(user.BrowserInfo))
	if err != nil {
		log.Println(err, userByip)
	}
	if userByip.ID == 0 && err == nil {
		err = models.UserFunc.Created(&user)
		if err != nil {
			log.Println(err, user)
		}
	} else {
		log.Println(err, userByip)
	}
}
