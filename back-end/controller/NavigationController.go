package controller

import (
	"Navigation-Web/models"
	"Navigation-Web/models/dto"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type navigationInfoController struct{}

var NavigationInfoController = new(navigationInfoController)

func (n *navigationInfoController) GetAll(c *gin.Context) {
	pageInfo:=dto.NewPageInfo()
	c.ShouldBind(pageInfo)
	pageInfo, err := models.NavigationInfofunc.GetAll(pageInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, pageInfo)
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
