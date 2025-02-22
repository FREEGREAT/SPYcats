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
	return t.repo.CompleteTarget(ctx, id)
}

// CreateTarget implements Target.
func (t *TargetService) CreateTarget(ctx context.Context, target *models.TargetModel) (int, error) {
	return t.repo.CreateTarget(ctx, target)
}

// DeleteTarget implements Target.
func (t *TargetService) DeleteTarget(ctx context.Context, id int64) error {
	return t.repo.DeleteTarget(ctx, id)
}

// GetTarget implements Target.
func (t *TargetService) GetTarget(ctx context.Context, id int64) (*models.TargetModel, error) {
	return t.repo.GetTarget(ctx, id)
}

// ListTargetsByMission implements Target.
func (t *TargetService) ListTargetsByMission(ctx context.Context, missionID int64) ([]models.TargetModel, error) {
	return t.repo.ListTargetsByMission(ctx, missionID)
}

func NewTargetService(repo storage.TargetRepository) Target {
	return &TargetService{repo: repo}
}
