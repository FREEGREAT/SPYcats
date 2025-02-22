package service

import (
	"context"
	"spy-cats/internal/models"
	"spy-cats/internal/storage"
)

type MissionService struct {
	repo storage.MissionRepository
}

// CreateMission implements Mission.
func (m *MissionService) CreateMission(ctx context.Context, mission *models.MissionModel) (int64, error) {
	return m.repo.CreateMission(ctx, mission)
}

// DeleteMission implements Mission.
func (m *MissionService) DeleteMission(ctx context.Context, id int64) error {
	return m.repo.DeleteMission(ctx, id)
}

// GetMission implements Mission.
func (m *MissionService) GetMission(ctx context.Context, id int64) (*models.MissionModel, error) {
	return m.repo.GetMission(ctx, id)
}

// ListMissions implements Mission.
func (m *MissionService) ListMissions(ctx context.Context) ([]models.MissionModel, error) {
	return m.repo.ListMissions(ctx)
}

// UpdateMission implements Mission.
func (m *MissionService) CompleteMission(ctx context.Context, missionID int64) error {
	return m.repo.CompleteMission(ctx, missionID)
}

func NewMissionService(repo storage.MissionRepository) Mission {
	return &MissionService{repo: repo}
}
