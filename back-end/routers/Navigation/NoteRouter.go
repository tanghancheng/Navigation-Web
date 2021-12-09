package Navigation

import (
	"Navigation-Web/controller"

	"github.com/gin-gonic/gin"
)

type noteRouter struct{}

var NoteRouter = new(noteRouter)

func(note *noteRouter) InitNoteRouter(r *gin.RouterGroup) (err error) {
	noteRouter:=r.Group("note")
	{
		noteRouter.GET("/getAll", controller.NoteController.GetAll)
		noteRouter.POST("/create",controller.NoteController.Create)
		noteRouter.PUT("/update/:id",controller.NoteController.Update)
		noteRouter.DELETE("/delete/:id",controller.NoteController.Delete)
	}
	return nil
}