package controller

import (
	"Navigation-Web/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type navigationInfoController struct{}

var NavigationInfoController = new(navigationInfoController)

func (n *navigationInfoController) GetAll(c *gin.Context) {
	//获取用户公网ip以及用户设备号信息
	log.Println(c.Request.Header.Get("X-Real-IP"))
	log.Println(c.Request.Header.Get("User-Agent")) //User-Agent
	log.Println(c.Request.RemoteAddr)
	navigationInfos, err := models.NavigationInfofunc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, navigationInfos)
	}
}

func (n *navigationInfoController) Create(c *gin.Context) {
	var navigationINfo models.NavigationInfo
	c.BindJSON(&navigationINfo)
	err := models.NavigationInfofunc.Create(&navigationINfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func (n *navigationInfoController) Update(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, "参数无效")
	}
	idInt, _ := strconv.Atoi(id)
	navigationINfo, err := models.NavigationInfofunc.GetOne(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	log.Println(navigationINfo)
	c.BindJSON(&navigationINfo)
	log.Println(navigationINfo)
	err = models.NavigationInfofunc.Update(navigationINfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func (n *navigationInfoController) Delete(c *gin.Context) {
	var navigationINfo models.NavigationInfo
	c.BindJSON(&navigationINfo)
	log.Println(navigationINfo)
	err := models.NavigationInfofunc.Delete(&navigationINfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
