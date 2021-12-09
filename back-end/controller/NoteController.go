package controller

import (
	"Navigation-Web/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type noteController struct {
}

var NoteController = new(noteController)

func (noteController *noteController) GetAll(c *gin.Context) {
	notes, err := models.NoteFunc.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, notes)
	}
}

func (noteController *noteController) Create(c *gin.Context) {
	var note models.Note
	c.BindJSON(&note)
	err := models.NoteFunc.Create(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func (noteController *noteController) Update(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, "参数无效")
	}
	idInt, _ := strconv.Atoi(id)
	note, err := models.NoteFunc.GetOne(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.BindJSON(&note)
	err = models.NoteFunc.Update(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func (noteController *noteController) Delete(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusInternalServerError, "参数无效")
	}
	idInt, _ := strconv.Atoi(id)
	note, err := models.NoteFunc.GetOne(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	note.DeleteStatus=1
	err = models.NoteFunc.Update(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, "success")
	}
}
