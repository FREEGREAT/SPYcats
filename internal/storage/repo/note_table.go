package repo

import (
	"context"
	"database/sql"
	"errors"
	"spy-cats/internal/models"
	"spy-cats/internal/storage"
	db_connection "spy-cats/pkg/pg_connection"
)

type noteTableImpl struct {
	client db_connection.Client
}

// AddNote implements storage.NoteRepository.
func (n *noteTableImpl) CreateNote(ctx context.Context, note *models.NoteModel) (int64, error) {
	var isCompleted bool
	err := n.client.QueryRow(ctx, `SELECT is_completed FROM targets WHERE id = $1`, note.TargetID).Scan(&isCompleted)
	if err != nil {
		return 0, err
	}
	if isCompleted {
		return 0, errors.New("cannot add notes to a completed target")
	}

	query := `
		INSERT INTO notes (target_id, content, created_at, updated_at) 
		VALUES ($1, $2, NOW(), NOW()) 
		RETURNING id`
	var id int
	err = n.client.QueryRow(ctx, query, note.TargetID, note.Content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

// GetNote implements storage.NoteRepository.
func (n *noteTableImpl) GetNote(ctx context.Context, id int64) (*models.NoteModel, error) {
	query := `
		SELECT id, target_id, content, created_at, updated_at 
		FROM notes WHERE id = $1`
	note := &models.NoteModel{}
	err := n.client.QueryRow(ctx, query, id).Scan(
		&note.ID, &note.TargetID, &note.Content, &note.CreatedAt, &note.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Cant find notes")
		}
		return nil, err
	}
	return note, nil
}

// ListNotesByTarget implements storage.NoteRepository.
func (n *noteTableImpl) ListNotesByTarget(ctx context.Context, targetID int64) ([]models.NoteModel, error) {
	query := `SELECT id, target_id, content, created_at, updated_at FROM notes WHERE target_id = $1`
	rows, err := n.client.Query(ctx, query, targetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.NoteModel
	for rows.Next() {
		var note models.NoteModel
		if err := rows.Scan(&note.ID, &note.TargetID, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	return notes, nil
}

// RemoveNote implements storage.NoteRepository.
func (n *noteTableImpl) DeleteNote(ctx context.Context, noteID int64) error {
	query := `DELETE FROM notes WHERE id = $1`
	_, err := n.client.Exec(ctx, query, noteID)
	return err
}

// UpdateNote implements storage.NoteRepository.
func (n *noteTableImpl) UpdateNote(ctx context.Context, noteID int64, newContent string) error {
	var isCompleted bool
	err := n.client.QueryRow(ctx, `SELECT t.is_completed FROM notes n JOIN targets t ON n.target_id = t.id WHERE n.id = $1`, noteID).Scan(&isCompleted)
	if err != nil {
		return err
	}
	if isCompleted {
		return errors.New("cannot update notes for a completed target")
	}

	query := `UPDATE notes SET content = $1, updated_at = NOW() WHERE id = $2`
	_, err = n.client.Exec(ctx, query, newContent, noteID)
	return err
}

func NewNoteRepository(client db_connection.Client) storage.NoteRepository {
	return &noteTableImpl{client: client}
}
