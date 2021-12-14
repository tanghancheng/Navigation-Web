package handler

import (
	"Navigation-Web/models"
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var mutex = sync.RWMutex{}

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
			//如果不是通过nginx代理转发的地址会有一个端口号 
			ip = strings.Split(context.Request.RemoteAddr, ":")[0]
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
	fmt.Printf("ip %s ,useDeviceInfo:%s  ", user.Ip, user.BrowserInfo)
	userByIp, err := models.UserFunc.GetUserByIp(user.Ip)
	user.BrowserInfo = base64.StdEncoding.EncodeToString([]byte(user.BrowserInfo))
	if err != nil {
		log.Println(err, userByIp)
	}
	if userByIp.ID == 0 && err == nil {
		mutex.Lock()
		err = models.UserFunc.Created(&user)
		mutex.Unlock()
		if err != nil {
			log.Println(err, user)
		}
	} else {
		models.UserFunc.Update(&userByIp)
		if err != nil {
			log.Println(err, user)
		}
	}
}
