package service

import (
	"context"
	middleware "spy-cats/internal/middlewae"
	"spy-cats/internal/models"
	"spy-cats/internal/storage/postgre"
	db_connection "spy-cats/pkg/pg_connection"
)

type SpyCat interface {
	CreateSpyCat(ctx context.Context, cat *models.CatModel) (int64, error)
	DeleteSpyCat(ctx context.Context, id *int64) error

	GetSpyCat(ctx context.Context, id int64) (*models.CatModel, error)
	ListSpyCats(ctx context.Context) ([]models.CatModel, error)

	UpdateSpyCatSalary(ctx context.Context, id int64, newSalary float64) (float64, error)
}

type Mission interface {
	CreateMission(ctx context.Context, mission *models.MissionModel) (int64, error)
	DeleteMission(ctx context.Context, id int64) error

	GetMission(ctx context.Context, id int64) (*models.MissionModel, error)
	ListMissions(ctx context.Context) ([]models.MissionModel, error)

	CompleteMission(ctx context.Context, missionID int64) error
}

type Target interface {
	CreateTarget(ctx context.Context, target *models.TargetModel) (int, error)
	DeleteTarget(ctx context.Context, id int64) error
	GetTarget(ctx context.Context, id int64) (*models.TargetModel, error)
	ListTargetsByMission(ctx context.Context, missionID int64) ([]models.TargetModel, error)
	CompleteTarget(ctx context.Context, id int64) error
}

type Note interface {
	CreateNote(ctx context.Context, note *models.NoteModel) (int64, error)
	GetNote(ctx context.Context, id int64) (*models.NoteModel, error)
	ListNotesByTarget(ctx context.Context, targetID int64) ([]models.NoteModel, error)
	UpdateNote(ctx context.Context, noteID int64, newContent string) error
	DeleteNote(ctx context.Context, noteID int64) error
}
type Service struct {
	SpyCat
	Mission
	Target
	Note
	logger middleware.Logger
}

func NewService(db db_connection.Client, log middleware.Logger) *Service {

	return &Service{
		SpyCat:  postgre.NewCatRepository(db),
		Mission: postgre.NewMissionRepository(db),
		Target:  postgre.NewTargetRepository(db),
		Note:    postgre.NewNoteRepository(db),
		logger:  log,
	}
}
