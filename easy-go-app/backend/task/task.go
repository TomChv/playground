package task

import (
	"github.com/gin-gonic/gin"
	"izi-go.com/ent"
)

func Init(db *ent.Client, r *gin.Engine) {
	repo := NewRepo(db)
	ctrl := NewController(repo)

	taskEndpoint := r.Group("/task")

	taskEndpoint.POST("", ctrl.Create)
	taskEndpoint.GET("", ctrl.List)
	taskEndpoint.GET(":id", ctrl.Get)
	taskEndpoint.PUT(":id", ctrl.Update)
	taskEndpoint.DELETE(":id", ctrl.Delete)
	taskEndpoint.PUT("/start/:id", ctrl.Start)
	taskEndpoint.PUT("/done/:id", ctrl.Done)
}
