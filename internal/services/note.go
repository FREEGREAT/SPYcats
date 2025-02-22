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
	return n.repo.CreateNote(ctx, note)
}

// GetNote implements Note.
func (n *NoteService) GetNote(ctx context.Context, id int64) (*models.NoteModel, error) {
	return n.repo.GetNote(ctx, id)
}

// ListNotesByTarget implements Note.
func (n *NoteService) ListNotesByTarget(ctx context.Context, targetID int64) ([]models.NoteModel, error) {
	return n.repo.ListNotesByTarget(ctx, targetID)
}

// RemoveNote implements Note.
func (n *NoteService) DeleteNote(ctx context.Context, noteID int64) error {
	return n.repo.DeleteNote(ctx, noteID)
}

// UpdateNote implements Note.
func (n *NoteService) UpdateNote(ctx context.Context, noteID int64, newContent string) error {
	return n.repo.UpdateNote(ctx, noteID, newContent)
}

func NewNoteService(repo storage.NoteRepository) Note {
	return &NoteService{repo: repo}
}
