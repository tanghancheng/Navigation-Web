package controller

import (
	"Navigation-Web/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	navigationINfo, err := models.GetOne(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	log.Println(navigationINfo)
	c.BindJSON(&navigationINfo)
	log.Println(navigationINfo)
	err = models.Update(navigationINfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func Delete(c *gin.Context) {
	var navigationINfo models.NavigationInfo
	c.BindJSON(&navigationINfo)
	log.Println(navigationINfo)
	err := models.Delete(&navigationINfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
