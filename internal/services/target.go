package service

import (
	"context"
	"spy-cats/internal/models"
	"spy-cats/internal/storage"
)

type TargetService struct {
	repo storage.TargetRepository
}

// CompleteTarget implements Target.
func (t *TargetService) CompleteTarget(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// CreateTarget implements Target.
func (t *TargetService) CreateTarget(ctx context.Context, target *models.TargetModel) (int, error) {
	panic("unimplemented")
}

// DeleteTarget implements Target.
func (t *TargetService) DeleteTarget(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetTarget implements Target.
func (t *TargetService) GetTarget(ctx context.Context, id int64) (*models.TargetModel, error) {
	panic("unimplemented")
}

// ListTargetsByMission implements Target.
func (t *TargetService) ListTargetsByMission(ctx context.Context, missionID int64) ([]models.TargetModel, error) {
	panic("unimplemented")
}

func NewTargetService(repo storage.TargetRepository) Target {
	return &TargetService{repo: repo}
}
