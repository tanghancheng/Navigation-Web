package controller

import (
	"Navigation-Web/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type NavigationInfoController struct{}

func GetAll(c *gin.Context) {
	navigationInfos, err := models.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, navigationInfos)
	}
}

func Create(c *gin.Context) {
	var navigationINfo models.NavigationInfo
	c.BindJSON(&navigationINfo)
	err := models.Create(&navigationINfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func Update(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, "参数无效")
	}
	idInt, _ := strconv.Atoi(id)
	one, err := models.GetOne(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.BindJSON(&one)
	err = models.Update(&one)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func Delete(c *gin.Context) {
	var navigationINfo models.NavigationInfo
	c.BindJSON(&navigationINfo)
	err := models.Delete(&navigationINfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
