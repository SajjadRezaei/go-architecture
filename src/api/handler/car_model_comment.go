package handler

import (
	"github.com/SajjadRezaei/go-clean-architecture/api/dto"
	_ "github.com/SajjadRezaei/go-clean-architecture/api/helper"
	"github.com/SajjadRezaei/go-clean-architecture/config"
	"github.com/SajjadRezaei/go-clean-architecture/dependency"
	_ "github.com/SajjadRezaei/go-clean-architecture/domain/filter"
	"github.com/SajjadRezaei/go-clean-architecture/usecase"
	"github.com/gin-gonic/gin"
)

type CarModelCommentHandler struct {
	usecase *usecase.CarModelCommentUsecase
}

func NewCarModelCommentHandler(cfg *config.Config) *CarModelCommentHandler {
	return &CarModelCommentHandler{
		usecase: usecase.NewCarModelCommentUsecase(cfg, dependency.GetCarModelCommentRepository(cfg)),
	}
}

// CreateCarModelComment godoc
// @Summary Create a CarModelComment
// @Description Create a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelCommentRequest true "Create a CarModelComment"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/ [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelComment, dto.ToCarModelCommentResponse, h.usecase.Create)
}

// UpdateCarModelComment godoc
// @Summary Update a CarModelComment
// @Description Update a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelCommentRequest true "Update a CarModelComment"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-comments/{id} [put]
// @Security AuthBearer
func (h *CarModelCommentHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelComment, dto.ToCarModelCommentResponse, h.usecase.Update)
}

// DeleteCarModelComment godoc
// @Summary Delete a CarModelComment
// @Description Delete a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-comments/{id} [delete]
// @Security AuthBearer
func (h *CarModelCommentHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCarModelComment godoc
// @Summary Get a CarModelComment
// @Description Get a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-comments/{id} [get]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelCommentResponse, h.usecase.GetById)
}

// GetCarModelComments godoc
// @Summary Get CarModelComments
// @Description Get CarModelComments
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.CarModelCommentResponse]} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelCommentResponse, h.usecase.GetByFilter)
}
