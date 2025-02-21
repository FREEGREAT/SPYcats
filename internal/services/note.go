package service

import (
	"context"
	"spy-cats/internal/models"
	"spy-cats/internal/storage"
)

type NoteService struct {
	repo storage.NoteRepository
}

// AddNote implements Note.
func (n *NoteService) CreateNote(ctx context.Context, note *models.NoteModel) (int64, error) {
	panic("unimplemented")
}

// GetNote implements Note.
func (n *NoteService) GetNote(ctx context.Context, id int64) (*models.NoteModel, error) {
	panic("unimplemented")
}

// ListNotesByTarget implements Note.
func (n *NoteService) ListNotesByTarget(ctx context.Context, targetID int64) ([]models.NoteModel, error) {
	panic("unimplemented")
}

// RemoveNote implements Note.
func (n *NoteService) DeleteNote(ctx context.Context, noteID int64) error {
	panic("unimplemented")
}

// UpdateNote implements Note.
func (n *NoteService) UpdateNote(ctx context.Context, noteID int64, newContent string) error {
	panic("unimplemented")
}

func NewNoteService(repo storage.NoteRepository) Note {
	return &NoteService{repo: repo}
}
