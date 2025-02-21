package service

import (
	"context"
	"spy-cats/internal/models"
	"spy-cats/internal/storage"
)

type SpyCatService struct {
	repo storage.CatRepository
}

// CreateSpyCat implements SpyCat.
func (s *SpyCatService) CreateSpyCat(ctx context.Context, cat *models.CatModel) (int64, error) {
	return s.repo.CreateSpyCat(ctx, cat)

}

// DeleteSpyCat implements SpyCat.
func (s *SpyCatService) DeleteSpyCat(ctx context.Context, id *int64) error {
	panic("unimplemented")
}

// GetSpyCat implements SpyCat.
func (s *SpyCatService) GetSpyCat(ctx context.Context, id int64) (*models.CatModel, error) {
	panic("unimplemented")
}

// ListSpyCats implements SpyCat.
func (s *SpyCatService) ListSpyCats(ctx context.Context) ([]models.CatModel, error) {
	panic("unimplemented")
}

// UpdateSpyCatSalary implements SpyCat.
func (s *SpyCatService) UpdateSpyCatSalary(ctx context.Context, id int64, newSalary float64) (float64, error) {
	panic("unimplemented")
}

func NewSpyCatService(repo storage.CatRepository) SpyCat {
	return &SpyCatService{repo: repo}
}
