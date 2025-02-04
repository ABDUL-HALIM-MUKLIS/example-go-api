package example

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.RouterGroup) {
	r.GET("/", index)
	r.POST("/", store)
	r.PUT("/:id", update)
	r.DELETE("/:id", destroy)
}