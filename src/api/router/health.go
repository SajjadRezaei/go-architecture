package router

import (
	"github.com/SajjadRezaei/go-clean-architecture/api/handler"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handler.NewHealthHandler()

	r.GET("/", handler.Health)
}
