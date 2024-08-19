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

type CarModelYearHandler struct {
	usecase *usecase.CarModelYearUsecase
}

func NewCarModelYearHandler(cfg *config.Config) *CarModelYearHandler {
	return &CarModelYearHandler{
		usecase: usecase.NewCarModelYearUsecase(cfg, dependency.GetCarModelYearRepository(cfg)),
	}
}

// CreateCarModelYear godoc
// @Summary Create a CarModelYear
// @Description Create a CarModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelYearRequest true "Create a CarModelYear"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/ [post]
// @Security AuthBearer
func (h *CarModelYearHandler) Create(c *gin.Context) {
	Create(c, dto.ToCreateCarModelYear, dto.ToCarModelYearResponse, h.usecase.Create)
}

// UpdateCarModelYear godoc
// @Summary Update a CarModelYear
// @Description Update a CarModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelYearRequest true "Update a CarModelYear"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-years/{id} [put]
// @Security AuthBearer
func (h *CarModelYearHandler) Update(c *gin.Context) {
	Update(c, dto.ToUpdateCarModelYear, dto.ToCarModelYearResponse, h.usecase.Update)
}

// DeleteCarModelYear godoc
// @Summary Delete a CarModelYear
// @Description Delete a CarModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-years/{id} [delete]
// @Security AuthBearer
func (h *CarModelYearHandler) Delete(c *gin.Context) {
	Delete(c, h.usecase.Delete)
}

// GetCarModelYear godoc
// @Summary Get a CarModelYear
// @Description Get a CarModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-years/{id} [get]
// @Security AuthBearer
func (h *CarModelYearHandler) GetById(c *gin.Context) {
	GetById(c, dto.ToCarModelYearResponse, h.usecase.GetById)
}

// GetCarModelYears godoc
// @Summary Get CarModelYears
// @Description Get CarModelYears
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param Request body filter.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=filter.PagedList[dto.CarModelYearResponse]} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelYearHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, dto.ToCarModelYearResponse, h.usecase.GetByFilter)
}
