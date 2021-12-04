package controller

import (
	"Navigation-Web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NavigationInfoController struct{}

func GetAll(c *gin.Context) {
	var navigation models.NavigationInfo
	navigationInfos, err := navigation.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, navigationInfos)
	}
}
