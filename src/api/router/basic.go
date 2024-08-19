package router

import (
	"github.com/SajjadRezaei/go-clean-architecture/api/handler"
	"github.com/SajjadRezaei/go-clean-architecture/config"
	"github.com/gin-gonic/gin"
)

const GetByFilterExp string = "/get-by-filter"

func Country(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCountryHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}

func City(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCityHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}

func File(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewFileHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}

func Company(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCompanyHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}

func Color(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewColorHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}

func Year(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewPersianYearHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}
