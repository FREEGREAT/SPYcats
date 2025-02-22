package service

import (
	"context"
	"spy-cats/internal/handler/middleware"
	"spy-cats/internal/models"
	"spy-cats/internal/storage"
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

type CatAPI interface {
	IsValidBreed(breedName string) (bool, error)
}

type Service struct {
	CatRepository     storage.CatRepository
	MissionRepository storage.MissionRepository
	TargetRepository  storage.TargetRepository
	NoteRepository    storage.NoteRepository
	logger            middleware.Logger
	CatApi            CatAPI
}

func NewService(catRepo storage.CatRepository, missionRepo storage.MissionRepository, noteRepo storage.NoteRepository,
	targetRepo storage.TargetRepository, log middleware.Logger, api CatAPI) *Service {

	return &Service{
		CatRepository:     catRepo,
		MissionRepository: missionRepo,
		TargetRepository:  targetRepo,
		NoteRepository:    noteRepo,
		logger:            log,
		CatApi:            api,
	}
}
