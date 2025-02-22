package storage

import (
	"context"
	"errors"
	"spy-cats/internal/models"
)

var (
	ErrNotFound            = errors.New("record not found")
	ErrMissionCompleted    = errors.New("mission is already completed")
	ErrTargetCompleted     = errors.New("target is already completed")
	ErrTooManyTargets      = errors.New("mission cannot have more than 3 targets")
	ErrMissionHasSpyCat    = errors.New("cannot delete mission assigned to spy cat")
	ErrInvalidStatus       = errors.New("invalid mission status")
	ErrDuplicateAssignment = errors.New("spy cat already assigned to another mission")
)

type CatRepository interface {
	CreateSpyCat(ctx context.Context, cat *models.CatModel) (int64, error)
	DeleteSpyCat(ctx context.Context, id *int64) error
	GetSpyCat(ctx context.Context, id int64) (*models.CatModel, error)
	ListSpyCats(ctx context.Context) ([]models.CatModel, error)
	UpdateSpyCatSalary(ctx context.Context, id int64, newSalary float64) (float64, error)
}

type MissionRepository interface {
	CreateMission(ctx context.Context, mission *models.MissionModel) (int64, error)
	DeleteMission(ctx context.Context, id int64) error
	GetMission(ctx context.Context, id int64) (*models.MissionModel, error)
	ListMissions(ctx context.Context) ([]models.MissionModel, error)
	CompleteMission(ctx context.Context, missionID int64) error
}

type TargetRepository interface {
	CreateTarget(ctx context.Context, target *models.TargetModel) (int, error)
	DeleteTarget(ctx context.Context, id int64) error
	GetTarget(ctx context.Context, id int64) (*models.TargetModel, error)
	ListTargetsByMission(ctx context.Context, missionID int64) ([]models.TargetModel, error)
	CompleteTarget(ctx context.Context, id int64) error
}

type NoteRepository interface {
	CreateNote(ctx context.Context, note *models.NoteModel) (int64, error)
	GetNote(ctx context.Context, id int64) (*models.NoteModel, error)
	ListNotesByTarget(ctx context.Context, targetID int64) ([]models.NoteModel, error)
	UpdateNote(ctx context.Context, noteID int64, newContent string) error
	DeleteNote(ctx context.Context, noteID int64) error
}
