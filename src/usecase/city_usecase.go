package usecase

import (
	"context"

	"github.com/SajjadRezaei/go-clean-architecture/config"
	"github.com/SajjadRezaei/go-clean-architecture/domain/filter"
	model "github.com/SajjadRezaei/go-clean-architecture/domain/model"
	"github.com/SajjadRezaei/go-clean-architecture/domain/repository"
	"github.com/SajjadRezaei/go-clean-architecture/usecase/dto"
)

type CityUsecase struct {
	base *BaseUsecase[model.City, dto.CreateCity, dto.UpdateCity, dto.City]
}

func NewCityUsecase(cfg *config.Config, repository repository.CityRepository) *CityUsecase {
	return &CityUsecase{
		base: NewBaseUsecase[model.City, dto.CreateCity, dto.UpdateCity, dto.City](cfg, repository),
	}
}

// Create
func (u *CityUsecase) Create(ctx context.Context, req dto.CreateCity) (dto.City, error) {
	return u.base.Create(ctx, req)
}

// Update
func (s *CityUsecase) Update(ctx context.Context, id int, req dto.UpdateCity) (dto.City, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CityUsecase) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CityUsecase) GetById(ctx context.Context, id int) (dto.City, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CityUsecase) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (*filter.PagedList[dto.City], error) {
	return s.base.GetByFilter(ctx, req)
}
