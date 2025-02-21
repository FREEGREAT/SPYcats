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
	panic("unimplemented")
}

// DeleteMission implements Mission.
func (m *MissionService) DeleteMission(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetMission implements Mission.
func (m *MissionService) GetMission(ctx context.Context, id int64) (*models.MissionModel, error) {
	panic("unimplemented")
}

// ListMissions implements Mission.
func (m *MissionService) ListMissions(ctx context.Context) ([]models.MissionModel, error) {
	panic("unimplemented")
}

// UpdateMission implements Mission.
func (m *MissionService) CompleteMission(ctx context.Context, missionID int64) error {
	panic("unimplemented")
}

func NewMissionService(repo storage.MissionRepository) Mission {
	return &MissionService{repo: repo}
}
