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

type CarModelHandler struct {
	usecase *usecase.CarModelUsecase
}

func NewCarModelHandler(cfg *config.Config) *CarModelHandler {
	return &CarModelHandler{
		usecase: usecase.NewCarModelUsecase(cfg, dependency.GetCarModelRepository(cfg)),
	}
}

// CreateCarModel godoc
// @Summary Create a CarModel
// @Description Create a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelRequest true "Create a CarModel"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/ [post]
// @Security AuthBearer
func (h *CarModelHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModel, dto.ToCarModelResponse, h.usecase.Create)
}

// UpdateCarModel godoc
// @Summary Update a CarModel
// @Description Update a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelRequest true "Update a CarModel"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-models/{id} [put]
// @Security AuthBearer
func (h *CarModelHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModel, dto.ToCarModelResponse, h.usecase.Update)
}

// DeleteCarModel godoc
// @Summary Delete a CarModel
// @Description Delete a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-models/{id} [delete]
// @Security AuthBearer
func (h *CarModelHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCarModel godoc
// @Summary Get a CarModel
// @Description Get a CarModel
// @Tags CarModels
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelResponse} "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-models/{id} [get]
// @Security AuthBearer
func (h *CarModelHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelResponse, h.usecase.GetById)
}

// GetCarModels godoc
// @Summary Get CarModels
// @Description Get CarModels
// @Tags CarModels
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.CarModelResponse]} "CarModel response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-models/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelResponse, h.usecase.GetByFilter)
}
